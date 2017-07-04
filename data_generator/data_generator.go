package main

import (
	"math/rand"
	"fmt"
	"time"
)

const (
	absMin int = 10000
	absMax int = 99999
	usalDiff int = 100
)
func main() {
	
	rand.Seed(time.Now().Unix())
	
	var values []int
	
	for i := 0; i < 100; i++ { 
		
		var val int
		
		if i == 0 {
			val = randInt(absMin,absMax)
		} else {
			val = randInt(values[i-1]-usalDiff,values[i-1]+usalDiff)
		}
		
		
		fmt.Println(val)
		values = append(values,val)
	}
	
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}