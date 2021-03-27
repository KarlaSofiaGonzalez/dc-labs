// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
	"fmt"
)

type Input struct {
	port     int
	timezone string
}

func handleConn(c net.Conn, timeZ string) {
	defer c.Close()
	locat, _ := time.LoadLocation(timeZ)
	for {
		/print time in locat/
		time_now := time.Now().In(locat).Format("15:04:05\n")
		response := timeZ + " " + time_now
		_, err := io.WriteString(c, response)

		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {

	var input Input
	var host string

	input = manageInput()
	host = "0.0.0.0:" + fmt.Sprintf("%d", input.port)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(host)
	for {
		conn, err := listener.Accept()
		log.Print(conn)
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, input.timezone) 
	}
}

func manageInput() Input {

	var input Input
	var tp *int

	input.timezone = os.Getenv("TZ")

	tp = flag.Int("port", 9000, "port number.")
	flag.Parse()
	input.port = *tp

	return input

}
