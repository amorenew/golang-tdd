package main

import "fmt"

func main(){
	fmt.Println(Hello("miro"));
}

const helloPrefix = "Hello"
//Hello message for testing
func Hello(name string) string {
	return helloPrefix+" " + name
}