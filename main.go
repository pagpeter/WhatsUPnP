package main

import (
	"flag"
	"fmt"
)

var waitTime int
var verbose bool
var useJson bool
var filename string
var outputMode string

func init() {
	flag.IntVar(&waitTime, "wait", 5, "how long to wait for responses")
	flag.StringVar(&filename, "file", "output.json", "output filename (only if json output is enabled)")
	flag.BoolVar(&useJson, "json", false, "Output in JSON format")
	flag.BoolVar(&verbose, "verbose", false, "verbose output (doesn't affect json)")

	flag.Parse()

	if useJson {
		outputMode = "json"
	} else {
		outputMode = "std"
	}

	if verbose {
		fmt.Println("Wait time:", waitTime)
		fmt.Println("Verbose:", verbose)
		fmt.Println("Filename:", filename)
		fmt.Println("Output mode:", outputMode)
	}

}

func main() {
	devices := make(map[string][]DeviceResponse)
	total := 0

	if verbose {
		fmt.Println("Starting...")
	}

	channel := make(chan string)
	go getAllDevices(waitTime, channel)

	for {
		res := <-channel
		parsed := ParseSSDPPacket(res)
		if parsed.stop {
			break
		}
		total++
		devices[parsed.Server] = append(devices[parsed.Server], parsed)
	}

	if outputMode == "std" {
		OutputToSTD(devices, total)
	} else if outputMode == "json" {
		fmt.Println("Writing to file " + filename)
		OutputToJSON(devices, total)
	}

}
