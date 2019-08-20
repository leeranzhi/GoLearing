package main

import "fmt"

//函数式编程
//函数可以作为参数  返回\

//闭包
func adder() func(int) int {
	sum := 0  //自由变量
	return func(v int) int {
		sum += v  //局部变量
		return sum
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		//fmt.Println(a(i))
		fmt.Printf("0+1+...+%d=%d\n", i, a(i))
	}

}
