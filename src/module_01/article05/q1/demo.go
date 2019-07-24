package main

import (
	"log"
)

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		log.Println("The block is ", block)
	}
	log.Println("The block is ", block)
	logBlock()
}

func logBlock() {
	log.Println("The block is ", block)
}
