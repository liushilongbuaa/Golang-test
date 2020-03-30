package main

import (
	"LslStudy/golang/duotai"
	"Lslstudy/golang/subdir"
	"bytes"
	"container/list"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/bluele/gcache"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	//	jdsync "jd.com/cc/jstack-cc-common/message/sync"
)

var Str = "乾"

type T struct {
	CCC []string `json:"-"`
	Ping
}
type Ping struct {
	SrcIps *[]string
	DstIps []string
}

var INT_MAX = 1<<31 - 1
var INT_MIN = -(1 << 31)

type Raw interface {
	String() string
}

func main() {
	fmt.Println(fmt.Sprintf("%s haha", strconv.Itoa()))
}
func exec_test() {
	f, err := os.OpenFile("./abc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 666)
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd := exec.Command("echo", time.Now().Format(time.RFC3339), "[ac] [awe] [eawf]")
	cmd.Stdout = f
	cmd.Stderr = f
	err = cmd.Run()
	fmt.Println(err)
}
func mem_test() {
	const Len int = 100000
	a := [Len]*string{}
	for i := 0; i < Len; i++ {
		tem := "long long long a string " + strconv.Itoa(i)
		a[i] = &tem
	}
	fmt.Println("mem allocated success.")
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * time.Duration(1))
	}
	for i := 0; i < Len; i++ {
		a[i] = nil
		fmt.Println("release ", i)
		time.Sleep(time.Microsecond * time.Duration(100))
	}
	runtime.GC()
	select {}
}

func time_test() {
	timmer := time.NewTimer(time.Second * time.Duration(10))
	for i := 0; i < 1; i++ {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
	for {
		select {
		case <-timmer.C:
			fmt.Println("timmer")
		default:
			fmt.Println("default")
			time.Sleep(time.Second)
		}
	}

	fmt.Println("finish")
}
func duotai_test() {
	bt, _ := ioutil.ReadFile("a.json")
	amap := map[string]interface{}{}
	json.Unmarshal(bt, &amap)
	for _, v := range duotai.SceneList {
		if mod, ok := amap[v.Id()]; ok {
			bt, _ = json.Marshal(mod)
			json.Unmarshal(bt, v)
		}
		v.Say()
	}
	bt, _ = json.Marshal(duotai.SceneList)
	fmt.Println(string(bt))
	return
}
func swi(s string) Raw {
	return reflect.ValueOf(s)
}
func fn(n, l, r int) int {
	if n == 1 {
		return 1
	}
	if n == l {
		return 1
	}
	if r == 1 {
		return fn(n-1, l-1, r-1)
	}
	if l <= r {
		return fn(n, l+1, r)
	}
	return fn(n, l+1, r) + fn(n, l, r+1)
}
func printQ(a string, ret *[]string, l, r int) {
	if l < 0 || r < 0 {
		return
	}
	if l == r && l == 0 {
		*ret = append(*ret, a)
		return
	}
	if l >= r {
		printQ(a+"(", ret, l-1, r)
	} else {
		printQ(a+"(", ret, l-1, r)
		printQ(a+")", ret, l, r-1)
	}
}
func tcp_test() {
	ln, err := net.Listen("tcp", ":888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		bt := make([]byte, 1024)
		n, err := conn.Read(bt)
		if err != nil {
			fmt.Println(err)
			return
		}
		bt = bt[:n]
		fmt.Println(string(bt))
	}
}
func signal_test() {
	quit := make(chan os.Signal)
	sigs := []os.Signal{}
	for i := 0; i < 16; i++ {
		sigs = append(sigs, syscall.Signal(i))
	}
	signal.Notify(quit, sigs...)
	sig := <-quit
	fmt.Println(sig.String())
	fmt.Println("Go to exit.")
}
func rawmessage_test() {
	//	a := Raw{Type: 111}
	//	from := "vpc-1"
	//	to := "vpc-2"
	//	b := jdsync.ResultVpcPeeing{From: &from, To: &to}
	//	a.Result = b
	//	bt, err := json.Marshal(a)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(string(bt))

	//	re := jdsync.QueryResult{}
	//	err = json.Unmarshal(bt, &re)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(*re.Type)
	//	fmt.Println(string(*re.Result))
}
func rsa_test() *rsa.PrivateKey {
	time1 := time.Now()
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	fmt.Println(time.Since(time1).Nanoseconds())
	time2 := time.Now()
	if err != nil {
		fmt.Println("generateKey error :", err)
	}
	fmt.Println(time.Since(time2).Nanoseconds())
	time3 := time.Now()
	fmt.Println("D ", *privateKey.D)
	fmt.Println(time.Since(time3).Nanoseconds())
	fmt.Println("E ", privateKey.E)
	fmt.Println("N ", *privateKey.N)
	fmt.Println("pub ", privateKey.PublicKey)
	return privateKey
}
func bit_test() {
	a := big.NewInt(123456789012341234)
	b := big.NewInt(123456789012341234)
	var c big.Int
	c.Mul(a, b)
	fmt.Println(a.BitLen(), b.BitLen(), c.BitLen())
	fmt.Println(c.String())
	fmt.Println(c.Bytes())
}
func unsafe_test() {
	t := T{
	///	CCC: "abc",
	}
	l := unsafe.Sizeof(t)
	pb := (*[1024]byte)(unsafe.Pointer(&t))
	fmt.Println("Struct:", t)
	fmt.Println(l, "Bytes:", (*pb)[:24])
	///	fmt.Println([]byte(t.CCC))
}
func slice_test() {
	b := []byte{192, 168, 1, 1, 0, 0, 0, 111, 192, 168, 0, 0, 0, 0, 111}
	a := b[:3]
	c := b[1:4]
	fmt.Println(b, a, c)
	a[0] = 111
	fmt.Println(b, a, c)
}
func ip_test() {
	ipv6cidr := "2001:2:1:110e:0:0:123:ffab/56"
	//ipv6cidr := "192.168.222.1/17"
	ipaddr, network, err := net.ParseCIDR(ipv6cidr)
	fmt.Println(ipaddr, network, err)
	_, size := network.Mask.Size()
	for i := 0; i < size/8; i++ {
		fmt.Printf(" %x", ^network.Mask[i])
	}
	fmt.Println("\n", network.Mask.String())
	fmt.Println("############################")
	fmt.Println(network.Mask.Size())
}
func convolve(u, v []int) []int {
	n := len(u) + len(v) - 1
	ret := make([]int, n)
	for k := 0; k < n; k++ {
		ret[k] = mul(u, v, k)
	}
	return ret
}
func convolve2(u, v []int) (w []int) {
	n := len(u) + len(v) - 1
	w = make([]int, n)

	// 将 w 切分成花费 ~100μs-1ms 用于计算的工作单元
	size := 1 << 22 / n

	wg := new(sync.WaitGroup)
	wg.Add(1 + (n-1)/size)
	for i := 0; i < n && i >= 0; i += size { // 整型溢出后 i < 0
		j := i + size
		if j > n || j < 0 { // 整型溢出后 j < 0
			j = n
		}

		// 这些goroutine共享内存，但是只读
		go func(i, j int) {
			for k := i; k < j; k++ {
				w[k] = mul(u, v, k)
			}
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	return
}
func mul(u, v []int, k int) (res int) {
	a := k
	tem := 0
	if k > len(u)-1 {
		a = len(u) - 1
	}
	if k > len(v)-1 {
		tem = k - len(v) + 1
	}
	for i := tem; i < a; i++ {
		res += u[i] + v[k-i]
	}
	return
}
func MultiProcess_test() {
	var u, v []int
	for i := 0; i < 50000; i++ {
		u = append(u, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}...)
		v = append(v, []int{1, 1, 1, 1, 1}...)
	}
	/*start := time.Now()
	a := Convolve(u, v)
	fmt.Println(time.Since(start), ":     ", a)*/
	start := time.Now()
	a := convolve2(u, v)
	fmt.Println(time.Since(start), ":     ", a)
}
func iota_test() {
	const (
		A = -1
		B = iota
		C
		D
		E
	)
	fmt.Println(A, B, C, D, E)
}
func int_test() {
	fmt.Println(uint(INT_MAX) + 2)
	fmt.Println(int(uint(INT_MAX) + 2))

	fmt.Printf("%x\t%d\n", INT_MAX, INT_MAX)
	fmt.Printf("%x\t%d\n", INT_MIN, INT_MIN)

	var INT32_MAX = int32(^uint32(0) >> 1)
	var INT32_MIN = ^INT_MAX
	fmt.Printf("%x\t%d\n", INT32_MAX, INT32_MAX)
	fmt.Printf("%x\t%d\n", INT32_MIN, INT32_MIN)
}
func fullrune_test() {
	buf := []byte{228, 184, 180}
	fmt.Println(string(buf))
	fmt.Printf("%b  %b  %b\n", buf[0], buf[1], buf[2])
	fmt.Println(utf8.FullRune(buf))
}
func unicode_test() {
	bt, err := ioutil.ReadFile("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("utf-8:")
	for _, v := range bt {
		if v == 0x0d || v == 0x0a {
			break
		}
		fmt.Printf("%x ", v)
	}
}
func rune_test() {
	descRe := regexp.MustCompile(`^(.){0,256}$`)
	str := "123456789012345aaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffffaaaaabbbbbcccccdddddeeeeefffff"
	b := descRe.MatchString(str)
	fmt.Println(b)
	fmt.Println(utf8.FullRune([]byte{0xe1, 0x99}))
	bts := []byte(str)
	bts = append(bts, []byte{0xe1}...)
	b = descRe.Match(bts)
	fmt.Println(b)

	str = "hello, 刘世龙"
	for k, v := range str {
		fmt.Printf("%x %d\n", v, k)
	}
	fmt.Println("\n", len(str))
	for k, v := range []rune(str) {
		fmt.Printf("%x %d\n", v, k)
	}
	fmt.Println(len([]rune(str)))
}
func rutime_test() {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
	subdir.TestSub()
}
func regexp_test() {
	uuidRe := regexp.MustCompile(`^([a-zA-Z0-9-_.]|[\p{Han}])*$`)
	tem := `路上高比例abc-._`
	fmt.Println(uuidRe.Match([]byte(tem)))

	macRe := regexp.MustCompile(`^([0-9A-F]{2}\:){5}[0-9A-F]{2}$`)
	fmt.Println(macRe.MatchString("FA:12:A2:21:22:11"))
}
func json_test() {
	jsonrow := `
	{
		"p1":"abc"
	}
	`
	a := C{}
	err := json.Unmarshal([]byte(jsonrow), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.P1, a.P2)
	bt, err := json.MarshalIndent(a, "", "    ")

	fmt.Println(err)
	fmt.Println(string(bt))
}
func mysql_test() {
	// connection
	strConn := "%s:%s@tcp(%s:%d)/%s?autocommit=true&parseTime=true&timeout=%dms&loc=Asia%%2FShanghai&tx_isolation='READ-COMMITTED'"
	url := fmt.Sprintf(strConn, "root", "admin", "10.226.137.197", 3306, "cc", 3000)
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("mysql open err: %s\n", err.Error())
		return
	}
	var ctx context.Context = context.WithValue(context.Background(), "trace_id", "xxxxxxxx")
	_ = ctx
	// Query 查不到不会报错，raws.next()=false
	sqlstr := "select updated_at from bgw where updated_at > ?;"
	raws, err := db.Query(sqlstr, time.Now().Truncate(time.Second*time.Duration(21100)))
	if err != nil {
		fmt.Println("418", err)
		return
	}
	fmt.Println(time.Now(), "#############", time.Now().Add(time.Second*time.Duration(-60)))
	var a time.Time
	for raws.Next() {
		err = raws.Scan(&a)
		if err != nil {
			fmt.Printf("raw.Scan err: %s\n", err.Error())
			fmt.Println(err == sql.ErrNoRows)
		}
		fmt.Println(a)
	}

}
func flag_test() {
	systemTest := flag.Bool("system-test", false, "Set to true when running system tests")
	flag.Parse()
	fmt.Printf("flag: %v\n", *systemTest)

	fmt.Printf("input params: %v\n", flag.Args())
}
func map_test() {
	a := map[int]int{}
	map_test_1(a)
	fmt.Println(a)
}
func map_test_1(A map[int]int) {
	for i := 2; i < 10; i++ {
		A[i] = i
	}
}
func redis_test() {
	//dialConnectTimeout := redis.DialConnectTimeout(time.Duration(1000) * time.Millisecond)
	//dialReadTimeout := redis.DialReadTimeout(time.Duration(3000) * time.Millisecond)
	//dialWriteTimeout := redis.DialWriteTimeout(time.Duration(3000) * time.Millisecond)
	dialDatabase := redis.DialDatabase(0)
	//password := redis.DialPassword("34fd10676ebdb68014888a0df2b0ab61")

	redis := &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: time.Duration(3000) * time.Millisecond,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "10.226.137.197:6379", dialDatabase)
			if err != nil {
				fmt.Printf("redis_test: %v\n", err)
				return nil, err
			}

			return c, err
		},
	}
	conn := redis.Get()
	defer conn.Close()
	conn.Do("hset", "task", "vpc-111:vpc222", "true")
	reply, err := conn.Do("hget", "task", "vpc-111:vpc222")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("reply: %s\n", reply)
	a, ok := reply.(string)
	fmt.Println(ok)
	fmt.Println(a)
}
func hash_test() {
	sha1.Sum([]byte("abc"))
}
func bitfuc_test() {
	a := 0x17
	b := 0x38
	c := b &^ a
	fmt.Println(c)
}
func gcache_test() {
	cache := gcache.New(1024).LRU().Build()
	re, err := cache.Get("CCBGW_API_LSL")
	fmt.Println(re, err)
	cache.SetWithExpire("CCBGW_API_LSL", "haha", time.Duration(1)*time.Second)
	for i := 0; i < 10; i++ {
		re, err = cache.Get("CCBGW_API_LSL")
		fmt.Println(re, err)
		time.Sleep(time.Millisecond * time.Duration(300))
	}
}
func list_test() {
	l := list.New()
	l.PushBack(time.Now())
	l.PushBack(time.Now().Add(time.Second))
	tem := time.Now().Add(time.Duration(500) * time.Millisecond)

	for i := l.Front(); i != nil; i = i.Next() {
		if i_, ok := i.Value.(time.Time); ok {
			if i_.After(tem) {
				l.InsertBefore(tem, i)
				break
			}
		}
	}
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
func pi_test() {
	const ARRSIZE = 10010
	const DISPCNT = 10000
	var x [ARRSIZE]int
	var z [ARRSIZE]int
	x[1] = 1
	z[1] = 1
	a := 1
	b := 3
	d := 0
	cnt := 0
	pre := 0
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
A:
	for {
		if cnt != 0 {
			flag := true
			for i := 0; i < ARRSIZE; i++ {
				if z[i] != 0 {
					flag = false
				}
			}
			if flag == true {
				break A
			}
		}
		d = 0
		for i := ARRSIZE - 1; i > pre; i-- {
			c := z[i]*a + d
			z[i] = c % 10
			d = c / 10
		}
		d = 0
		for i := 1 + pre; i < ARRSIZE; i++ {
			c := z[i] + d*10
			z[i] = c / b
			d = c % b
		}
		a++
		b += 2
		d = 0
		for i := ARRSIZE - 1; i > pre; i-- {
			c := x[i] + z[i] + d
			x[i] = c % 10
			d = c / 10
		}
		cnt++
		if cnt != 0 && cnt%20 == 0 {
			i := 0
			for i = 0; i < ARRSIZE; i++ {
				if z[i] != 0 {
					break
				}
			}
			if i > pre+20 {
				pre = i - 10
			}

		}
	}
	d = 0
	for i := ARRSIZE - 1; i > 0; i-- {
		c := x[i]*2 + d
		x[i] = c % 10
		d = c / 10
	}
	fmt.Printf("pi compute %d times.\n", cnt)

	for i := 1; i < DISPCNT; i++ {
		fmt.Print(x[i])
		if i%100 == 0 {
			fmt.Println()
		}
		if i == 1 {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
func httpclient_test() {
	input := []byte{}
	req, err := http.NewRequest("POST", "http://10.226.137.196:9698/cc-server", bytes.NewReader(input))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "CcClient")

	client := http.Client{Timeout: 100000 * time.Microsecond}
	resp, err := client.Do(req)
	fmt.Println(resp, err)
}
func http_server_test() {
	handler := &handler{}
	http.Handle("/", handler)
	err := http.ListenAndServe(":801", nil)
	fmt.Println("finished", err)
}
func deferreturn_test() (int, error) {
	fmt.Println("1")
	defer fmt.Println("2")
	return fmt.Println("3")
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
