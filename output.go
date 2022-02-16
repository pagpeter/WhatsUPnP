package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// Output to stdout:
func OutputToSTD(devices map[string][]DeviceResponse, total int) {
	fmt.Println("Total responses:", total)
	fmt.Println("Unique devices:", len(devices), "\n")

	for server, responses := range devices {
		otherHeaders := make(map[string]string)

		verboseText := ""

		for i, response := range responses {
			if verbose {
				verboseText += "  - Location: " + response.Location + "\n"
				verboseText += "  - USN: " + response.USN + "\n"
				verboseText += "  - ST: " + response.ST + "\n"
				if i < len(responses)-1 {
					verboseText += "    -----\n"
				}
			}
			for k, _ := range response.OtherHeaders {
				otherHeaders[k] = ""
			}
		}

		// Extract the IP from the first response
		IP := strings.Split(strings.Split(responses[0].Location, "://")[1], ":")[0]
		fmt.Println("\nServer:", server)
		fmt.Println("-", "Total responses:", len(responses))
		fmt.Println("-", "IP:", IP)
		if len(otherHeaders) > 0 {
			fmt.Println("-", "Custom headers:")
			for k, _ := range otherHeaders {
				fmt.Println("  -", k)
			}
		}
		if verbose {
			fmt.Println("- Responses:")
			fmt.Println(verboseText)
		}

	}
}

func OutputToJSON(devices map[string][]DeviceResponse, total int) {
	j, _ := json.MarshalIndent(devices, "", "  ")
	ioutil.WriteFile(filename, j, 0644)
}
