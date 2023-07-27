package internal

import (
	"bufio"
	"fmt"
	"net"
	"time"
	p "zsandibe/pkg"
)

type Message struct {
	from net.Conn
	body string
	info string
}

var (
	Data        []string
	Connections int
	Allconn     map[string]net.Conn
)

func init() {
	Allconn = make(map[string]net.Conn)
}

func HandleConnection(conn net.Conn, ch1 chan<- Message) {
	defer conn.Close()

	p.PrintLogo(conn)

	name := p.EnterName(conn)
	Allconn[name] = conn

	for _, message := range Data {
		conn.Write([]byte(message))
		conn.Write([]byte("\n"))
	}

	onetime := time.Now().Format("2006-01-02 15:04:05")
	connMessage := Message{info: "\n" + name + " has joined our chat...\n", body: "[" + onetime + "]", from: conn}
	ch1 <- connMessage

	for {
		time := time.Now().Format("2006-01-02 15:04:05")
		terminal := "[" + time + "]" + "[" + name + "]" + ":"
		conn.Write([]byte(terminal))
		msg, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			Connections--
			fmt.Println(name + " disconnected")
			connMessage := Message{info: "\n" + name + " has left our chat...\n", body: "[" + time + "]", from: conn}
			ch1 <- connMessage
			break
		}

		var connMessage Message
		if isPrintable(string(msg)) {
			connMessage = Message{body: "\n" + terminal + string(msg) + "\n" + "[" + time + "]", from: conn}
			ch1 <- connMessage
			Data = append(Data, terminal+string(msg))
		}
		// if string(msg) != "" {
		// 	connMessage = Message{body: "\n" + terminal + string(msg) + "\n" + "[" + time + "]", from: conn}
		// 	ch1 <- connMessage
		// 	Data = append(Data, terminal+string(msg))
		// }
	}
}

func isPrintable(message string) bool {
	printableFlag := false
	for _, char := range message {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			printableFlag = true
			if char < 32 || char > 126 {
				return false
			}
		}
	}
	return printableFlag
}
