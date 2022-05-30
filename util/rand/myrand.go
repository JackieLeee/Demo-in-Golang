package myrand

import (
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

/**
 * @Author  Flagship
 * @Date  2022/4/20 9:44
 * @Description
 */

type RandMode struct {
	Numbers        bool
	LowerLetters   bool
	UpperLetters   bool
	SpecialSymbols bool
}

const (
	numbers        = "0123456789"
	lowerLetters   = "abcdefghijklmnopqrstuvwxyz"
	upperLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialSymbols = "~!@#$%^&*()[{]}-_=+|;:'\",<.>/?`"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandStr(n int, mode RandMode) string {
	var baseStr strings.Builder
	switch {
	case mode.Numbers:
		baseStr.WriteString(numbers)
		fallthrough
	case mode.LowerLetters:
		baseStr.WriteString(lowerLetters)
		fallthrough
	case mode.UpperLetters:
		baseStr.WriteString(upperLetters)
		fallthrough
	case mode.SpecialSymbols:
		baseStr.WriteString(specialSymbols)
	default:
		return ""
	}

	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < baseStr.Len() {
			b[i] = baseStr.String()[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
