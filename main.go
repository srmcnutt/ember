package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// store credentials in a map for easy retrieval
var creds = make(map[string]string)

// store api endpoints in a map for easy retrieval
var endPoints = make(map[string]string)

func main() {
	banner()
	creds = getCreds()
	color.Cyan("Authenticating to FMC, hang tight...\n")
	fmt.Println("")
	initEndpoints()
	if !strings.Contains(creds["fmc_host"], "cdo.cisco.com") {
		getAuthToken(endPoints["auth"])
	} else {
		creds["domain"] = getDomains(creds["fmc_host"])[0].ID
	}

	// need to init a second time now that we have the uuid for the domain
	initEndpoints()

	// launch menu
	menu()
}

// read in environment vars to connect to FMC
func getCreds() map[string]string {
	creds["fmc_host"] = os.Getenv("FMC_HOST")
	creds["fmc_user"] = os.Getenv("FMC_USER")
	creds["fmc_password"] = os.Getenv("FMC_PASSWORD")

	if creds["fmc_host"] == "" {
		//fmt.Println("FMC_HOST Environment Variable not set")
		var fmc_host string
		fmt.Print("Enter FMC Hostname or Address (c to cancel and exit): ")
		fmt.Scanln(&fmc_host)
		creds["fmc_host"] = fmc_host
		if fmc_host == "c" {
			fmt.Println("Exiting...")
			os.Exit(1)
		}
		//fmt.Println("FMC_HOST set to:", fmc_host)
	}

	if creds["fmc_user"] == "" {
		if !strings.Contains(creds["fmc_host"], "cdo.cisco.com") {
			//fmt.Println("FMC_USER Environment Variable not set")
			var fmc_user string
			fmt.Print("Enter FMC Username (c to cancel and exit): ")
			fmt.Scanln(&fmc_user)
			creds["fmc_user"] = fmc_user
			if fmc_user == "c" {
				fmt.Println("Exiting...")
				os.Exit(1)
			}
			//fmt.Println("FMC_USER set to:", fmc_user)
		}
	}

	if creds["fmc_password"] == "" {
		//fmt.Println("FMC_PASSWORD Environment Variable not set")
		var fmc_password string
		fmt.Print("Enter FMC Password (c to cancel and exit): ")
		fmt.Scanln(&fmc_password)
		creds["fmc_password"] = fmc_password
		if fmc_password == "c" {
			fmt.Println("Exiting...")
			os.Exit(1)
		}
		//fmt.Println("FMC_PASSWORD set")
	}
	return creds
}

func menu() {

	menuOptions := []string{
		"1. Get FMC Details",
		"2. Get Device List",
		"3. Get Device Details",
		"0. Exit",
	}

	for {
		for _, option := range menuOptions {
			fmt.Println(option)
		}
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter number of your choice: ")
		input, _, err := r.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		choice := int(input - 48)
		//choice, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("There is an error: ", err)
		}

		switch choice {
		case 1:
			//aaa
		case 2:
			devices := getDevices()
			fmt.Print("\n")
			printTable(devices)
			color.Blue("\nTotal number of sensors: %s", strconv.Itoa(len(devices)))
			fmt.Print("\n")
		case 3:
			//aaa
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input - please enter a number")
		}
	}
}

// print banner
func banner() {
	fmt.Println(art)
}

// populate endpoints map
func initEndpoints() {
	endPoints["auth"] = fmt.Sprintf("https://%s/api/fmc_platform/v1/auth/generatetoken", creds["fmc_host"])
	endPoints["devices"] = fmt.Sprintf("https://%s/api/fmc_config/v1/domain/%s/devices/devicerecords", creds["fmc_host"], creds["domain"])
}

// generic function to make rest api call to FMC and pass the body back
func fmcCall(url string) []byte {
	// make a transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}

	// make a client
	client := &http.Client{Transport: tr}

	// set up request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// build header
	if !strings.Contains(creds["fmc_host"], "cdo.cisco.com") {
		req.Header = http.Header{
			"Content-Type":        {"application/json"},
			"Accept":              {"application/json"},
			"X-auth-access-token": {creds["token"]},
		}
	} else {
		var bearer = "Bearer " + creds["fmc_password"]
		req.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
			"Authorization": {bearer},
		}
	}

	// add some parameters to the request
	q := req.URL.Query()
	q.Add("limit", "150")
	q.Add("expanded", "true")

	req.URL.RawQuery = q.Encode()

	// execute request & assign to res variable
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// dump the header
	//fmt.Println(res)

	//response body
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	//fmt.Println(string(b))
	//fmt.Println(url)
	return b
}

func getResponse(url string) APIResponse {
	// make a transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}

	// make a client
	client := &http.Client{Transport: tr}

	// set up request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// build header
	if !strings.Contains(creds["fmc_host"], "cdo.cisco.com") {
		req.Header = http.Header{
			"Content-Type":        {"application/json"},
			"Accept":              {"application/json"},
			"X-auth-access-token": {creds["token"]},
		}
	} else {
		var bearer = "Bearer " + creds["fmc_password"]
		req.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
			"Authorization": {bearer},
		}
	}

	// execute request & assign to res variable
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//response body
	defer res.Body.Close()

	// store the domain from the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(body))
	var r APIResponse
	json.Unmarshal(body, &r)

	return r
}

func getDomains(baseUrl string) []Domain {
	url := "https://" + baseUrl + "/api/fmc_platform/v1/info/domain"
	r := getResponse(url)
	var domains []Domain
	domainJson, _ := json.Marshal(r.Items)
	json.Unmarshal(domainJson, &domains)

	return domains
}

func getAuthToken(url string) {
	// make a transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
	}

	// make a client
	client := &http.Client{Transport: tr}

	// set up request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	// build header
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}

	// add basic authentication to our header
	req.SetBasicAuth(creds["fmc_user"], creds["fmc_password"])

	// fmt.Println(req.URL.String())

	// execute request & assign to res variable
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// populate the auth map with the tokens
	creds["token"] = res.Header["X-Auth-Access-Token"][0]
	creds["refresh_token"] = res.Header["X-Auth-Refresh-Token"][0]
	creds["domain"] = res.Header["Domain_uuid"][0]

	//response body
	defer res.Body.Close()
}

func getFMCInfo() {

}

func getDevices() []Device {
	//use our nodelist struct to store the response
	var devicelist DeviceList

	// make the api call
	res := fmcCall(endPoints["devices"])

	error := json.Unmarshal(res, &devicelist)
	if error != nil {
		log.Println(error)
	}

	// build a slice of node items using the node struct
	var x []Device
	for i := range devicelist.Items {

		x = append(x, devicelist.Items[i])

	}

	return x
}

func printTable(devices []Device) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Model", "Software Version", "Is Connected")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, device := range devices {
		tbl.AddRow(device.Name, device.Model, device.SwVersion, device.IsConnected)
	}

	tbl.Print()
}

var art string = `


███████╗███╗   ███╗██████╗ ███████╗██████╗ 
██╔════╝████╗ ████║██╔══██╗██╔════╝██╔══██╗
█████╗  ██╔████╔██║██████╔╝█████╗  ██████╔╝
██╔══╝  ██║╚██╔╝██║██╔══██╗██╔══╝  ██╔══██╗
███████╗██║ ╚═╝ ██║██████╔╝███████╗██║  ██║
╚══════╝╚═╝     ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝ v0.1
                                           
Ember - A Cisco FMC API Client by Steven McNutt CCIE #6495
Github:	srmcnutt/ember
`
