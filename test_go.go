package main

import (
	"fmt"
)

func main() {
	ch := make( chan int, 1024 )
	for {
		select {
			case ch <- 1:
			case ch <- 2:
		}
		
		//c := <- ch
		//fmt.Println( "Receive value:", c )
		
		for i := range ch {
			fmt.Println( "Receive value:", i )
		}
	}
}