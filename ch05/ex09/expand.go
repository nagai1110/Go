// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "$a a $A $bc bc $BC"
	fmt.Println(str)

	str = expand(str, strings.ToUpper)
	fmt.Println(str)
}

func expand(s string, f func(string) string) string {
	replaceFunc := func(s string) string {
		s = strings.TrimPrefix(s, "$")
		return f(s)
	}

	regex := regexp.MustCompile("\\$([^\\s]*)")
	return regex.ReplaceAllStringFunc(s, replaceFunc)
}
