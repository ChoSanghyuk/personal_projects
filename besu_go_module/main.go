package main

import "fmt"

var data string = "0x0e37008a%064x"

func main() {

	fmt.Printf(fmt.Sprintf(data, 11))

}
