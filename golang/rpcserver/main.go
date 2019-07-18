package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"runtime"
	"strconv"
	"text/template"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func init() {
	dail := func() (redis.Conn, error) {
		return redis.Dial("tcp", "192.168.217.131:6379")
	}
	redisPool = redis.NewPool(dail, 10)
	if redisPool == nil {
		os.Exit(1)
	}
}
func checkErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line, " ", err.Error())
		os.Exit(2)
	}
}

func main() {
	rpcServer()
	return
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	checkErr(http.ListenAndServe(":180", nil))
}
func login(w http.ResponseWriter, r *http.Request) {
	checkErr(r.ParseForm())
	fmt.Println("login: ", r.Method)
	start := time.Now()
	select {
	case <-r.Context().Done():
		fmt.Println(time.Since(start))
		return
	}
	if r.Method == "GET" {
		tem, err := template.ParseFiles("login.html")
		checkErr(err)
		checkErr(tem.Execute(w, nil))
	} else {
		name := r.FormValue("username")
		password := r.FormValue("password")
		interest := r.FormValue("interest")
		fmt.Println(name, password, interest)
		conn := redisPool.Get()
		strmap, err := redis.StringMap(conn.Do("HMGET", "user", "name", "password"))
		checkErr(err)
		if strmap[name] == password {
			bt, err := ioutil.ReadFile("Lighthouse.jpg")
			checkErr(err)
			w.Write(bt)
		} else {
			w.Write([]byte("login fail."))
		}
	}
}
func upload(w http.ResponseWriter, r *http.Request) {
	checkErr(r.ParseForm())
	fmt.Println("upload: ", r.Method)
	if r.Method == "GET" {
		tem, err := template.ParseFiles("upload.html")
		checkErr(err)
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(time.Now().Unix(), 10))
		aoken := fmt.Sprintf("%x", h.Sum(nil))
		checkErr(tem.Execute(w, aoken))
	} else {
		checkErr(r.ParseMultipartForm(32 << 20))
		file, handler, err := r.FormFile("uploadfile")
		checkErr(err)
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		checkErr(err)
		defer f.Close()
		io.Copy(f, file)
	}
}
func rpcServer() {
	server := rpc.NewServer()
	A := &Hello{}
	server.Register(A)

	l, err := net.Listen("tcp", ":1234")
	checkErr(err)
	defer l.Close()
	go func() {
		for {
			c, err := l.Accept()
			checkErr(err)
			server.ServeConn(c)
		}
	}()

	//	l2, err := net.Listen("unix", "/tmp/haha")
	//	checkErr(err)
	//	defer l2.Close()
	//	go func() {
	//		for {
	//			c, err := l2.Accept()
	//			checkErr(err)
	//			server.ServeConn(c)
	//		}
	//	}()
	select {}
}

type Hello struct{}

func (h *Hello) Haha(args string, reply *string) error {
	tem := "this is Hello.Haha " + args
	*reply = tem
	return nil
}
