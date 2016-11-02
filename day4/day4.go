package day4

import (
	"crypto/md5"
	"strconv"
	"strings"
	"encoding/hex"
	"fmt"
)

func findMatchingMD5(prefix string) int64 {
	suffix := int64(0)
	for {
		md5hash := calculateMD5(prefix, suffix)
		if md5Matches(md5hash) {
			break;
		}
		suffix += 1
	}
	return suffix
}

func md5Matches(md5Hash [16]byte) bool {
	md5String := hex.EncodeToString(md5Hash[:])
	return strings.HasPrefix(md5String, `000000`)
}

func calculateMD5(prefix string, suffix int64) [16]byte {
	var md5hash [16]byte
	inputString := prefix + strconv.FormatInt(suffix, 10)
	md5hash = md5.Sum([]byte(inputString))
	return md5hash
}

func Day4() {
	suffix := findMatchingMD5("yzbqklnj")
	fmt.Printf("Answer is %d\n", suffix)
}
