package main

type BoxHeader struct {
	Size uint32
	Type string
}

type Box struct {
	Head BoxHeader
	Data []byte
}
