package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, e := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if e != nil {
		log.Println(e)
		return
	}
	defer f.Close()
	ret, e := f.Seek(2, 0)
	if e != nil {
		log.Println(e)
		return
	}
	fmt.Println("ret=", ret)
	_, e = f.WriteAt([]byte("1"), ret)
	if e != nil {
		log.Println(e)
		return
	}
}
