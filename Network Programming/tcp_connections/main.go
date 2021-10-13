package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func createConn(addr *net.TCPAddr) {
	defer log.Println("-> Closing")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln("-> Connection:", err)
	}
	log.Println("-> Connection to", addr)
	r := bufio.NewReader(os.Stdin)
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
	}
}

func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	time.Sleep(time.Second / 2)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println("<- Message error:", err)
			continue
		}
		switch msg = strings.TrimSpace(msg); msg {
		case `\q`:
			log.Println("Exiting...")
			if err := conn.Close(); err != nil {
				log.Println("<- Close:", err)
			}
			time.Sleep(time.Second / 2)
			return
		case `\x`:
			log.Println("<- Special message `\\x` received!")
		default:
			log.Println("<- Message Received:", msg)
		}
	}
}

/*func main() {
				if len(os.Args) != 2 {
				log.Fatalln("Please specify an address.")
				}
				addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
				if err != nil {
				log.Fatalln("Invalid address:", os.Args[1], err)
				}
				listener, err := net.ListenTCP("tcp", addr)
				if err != nil {
				log.Fatalln("Listener:", os.Args[1], err)
				}
				for {
				time.Sleep(time.Millisecond * 100)
				conn, err := listener.AcceptTCP()
				if err != nil {
				log.Fatalln("<- Accept:", os.Args[1], err)
				}
				go handleConn(conn)
				}
}

func main() {
					if len(os.Args) != 2 {
					log.Fatalln("Please specify an address.")
					}
					addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
					if err != nil {
					log.Fatalln("Invalid address:", os.Args[1], err)
					}
					createConn(addr)
}
*/
