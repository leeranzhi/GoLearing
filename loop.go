package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//将整数转为二进制
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//循环方式读取文件
//相当于while
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

//没有循环条件
//无限循环
func forver() {
	for {
		fmt.Println("abc")
	}

}
func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1011-->1101
		convertToBin(13132131231),
		convertToBin(0),
	)

	readFile("abc.txt")
	forver()

}
