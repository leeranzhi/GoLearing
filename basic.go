package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//不是全局变量
//属于包内变量
//函数外部不能使用:来代替var
var a, b, c = 2, 3, "godV"
var cc = "gkd"

//使用() 来集中定义变量，一般定义包内变量使用
var (
	aa = 3
	ss = "wallsFlower"
	bb = true
)

//普通写法
func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//类型推断
func variableInit() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//不用var 使用:
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	s = "kkp"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello World !!")
	variable()
	variableInit()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, c, aa, bb, cc)

	euler()

	triangle()

	consts()

	enums()
}

//欧拉公式
func euler() {
	////Go语言支持复数表示
	//c := 3 + 4i
	////取模（绝对值）
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n",
		//自动以e为底
		cmplx.Exp(1i*math.Pi)+1)
	//cmplx.Pow(math.E, 1i*math.Pi) + 1
}

//强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量
//命名一般不大写
func consts() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt((a*a + b*b)))
	fmt.Println(fileName, c)
}

//枚举类型
func enums() {
	const (
		//cpp    = 0
		//java   = 1
		//python = 2
		//golang = 3

		//自增
		cpp = iota
		_
		_
		python
		golang
		javascript
	)

	fmt.Println(cpp, python, golang, javascript)
}
