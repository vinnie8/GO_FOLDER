package main

import (
  //  "io"
    "fmt"
   // "log"
    "strings"
)

type command struct{
    name string
    mp int	//Matches Played
    w int	//Matches Won
    d int	//Matches Drawn (Tied)
    l int	//Matches Lost
    p int	//Points
}


func main() {
	var(
        input string
    ) 
    
    input="Allegoric Alaskians;Blithering Badgers;win;Devastating Donkeys;Courageous Californians;draw;Devastating Donkeys;Allegoric Alaskians;win;Courageous Californians;Blithering Badgers;loss;Blithering Badgers;Devastating Donkeys;loss;Allegoric Alaskians;Courageous Californians;win"
	//working section
    //loading command names
    splitted:=strings.Split(string(input), ";")
    commandsMap := make(map[string]command)

    
    //creating commandslist here
    for _,v:=range splitted{
        //if there's no such command here - add new
        if _, ok := commandsMap[v]; !ok{
    		commandsMap[v]=createCommand(v)
		}	
    }
	
	for k,_:=range commandsMap{
       	fmt.Printf("%s\n", commandsMap[k].name) 
    }
	//fmt.Printf("%s", splitted[0])
  }

//creates command
func createCommand(s string) command{
    c:=command{name: s}
    return c
}