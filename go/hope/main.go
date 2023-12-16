package main

import (
	"fmt"
)

func main() {
	var p int
	var q int

	fmt.Scanf("%d %d", &p, &q)

//	if q < p {
//		fmt.Println("Error")
//		return
//	}
	var message string
	for i := 1; i <= q; i++ {
		var modulo int = i % p;
		if modulo == 0 {
			message = message + "Hope ";
			fmt.Println(message);
		} else {
			fmt.Println(i);
		}
	}
}
