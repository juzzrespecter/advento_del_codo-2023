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

func validateRow(line, config string) {
     groups := line.split('.')
     sizes := config.split(',')

     if len(sizes) != len(groups) {
     	return 0
     }
     for i, group := range groups {
     	 if len(group) != sizes[i] - '0' {
	    return 0
	 }
     }
     return 1
}

func chooseState(line, config string, idx int) int {
     val := 0
     
     if idx == len(line) {
     	return validateRow(line)
     }
     for _, c := range line {
     	 if line[idx] == '?' {
     	    line[idx] = '.'
	    val = chooseState(line, config, idx + 1)
	    line[idx] = '#'
	    return val + chooseState(line, config, idx + 1)
     	 }
	 idx++
     }
}

func main() {
     input, e := os.Open("./input")

     check_err(e)
     rd := bufio.NewReader(input)

     cont := 0
     for {
     	 line, _, e := rd.ReadLine()

	 if e == io.EOF {
	    break 
	 }
	 row, config := line.split(' ')[:2]
	 test := chooseState(row, config, 0)
	 cont = cont + test
     }
     fmt.Printf("Num. combinaciones totales: %d\n", cont)
}