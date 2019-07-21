package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//每个中文占3个字节
	//其他的占1个字节
	s := "Yes这里是慕课网！！" //UTF-8
	fmt.Println(s)
	fmt.Println(len(s))

	//使用range遍历pos，rune对
	//使用len获得字节长度
	//使用[]byte获得字节
	for _, b := range []byte(s) {
		//ASCII码
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()


	//使用utf8.RuneCountInString获得字符的数量
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

}
