package main

import "fmt"
import "flag"
import "os"
import "io"
import "bufio"
import "strconv"
import "algorithms/bubblesort"
import "algorithms/qsort"

var infile *string = flag.String( "i", "infile", "" );
var outfile *string = flag.String( "o", "outfile", "" );
var algorithm *string = flag.String( "a", "algorithm", "" );

func readValue( infile string ) ( values []int, err error ) {
	file, err := os.Open( infile )
	if err != nil {
		fmt.Println( "Failed to open file", infile )
		return
	}
	
	defer file.Close()
	
	br := bufio.NewReader( file )
	values = make( []int, 0 )
	for {
		line, isPrefix, err1 := br.ReadLine();
		
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		
		if isPrefix {
			fmt.Println( "The line is too long, seems unexpected!" )
			return
		}
		
		str := string(line)
		
		value1, err1 := strconv.Atoi( str )
		if err1 != nil {
			err = err1
			break
		}
		
		values = append( values, value1 )
	}
	
	return
}

func writeValue( outfile string, values []int ) error {
	file, err := os.Create( outfile )
	if err != nil {
		fmt.Println( "Failed to open file ", outfile )
		return nil
	}
	
	defer file.Close()
	
	for _, value := range values {
		str := strconv.Itoa( value )
		file.WriteString( str + "\r\n" )
	}
	
	return nil
}

func main() {
	flag.Parse();
	
	if ( infile != nil ) {
		fmt.Println( "infile =", *infile, ", outfile =",*outfile, ", algorithm =",*algorithm );
	}
	
	values, err := readValue( *infile )
	if err != nil {
		fmt.Println( err );
		return
	}
	
	switch *algorithm {
		case "bubblesort":
			bubblesort.BubbleSort( values )
		case "qsort":
			qsort.Qsort( values )
		default:
		
	}
	
	fmt.Println( "Write to file:", values );
	err = writeValue( *outfile, values )
	if err != nil {
		fmt.Println( err )
	}
}