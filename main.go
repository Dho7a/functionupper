package main

import "fmt"

func toUpper(c byte) byte {
	if (c > 96) && (c < 123) {
		return c - 32
	}
	return c
}

func toUpperString(str string) string {
	ret := ""
	thereisaspace := true
	for i := 0; i < len(str); i++ {

		var c = str[i]
		if !thereisaspace {
			ret = ret + string(c)
		} else {
			ret = ret + string(toUpper(c))
		}

		if c == ' ' {
			thereisaspace = true
		} else {
			thereisaspace = false
		}

	}

	return ret
}

func main() {

	fmt.Printf(toUpperString("Hello i am bacem"))

}
