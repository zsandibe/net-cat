package pkg

import (
	"fmt"
	"io/ioutil"
	"net"
)

func PrintLogo(conn net.Conn) {
	file, err := ioutil.ReadFile("pig.txt")
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(file)
}
