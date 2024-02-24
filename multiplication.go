package main

import "fmt"

func table(x int, y int){
	answer := x * y
	fmt.Printf("%d x %d = %d \n", x, y, answer)
}

func main(){
	for i:= 1; i<= 12; i++ {
		fmt.Printf("\n")
		fmt.Printf("multiplication table of %d \n", i)

		for j:= 1; j<= 12; j++ {

		table(i, j)

		}
	}
}