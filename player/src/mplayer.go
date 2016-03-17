package main 

import ( 
	"fmt"
	"bufio"
	"os"
	"strings"
	
	"library"
	"mp"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands( tokens []string ) {
	if len(tokens) < 2 {
		fmt.Println( "args error!" )
		return
	} 
	
	switch ( tokens[1] ) {
		case "add":
			if len(tokens) == 4 {
				id++
				lib.Add( &library.MusicEntry{ tokens[2], tokens[3] } )
				fmt.Println( "add ", id, ":", tokens[2], ".", tokens[3] )
			}
		case "list":
			for i:=0; i<lib.Len(); i++ {
				e, _ := lib.GetAt(i)
				fmt.Println( i,  ":", e.Name )
			}
		case "remove":
		default:
			fmt.Println( "Unsupport command", tokens[1] )
	}
}

func handlePlayCommand( tokens []string ) {
	music := lib.Find( tokens[1] )
	if music == nil {
		fmt.Println( "find failed ", tokens[1] )
		return
	}
	mp.Play( music.Name, music.Type )
}

func main() {
	lib = library.NewMusicManager()
	r := bufio.NewReader( os.Stdin )
	
	for {
		fmt.Println( "Enter command->" )
		rawLine, _, _ := r.ReadLine()
		line := string( rawLine )
		
		if line == "q" || line == "e" {
			break;
		}
		
		tokens := strings.Split( line, " " )
		
		switch( tokens[0] ) {
			case "lib":
				handleLibCommands( tokens )
			case "play":
				handlePlayCommand( tokens )
			default:
				return;
		}
	}
}