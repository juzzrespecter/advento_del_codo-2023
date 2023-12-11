package main

import (
       "os"
       "io"
       "bufio"
       "fmt"
)

func check_err(e error) {
     if e != nil {
     	panic(e)
     }
}

type Coord struct {
     i int;
     j int;
}


func main() {
     input, e := os.Open("./input")

     check_err(e)
     rd := bufio.NewReader(input)

     i, j, jdy := 0, 0, 0
     coords := make([]Coord, 0)
     spaceMap := make([]string, 0)
     empty := true
     for {
     	 line, _, e := rd.ReadLine()

	 if e == io.EOF {
	    break 
	 }
	 for idx, c := range string(line) {
	     if c == '#' {
	     	empty = false
		coords = append(coords, Coord{
		       i: idx,
		       j: jdy,
		       })
	     }
	 }
	 if empty == true {
	    jdy++
	 }
	 spaceMap = append(spaceMap, string(line))
	 empty = true
	 i++
	 jdy++
     }
     fmt.Printf("coords: %+v\n", coords)
}