package main

import (
	"strings"
	"testing"
)

func TestBigInt(t *testing.T) {
	strs := []string{"123456789", "-12345", "001232382345", "1291-12"}
	for _, str := range strs {
		test(t, str)
	}
}

func test(t *testing.T, str string) {
	a := BigInt(str)
	if a == nil {
		t.Logf("wrong num: %s", str)
		return
	}
	if str != a.String() {
		if strings.HasSuffix(str, a.String()) {
			t.Logf("str: %s, BigInt: %s", str, a.String())
		} else {
			t.Error(str, a.value)
		}
	}
}

func TestLess(t *testing.T) {
	a := map[*Int]*Int{
		BigInt("123151"):  BigInt("321456"),
		BigInt("1231151"): BigInt("321456"),
		BigInt("-123151"): BigInt("321456"),
		BigInt("-123151"): BigInt("-321456"),
		BigInt("123151"):  BigInt("123456"),
	}
	for k, v := range a {
		if k.Less(v) {
			t.Logf("%s  < %s", k, v)
		} else {
			t.Logf("%s >= %s", k, v)
		}
	}
}
