package main

import "fmt"
import "os"
import "simplemath"
import "strconv"

func Usage() {
	fmt.Println( "Usage: calc command [argments]..." )
	fmt.Println( "\nThe commands are add and sqrt" )
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	
	switch args[1] {
		case "add":
			arg1, err1 := strconv.Atoi(args[2])
			arg2, err2 := strconv.Atoi(args[3])
			if  err1 != nil || err2 != nil {
				fmt.Println( "err" )
				return
			}
			ret := simplemath.Add( arg1, arg2 )
			fmt.Println( "Return ", ret )
		case "sqrt":
		default:
			Usage();
	}
}