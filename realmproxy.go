package main

import (
	"bytes"
	"fmt"
	"strings"
)

func login() {

	fmt.Print("Hello world\n")

	p := Server{
		Addr:   "127.0.0.1:478",
		Target: "34.251.172.139:443",
		ModifyResponse: func(b *[]byte, id string) {
			packet := string(*b)
			//fmt.Println("[login] WebSocket->client: " + packet)
			if strings.HasPrefix(packet, "AXK3413389?ag7") {
				token := packet[14:]
				fmt.Println("bibly packet transform!,  token=" + token)
				by := bytes.NewBufferString("AYK127.0.0.1:5555;" + token)
				*b = by.Bytes()
				//fmt.Println(*b)
			}
		},
	}

	err := p.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}