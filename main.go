package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/StevenZack/demo/network/tcp"
)

func main() {
	offset, e := strconv.ParseUint("1110110111100000", 2, 16)
	if e != nil {
		log.Println(e)
		return
	}
	fmt.Printf("%b\n", offset)
	m := tcp.Message{Flags: uint16(offset)}
	fmt.Printf("%v\n", m.URG())
}
