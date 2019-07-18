package duotai

import (
	"fmt"
)

var SceneList []A

func init() {
	SceneList = []A{&Ping{}, &Acl{}}
}

type A interface {
	Say()
	Id() string
}

type Ping struct {
	Config  string
	Context string
}

func (p *Ping) Id() string {
	return "Ping"
}
func (p *Ping) Say() {
	fmt.Printf("this is %s: %v\n", p.Id(), *p)
}

type Acl struct {
	Config  string
	Context string
}

func (p *Acl) Id() string {
	return "Acl"
}
func (p *Acl) Say() {
	fmt.Printf("this is %s: %v\n", p.Id(), *p)
}
