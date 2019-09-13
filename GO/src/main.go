package main

import "fmt"

func main() {
	m := map[string]int{}
	letter := 'a'
	for i := 0; i < 26; i++ {

		m[string(letter)] = i

		letter++
	}
	fmt.Println(m["f"])

}
