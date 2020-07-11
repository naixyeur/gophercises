package main

import (
	"fmt"
)

func main() {

	var length, delta int
	var input string

	fmt.Scanf("%d", &length)
	fmt.Scanf("%s", &input)
	fmt.Scanf("%d", &delta)

	fmt.Printf("%v \n", input)
	fmt.Printf("%v \n", cipher(input, delta))

}

func rotate(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rebaseRotate(r, 'A', delta)
	} else if r >= 'a' && r <= 'z' {
		return rebaseRotate(r, 'a', delta)
	}
	return r
}

func rebaseRotate(r rune, base int, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func cipher(str string, delta int) string {
	ret := make([]byte, 0, len(str))
	for _, r := range str {
		ret = append(ret, byte(rotate(r, delta)))

	}
	return string(ret)
}

// func cipherV2(str string, rotation int) string {
//   rot := uint8(rotation)
//   ret := make([]byte, 0, len(str))
//   for _, r := range []byte(str) {

//     if r >= 'A' && r <= 'Z' {
//       r += rot
//       if r > 'Z' {
//         r -= 26
//       }
//     } else if r >= 'a' && r <= 'z' {
//       r += rot
//       if r > 'z' {
//         r -= 26
//       }
//     }

//     ret = append(ret, byte(r))
//   }
//   return string(ret)
// }
