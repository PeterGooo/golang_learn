package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "everyone", "Your name, fun")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)

}
