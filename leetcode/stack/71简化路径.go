package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/.../a/../b/c/../d/./"
	res := simplifyPath(path)
	fmt.Println("res", res)
}

//输入：path = "/.../a/../b/c/../d/./"
//
//输出："/.../b/d"
func simplifyPath(path string) string {
	var resStack []string
	pathSplit := strings.Split(path, "/")
	for i := range pathSplit {
		switch pathSplit[i] {
		case "..":
			if len(resStack) > 0 {
				resStack = resStack[:len(resStack)-1]
			}
		case ".":
			continue
		case "":
			continue
		default:
			resStack = append(resStack, pathSplit[i])
		}
	}
	var resStr string
	for j := range resStack {
		resStr += "/" + resStack[j]
	}
	if resStr == "" {
		resStr = "/"
	}
	return resStr
}
