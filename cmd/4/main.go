package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	lib "github.com/thebenkogan/Advent-Of-Code-2016"
)

func main() {
	key := lib.GetInput()
	fiveFound := false
	for n := 0; ; n++ {
		hash := md5.Sum([]byte(key + strconv.Itoa(n)))
		hex := hex.EncodeToString(hash[:])
		if hex[:5] == "00000" && !fiveFound {
			fmt.Println(n)
			fiveFound = true
		}
		if hex[:6] == "000000" {
			fmt.Println(n)
			break
		}
	}
}
