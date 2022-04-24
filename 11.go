package main

import "fmt"

func main() {
	n := map[string]struct{}{
		"a":{},
		"b":{},
		"c":{},
		"d":{},
		"e":{},
		"k":{},
	}
	m := map[string]struct{}{
		"d":{},
		"e":{},
		"f":{},
		"g":{},
		"h":{},
		"k":{},
	}
	inter := make(map[string]struct{})
	if len(n) > len(m) {
		n, m = m, n
	}

	for key_n := range n {
		for key_m := range m{
			if key_n == key_m {
				inter[key_m] = struct{}{}
			}
		}
	}
	fmt.Println("Intersection of given sets:", inter)
}
