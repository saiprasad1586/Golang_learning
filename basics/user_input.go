package main

import (
	"bufio" //user input
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var mystring string
	// fmt.Scanln(&mystring)
	// fmt.Println(mystring)

	//errors in fmt

	// var name string = "Saiprasad"

	// var a,_=fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your Full name ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating ")

	myrating, _ := reader.ReadString('\n')

	num_rating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)

	fmt.Println(num_rating + 2)

}
