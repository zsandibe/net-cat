package pkg

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func EnterName(conn net.Conn) string {
	takeName, _, err := bufio.NewReader(conn).ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	name := string(takeName)
	for {
		if name == "" {
			conn.Write([]byte("[ENTER YOUR NAME]:"))
			takeName, _, err = bufio.NewReader(conn).ReadLine()
			if err != nil {
				log.Fatalln(err)
			}
			name = string(takeName)
		} else {
			break
		}
	}
	fmt.Println(name + " connected")
	return name
}
