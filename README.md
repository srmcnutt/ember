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



## Usage:
- (optional) Set the environment variables
- Run the program

Binaries for windows, mac, and linux are available under the binaries folder.

## TODOs
- unit tests
- additional features (shoot me some ideas or even better make a pull request)






