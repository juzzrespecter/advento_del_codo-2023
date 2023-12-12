package main

import (
"fmt"
"os"
)

func check_err(e error) {
     if e != nil {
     	panic(e)
     }
}

type Page struct {
     val string,
     n int,
}

type Node struct {
     paths []string,
     ids []int,
}


func main() {
     input, e := os.Open("./input")

     check_err(e)
     rd := bufio.NewReader(input)

     i := 0
     dict := make([]Page, 0)
     table := make([]Node, 0)
     movs, _, e := rd.ReadLine()
     start, end := 0, 0
     for {
     	 line, _, e := rd.ReadLine()

	 if line == "" {
	    continue
	 }
	 if e == io.EOF {
	    break 
	 }
	 index, pair := line.split()...
	 
	 if index == "AAA" {
	    start = i
	 }
	 if index == "ZZZ" {
	    end = i
	 }
	 dict = append(dict, Page{
	      val: index.trim(),
	      n: i
	      })
	 table = append(table, Node{
	       paths: pair,
	       idx: []
	       })
	 i++
     }
     for idx, node := range table {
     	 done := 0
     	 for jdx, page := range dict {
	     if page.val == node.paths[0] {
	     	node.ids[0] = page.val
		done++
	     }
	     if page.val == node.paths[1] {
	     	node.ids[1] = page.val
		done++
	     }
	     if done == 2 {
	     	break
	     }
	 }
     }

     pos := start
     iter := 0
     current := table[start]
     for {
     	 for _, m := range movs {
	     if m == 'R' {
	     	pos = current.ids[1]
	     }
	     if m == 'L' {
	     	pos = current.ids[0]
	     }
	     current = table[pos]
	 }
	 iter = iter + len(movs)
	 if pos == end {
	    break
	 }
     }
     fmt.Printf("Steps: %d\n", iter)
     
}