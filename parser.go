package main

import (
	"log"
	"strings"
)

// https://web.archive.org/web/20151107123618/http://upnp.org/specs/arch/UPnP-arch-DeviceArchitecture-v2.0.pdf
// "Device Headers"
type DeviceResponse struct {
	// Should always be 200 (OK).
	// These are normal http status codes
	StatusCode string

	CacheControl string

	// ST header
	// Can be the same as USN
	// ServiceType / SearchTarget
	ST string `json:"ST,omitempty"`

	// USN header
	// Unique Service Name
	USN string `json:"USN,omitempty"`

	// EXT header
	// Required for backwards compatibility with UPnP 1.0. (Header field name only; no field value.)
	EXT string `json:"EXT,omitempty"`

	// MX Header
	// Specifies the seconds that responses should be delayed
	MX string `json:"MX,omitempty"`

	// Location Header
	// URI (http/https)
	Location string `json:"location,omitempty"`

	// Server header
	// Example: "unix/5.1 UPnP/2.0 MyProduct/1.0"
	Server string `json:"server,omitempty"`

	// BOOTID.UPNP.ORG
	// Example: "1"
	BootID string `json:"bootid,omitempty"`

	// CONFIGID.UPNP.ORG
	// Example: "1"
	ConfigID string `json:"configid,omitempty"`

	// SECURELOCATION.UPNP.ORG
	// Example: "https://.../description.xml"
	SecureLocation string `json:"secureLocation,omitempty"`

	// Custom headers
	OtherHeaders map[string]string `json:"otherHeaders,omitempty"`

	// small hack
	stop bool `json:"-"`
}

func ParseSSDPPacket(data string) DeviceResponse {
	if data == "" {
		return DeviceResponse{stop: true}
	}
	lines := strings.Split(data, "\r\n")
	response := DeviceResponse{}

	if len(lines) < 1 {
		log.Println("Invalid SSDP packet", data)
		return response
	}

	// First line is the status code
	parts := strings.Split(lines[0], " ")
	if len(parts) < 2 {
		log.Println("Invalid SSDP packet", data)
		return response
	}
	response.StatusCode = strings.Split(lines[0], " ")[1]

	for i, line := range lines {
		// skip the first line
		if i == 0 || line == "" {
			continue
		}

		key := strings.Split(line, ":")[0]
		parts := strings.Split(line, ": ")
		if len(parts) < 1 {
			log.Println("Invalid SSDP packet", data)
			return response
		}
		value := strings.Join(parts[1:], ": ")

		switch key {
		case "ST":
			response.ST = value
		case "USN":
			response.USN = value
		case "EXT":
			response.EXT = value
		case "MX":
			response.MX = value
		case "LOCATION":
			response.Location = value
		case "CACHE-CONTROL":
			response.CacheControl = value
		case "SERVER":
			response.Server = value
		case "BOOTID.UPNP.ORG":
			response.BootID = value
		case "CONFIGID.UPNP.ORG":
			response.ConfigID = value
		case "SECURELOCATION.UPNP.ORG":
			response.SecureLocation = value
		default:
			if response.OtherHeaders == nil {
				response.OtherHeaders = map[string]string{}
			}
			response.OtherHeaders[key] = value
		}
	}
	return response
}
