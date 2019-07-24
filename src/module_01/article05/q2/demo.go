package main

import "fmt"

var container = []string{"zero", "one1", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one2", 2: "two"}
	value, ok := interface{}(container).(map[int]string)
	if ok {
		fmt.Println("value: ", value)
	}
	fmt.Printf("The element is %q.\n", container[1])
}
