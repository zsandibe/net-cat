package internal

func Hub(ch <-chan Message) {
	for {
		msg := <-ch
		for name, conn := range Allconn {
			if conn == msg.from {
				continue
			}

			if msg.info == "" {
				conn.Write([]byte(msg.body))
				bname := "[" + name + "]" + ":"
				conn.Write([]byte(bname))
			} else {
				conn.Write([]byte(msg.info))
				aname := msg.body + "[" + name + "]" + ":"
				conn.Write([]byte(aname))
			}
		}
		msg.info = ""
	}
}
