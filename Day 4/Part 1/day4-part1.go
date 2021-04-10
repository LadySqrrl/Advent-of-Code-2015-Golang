//finds the lowest number that produces an MD5 hash that starts with 5 zeros when appended to the key
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// MD5 hashes using md5 algorithm
func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func main() {

	key := "iwrupvqb"

	tempHash := ""
	found := false

	for i := 0; found == false; i++ {
		keyNum := key + strconv.Itoa(i)
		tempHash = MD5(keyNum)

		if strings.HasPrefix(tempHash, "00000") {
			fmt.Println(keyNum)
			found = true
		}
	}

}
