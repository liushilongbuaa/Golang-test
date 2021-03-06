package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/bluele/gcache"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	json_test()
}

func json_test() {
	type A struct {
		Name []string `json:"name,omitempty"`
		Age  []*int   `json:"age"`
		City *[]int   `json:"city"`
	}
	a := A{}
	bt, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}
func mysql_test() {
	strConn := "%s:%s@tcp(%s:%d)/%s?autocommit=true&parseTime=true&timeout=%dms&loc=Asia%%2FShanghai&tx_isolation='READ-COMMITTED'"
	url := fmt.Sprintf(strConn, "root", "admin", "192.168.178.128", 3306, "cc", 3000)
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("mysql open err: %s\n", err.Error())
		return
	}

	result, err := db.Query("select id,name from vpc where id=? limit 10", "")
	if err != nil {
		fmt.Printf("db.query err: %s\n", err.Error())
	}
	defer result.Close()
	var id, name string

	for result.Next() {
		err = result.Scan(&id, &name)
		if err != nil {
			fmt.Printf("vpc.Scan err: %s\n", err.Error())
		}
	}
	fmt.Printf("id: %s, name: %s", id, name)
}
func flag_test() {
	systemTest := flag.Bool("system-test", false, "Set to true when running system tests")
	flag.Parse()
	fmt.Printf("flag: %v\n", *systemTest)

	fmt.Printf("input params: %v\n", flag.Args())
}
func map_test() {
	A := map[string]interface{}{}
	A["a"] = 123
	A["b"] = "abc"
	if A["a"] != nil {
		fmt.Printf("map_test: not nil\n")
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
			c, err := redis.Dial("tcp", "192.168.217.131:6379", dialDatabase)
			if err != nil {
				fmt.Printf("redis_test: %v\n", err)
				return nil, err
			}

			return c, err
		},
	}
	conn := redis.Get()
	defer conn.Close()

	reply, err := conn.Do("hgetall", "post:2")
	fmt.Printf("err: %v", err)
	fmt.Printf("reply: %s", reply)

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
	cache.SetWithExpire("CCBGW_API_LSL", "haha", time.Minute)
	re, err = cache.Get("CCBGW_API_LSL")
	fmt.Println(re, err)
}
func logic_test() {
	count := func(b int, c [10]int) int {
		var sum int
		for i := 0; i < 10; i++ {
			if c[i] == b {
				sum += 1
			}
		}
		return sum
	}
	var a [10]int
	var i int
	for i = 0; i < 1024*1024; i++ {
		tem := i
		for j := 0; j < 10; j++ {
			a[j] = tem % 4
			tem = tem / 4
		}
		switch a[1] {
		case 0:
			if a[4] != 2 {
				continue
			}
		case 1:
			if a[4] != 3 {
				continue
			}
		case 2:
			if a[4] != 0 {
				continue
			}
		case 3:
			if a[4] != 1 {
				continue
			}
		}
		switch a[2] {
		case 0:
			if (a[5] == a[1] && a[1] == a[3] && a[5] != a[2]) != true {
				continue
			}
		case 1:
			if (a[2] == a[1] && a[1] == a[3] && a[2] != a[5]) != true {
				continue
			}
		case 2:
			if (a[2] == a[5] && a[5] == a[3] && a[2] != a[1]) != true {
				continue
			}
		case 3:
			if (a[2] == a[5] && a[5] == a[1] && a[2] != a[3]) != true {
				continue
			}
		}
		switch a[3] {
		case 0:
			if a[0] != a[4] {
				continue
			}
		case 1:
			if a[1] != a[6] {
				continue
			}
		case 2:
			if a[0] != a[8] {
				continue
			}
		case 3:
			if a[5] != a[9] {
				continue
			}
		}
		switch a[4] {
		case 0:
			if a[4] != a[7] {
				continue
			}
		case 1:
			if a[4] != a[3] {
				continue
			}
		case 2:
			if a[4] != a[8] {
				continue
			}
		case 3:
			if a[4] != a[6] {
				continue
			}
		}
		switch a[5] {
		case 0:
			if (a[7] == a[1] && a[1] == a[3]) == false {
				continue
			}
		case 1:
			if (a[7] == a[0] && a[0] == a[5]) == false {
				continue
			}
		case 2:
			if (a[7] == a[2] && a[2] == a[9]) == false {
				continue
			}
		case 3:
			if (a[7] == a[4] && a[4] == a[8]) == false {
				continue
			}
		}
		C0 := count(0, a)
		C1 := count(1, a)
		C2 := count(2, a)
		C3 := count(3, a)
		fmt.Println(C0+C1+C2+C3, i)
		switch a[6] {
		case 0:
			if C2 > C1 || C2 > C3 || C2 > C0 {
				continue
			}
		case 1:
			if C1 > C0 || C1 > C2 || C1 > C3 {
				continue
			}
		case 2:
			if C0 > C1 || C0 > C2 || C0 > C3 {
				continue
			}
		case 3:
			if C3 > C1 || C3 > C2 || C3 > C0 {
				continue
			}
		}
		switch a[7] {
		case 0:
			if a[6] == a[0]-1 || a[6] == a[0]+1 {
				continue
			}
		case 1:
			if a[4] == a[0]-1 || a[4] == a[0]+1 {
				continue
			}
		case 2:
			if a[1] == a[0]-1 || a[1] == a[0]+1 {
				continue
			}
		case 3:
			if a[9] == a[0]-1 || a[9] == a[0]+1 {
				continue
			}
		}
		switch a[8] {
		case 0:
			if (a[0] == a[5]) == (a[4] == a[5]) {
				continue
			}
		case 1:
			if (a[0] == a[5]) == (a[4] == a[9]) {
				continue
			}
		case 2:
			if (a[0] == a[5]) == (a[4] == a[1]) {
				continue
			}
		case 3:
			if (a[0] == a[5]) == (a[4] == a[8]) {
				continue
			}
		}
		min := C0
		max := C0
		if min > C1 {
			min = C1
		}
		if min > C2 {
			min = C2
		}
		if min > C3 {
			min = C3
		}
		if max < C1 {
			max = C1
		}
		if max < C2 {
			max = C2
		}
		if max < C3 {
			max = C3
		}
		switch a[9] {
		case 0:
			if max-min != 3 {
				continue
			}
		case 1:
			if max-min != 2 {
				continue
			}
		case 2:
			if max-min != 4 {
				continue
			}
		case 3:
			if max-min != 1 {
				continue
			}
		}
		for i := 0; i < 10; i++ {
			switch a[i] {
			case 0:
				fmt.Print("A")
			case 1:
				fmt.Print("B")
			case 2:
				fmt.Print("C")
			case 3:
				fmt.Print("D")
			}
		}
		fmt.Println(" ")
		break
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
