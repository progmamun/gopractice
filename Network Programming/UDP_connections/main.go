package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address.")
	}
	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalln("-> Connection:", err)
	}
	log.Println("-> Connection to", addr)
	r := bufio.NewReader(os.Stdin)
	b := make([]byte, 1024)
	for {
		fmt.Print("# ")
		msg, err := r.ReadBytes('\n')
		if err != nil {
			log.Println("-> Message error:", err)
		}
		if _, err := conn.Write(msg); err != nil {
			log.Println("-> Connection:", err)
			return
		}
		n, err := conn.Read(b)
		if err != nil {
			log.Println("<- Receive error:", err)
		}
		msg = bytes.TrimSpace(b[:n])
		log.Printf("<- %q", msg)
	}
}

/*func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address.")
	}
	addr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatalln("Invalid address:", os.Args[1], err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln("Listener:", os.Args[1], err)
	}
	b := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Println("<-", addr, "Message error:", err)
			continue
		}
		msg := bytes.TrimSpace(b[:n])
		log.Printf("<- %q from %s", msg, addr)
		for i, l := 0, len(msg); i < l/2; i++ {
			msg[i], msg[l-1-i] = msg[l-1-i], msg[i]
		}
		msg = append(msg, '\n')
		if _, err := conn.WriteTo(b[:n], addr); err != nil {
			log.Println("->", addr, "Send error:", err)
		}
	}
}
*/
