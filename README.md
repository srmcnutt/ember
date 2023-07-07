Ember is a utility written in Go that will query Cisco Firepower Management Center for a list of devices, then output a summarized list.

## Requirements:
REST API needs to be enabled on Firewall Management Center.
   1. Navigate to System>Configuration>REST API Preferences>Enable REST API.
   2. Check the "Enable REST API" checkbox.
   3. Click "Save". A "Save Successful" dialog will display when the REST API is enabled.

Embers checks for the following environment variables on startup:
- FMC_HOST
- FMC_USER
- FMC_PASSWORD

If any of them are not present, Ember will interactively prompt you for the information.

                                    ** note **
The library used to hide the password for interactive entry is not compatible with windows. 
For windows set the FMC_Password environment variable for the time being.  Here's an example of how to do it in powershell:

```
$env:FMC_PASSWORD="mypassword123"
$env:FMC_HOST="10.1.1.1"
$env:FMC_USER="admin"
```



## Usage:
- (optional) Set the environment variables
- Run the program

Binaries for windows, mac, and linux are available under the binaries folder.

## TODOs
- unit tests
- additional features (shoot me some ideas or even better make a pull request)
- allow interactive password entry on windows






