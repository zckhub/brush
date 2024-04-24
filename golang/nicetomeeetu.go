package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToNumber(str string, offset int) int {
	number := ""
	for _, char := range str {
		number += fmt.Sprintf("%d", char)
	}
	var result int
	fmt.Sscanf(number, "%d", &result)
	return result - offset
}

func main() {
	//defer waitForSignal()

	var input string
	fmt.Print("请输入字符串: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	strList := strings.Split(input, " ")
	if len(strList) != 4 {
		fmt.Println("长度不是4，重新来")
	}
	result1 := strconv.Itoa(stringToNumber(strList[0], 11010599006))
	result2 := strconv.Itoa(stringToNumber(strList[1], 116111))
	result3 := strconv.Itoa(stringToNumber(strList[2], 109101101034))
	result4 := strconv.Itoa(stringToNumber(strList[3], 0)%10 - 1)
	fmt.Println("输出结果:", result1+result2+result3+result4)
}
