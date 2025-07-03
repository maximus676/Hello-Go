package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(ModifySpaces("hello world", "dash"))       		//hello-world
	fmt.Println(ModifySpaces("hello world", "underscore")) 		//hello_world
	fmt.Println(ModifySpaces("hello world", "unknown"))    		//hello*world
	fmt.Println(ModifySpaces("hello world", ""))          		//hello*world
	fmt.Println(ModifySpaces("hello world", "asvbawdawd"))		//hello world

}

func ModifySpaces(s, mode string) string {
	if mode == "dash" {
		return strings.ReplaceAll(s, " ", "-")
	} else if mode == "underscore" {
		return strings.ReplaceAll(s, " ", "_")
	} else if mode == "unknown" || mode == "" {
		return strings.ReplaceAll(s, " ", "*")
	}
	return s
}
