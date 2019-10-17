package main

type A struct {
	Name string
	Age  int
}

type B struct {
	A
	Birthday string
}

type C struct {
	P1 string `josn:"p1"`
	P2 string `josn:"p2"`
}
