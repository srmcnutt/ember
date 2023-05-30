package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// store credentials in a map for easy retrieval
var creds = make(map[string]string)

// store api endpoints in a map for easy retrieval
var endPoints = make(map[string]string)

func main() {
	creds = getEnv()
	banner()
	initEndpoints()
	getAuthToken(endPoints["auth"])
	// need to init a second time now that we have the uuid for the domain
	initEndpoints()
	devices := getDevices()
	printTable(devices)
	fmt.Println("\n Total number of sensors: ", len(devices))
}

// read in environment vars to connect to FMC
func getEnv() map[string]string {
	creds["fmc_host"] = os.Getenv("FMC_HOST")
	creds["fmc_user"] = os.Getenv("FMC_USER")
	creds["fmc_password"] = os.Getenv("FMC_PASSWORD")

	if creds["fmc_host"] == "" {
		fmt.Println("FMC_HOST Environment Variable missing!")
		os.Exit(1)
	}

	if creds["fmc_user"] == "" {
		fmt.Println("FMC_USER Environment Variable missing!")
		os.Exit(1)
	}

	if creds["fmc_password"] == "" {
		fmt.Println("FMC_PASSWORD Environment Variable missing!")
		os.Exit(1)
	}
	return creds
}

// print banner
func banner() {
	fmt.Println(art)
	return
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
	req.Header = http.Header{
		"Content-Type":        {"application/json"},
		"Accept":              {"application/json"},
		"X-auth-access-token": {creds["token"]},
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
  { 
    Twitter:  @densem0de,
    Github:   srmcnutt,
    mastodon: @densemode@infosec.exchange,
  }
`
