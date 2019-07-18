package main

import (
	"fmt"
	"net/rpc"
	"os"
	"runtime"
)

func checkErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line, " ", err.Error())
		os.Exit(2)
	}
}

func main() {
	args := os.Args[1]
	if args == "tcp" {
		fmt.Println("tcp")
		client, err := rpc.Dial("tcp", "192.168.217.131:1234")
		checkErr(err)
		defer client.Close()
		var reply string
		err = client.Call("Hello.Haha", args, &reply)
		checkErr(err)
		fmt.Println("recieve from tcp :1234    ", reply)
	} else {
		client1, err := rpc.Dial("unix", "/tmp/haha")
		checkErr(err)
		defer client1.Close()
		var reply1 string
		err = client1.Call("Hello.Haha", args, &reply1)
		checkErr(err)
		fmt.Println("recieve from unix:/tmp/haha    ", reply1)
	}
}
