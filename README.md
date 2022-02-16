# WhatsUp(np) - Discover all UPnP services in a network

WhatsUpnp is a pure go CLI app to track down all UPnP devices on your network. It doesn't have any third-party dependencies.

## Install

The easiest way to install WhatsUpnp is with the go install command.

`go install github.com/wwhtrbbtt/WhatsUPnP@latest`

## Usage

```
Usage of WhatsUPnP:
  -file string
        output filename (only if json output is enabled) (default "output.json")
  -json
        Output in JSON format
  -verbose
        verbose output (doesn't affect json)
  -wait int
        how long to wait for responses (default 5)
  -h
        Display this help message
```

Usage example:

```sh
$ WhatsUPnP -v

Total responses: 111
Unique devices: 5

...
```
