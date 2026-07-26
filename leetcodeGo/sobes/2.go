package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://sobes.tech/bank/go/livecode/483c9756-1066-4814-bb22-a000c90ffd39

func main() {
	fmt.Println(rle("AAAABBCCXYZDDDDEEEFFFAAAAABBBBBBBBBBBBBBBBBBBBBBBB"))
}

func rle(str string) string {
	ret := strings.Builder{}
	counter := 0
	var buf byte
	for i := range str {
		if buf == 0 {
			buf = str[i]
			counter++
		} else if buf == str[i] {
			counter++
		} else {
			ret.WriteByte(buf)
			if counter != 1 {
				ret.WriteString(strconv.Itoa(counter))
			}
			buf = str[i]
			counter = 1
		}
	}

	if counter != 0 {
		ret.WriteByte(buf)
		if counter != 1 {
			ret.WriteString(strconv.Itoa(counter))
		}
	}

	return ret.String()
}

//A4B2C2XYZD4E3F3A5
//A4B3C2XYZD4E3F3A6B28
