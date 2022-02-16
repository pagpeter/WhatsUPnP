package main

import (
	"log"
	"net"
	"time"
)

const message = "M-SEARCH * HTTP/1.1\r\n" +
	"HOST: 239.255.255.250:1900\r\n" +
	"MAN: \"ssdp:discover\"\r\n" +
	"MX: 1\r\n" +
	"ST: ssdp:all\r\n" +
	"\r\n"

func getAllDevices(timeout int, channel chan string) []string {
	var responses []string

	addr := &net.UDPAddr{IP: net.IPv4(239, 255, 255, 250), Port: 1900}
	ourAddr, _ := net.ResolveUDPAddr("udp", "192.168.178.30:10000")

	conn, err := net.ListenUDP("udp", ourAddr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	conn.WriteToUDP([]byte(message), addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			buf := make([]byte, 2048)
			n, _, _ := conn.ReadFromUDP(buf)
			msg := string(buf[:n])
			responses = append(responses, msg)
			channel <- msg
		}
	}()
	time.Sleep(time.Duration(timeout) * time.Second)
	conn.Close()
	return responses
}
