package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().Unix())
	var ar [5][5] int
	for j:=0; j<len(ar);j++ {
		for i:=0; i<len(ar);i++{
			ar[i][j]=rand.Intn(10)
			fmt.Print(ar[i][j], " ")
		}
		fmt.Println()
	}
}