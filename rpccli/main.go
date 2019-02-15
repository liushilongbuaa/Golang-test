package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("check success")
}
func main() {
	pwd, err := os.Getwd()
	checkErr(err)
	i := strings.LastIndex(pwd, "/")
	addr := pwd[:i+1] + "gsock.sock"
	fmt.Println(addr)
	Naddr, err := net.ResolveUnixAddr("unix", addr)
	checkErr(err)

	conn, err := net.DialUnix("unix", nil, Naddr)
	checkErr(err)
	defer conn.Close()
	n, err := conn.Write([]byte("this is cli"))
	checkErr(err)
	fmt.Println(n)
	bt := make([]byte, 100)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		_, err = conn.Read(bt)
		checkErr(err)
		fmt.Println(string(bt))
	}
}
