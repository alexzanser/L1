package main

import "fmt"

func main() {
	rawData := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	
	for i:= -50; i < 60; i += 10 {
		for _, val := range rawData {
			if float64(i) <= val && val < float64(i + 10) {
				groups[i] = append(groups[i], val)
			} 
		}
	}
	fmt.Println(groups)
}
