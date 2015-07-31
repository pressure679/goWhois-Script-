package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"net"
)

func main() {

	// init
	server := "whois.verisign-grs.com:43"
	// var search net.IP = []byte{31, 13, 93, 145}
	ipaddr, err := net.ResolveIPAddr("ip4", "31.13.93.145")
	fmt.Println(ipaddr)

	// connect
	conn, err := net.Dial("tcp", server)

	// connection error
	if err != nil {
		log.Fatal(err)
	}

	// make the connection
	fmt.Fprintf(conn, "%s\r\n", ipaddr)
	scanner := bufio.NewScanner(conn)

	// response
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '%' {
			fields := strings.Split(line, ": ") // k,v split
			fmt.Println(fields)
		}
	}

	// response error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}