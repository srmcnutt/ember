Fire-T is currently a super basic tool written in Go that will query Cisco Firepower Management Center for a list of devices, then output a summarized list.

Requirements:
To run it requires 3 environment variables to be set:
- FMC_HOST
- FMC_USER
- FMC_PASSWORD

Also the REST API needs to be enabled on FMC

Usage:
Set the environment variables
Run the program

Binaries for windows, mac, and linux are available under the *shocked face* binaries folder.

TODOs
In the future if I have time to work on it, I'll make it a lot nicer and allow the creds to be entered interactively, output the device list as a table, add other functions, etc.





