package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"scribe"
	"strings"
)

const DELIMITER  = "::"

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	protocolOpt := flag.String("protocol", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", true, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	flag.Parse()

	transport, protocol := obtainConnection(protocolOpt, buffered, framed)

	if client, err := runClient(transport, protocol, *addr, *secure); err == nil {
		fmt.Printf("Connected to %v \n", *addr)
		logic(client)
	} else {
		fmt.Println("error running client:", err)
	}
}

func logic(client *scribe.ScribeClient) {
	reader := bufio.NewReader(os.Stdin)

	bulk := make([]*scribe.LogEntry, 0, 10)

	for {
		line := readLine(reader)
		if line == "SEND" {
			if _, err := client.Log(bulk); err == nil{
				fmt.Printf("%v messages sent \n", len(bulk))
				bulk = bulk[:0]
			} else {
				fmt.Println(err)
			}
		} else if (line == "EXIT") {
			os.Exit(0)
		} else if line == "CLEAR" {
			bulk = bulk[:0]
		} else if strings.Contains(line, DELIMITER) {
			log := composeLog(line)
			bulk = append(bulk, log)
			fmt.Printf("%v added to batch \n", log)
		} else {
			fmt.Println("Please use one of: CLEAR, EXIT or SEND")
		}
	}
}

func composeLog(plain string) *scribe.LogEntry {
	splitted := strings.Split(plain, DELIMITER)
	return &scribe.LogEntry{
		Category: splitted[0],
		Message:  splitted[1],
	}
}

func readLine(r *bufio.Reader) string {
	text, _ := r.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}