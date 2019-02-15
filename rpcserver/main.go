package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	checkErr(err)
	fmt.Println(addr)

	listener, err := net.Listen("unix", addr)
	checkErr(err)
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		checkErr(err)
		defer conn.Close()
		serve(conn)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
func serve(conn net.Conn) {
	bt := make([]byte, 1024)
	n, err := conn.Read(bt)
	checkErr(err)
	fmt.Println(len(bt), n, string(bt))
	time.Sleep(time.Duration(3))
	conn.Write([]byte("this is server"))
}
