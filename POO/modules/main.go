package main

import (
	"github.com/donvito/hellomod"
	alias "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	alias.SayHello("diego")
}
