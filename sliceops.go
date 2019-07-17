package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("create slice")
	//声明式默认初始化元素为默认值
	var s []int //zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	//数组的创建
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	//数组的创建
	//声明式默认初始化元素为默认值
	s2 := make([]int, 16)
	//指定数组slice大小，底层实际大小
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	//数组拷贝
	fmt.Println("Coping slice")
	copy(s2, s1)
	printSlice(s2)

	//删除数组指定位置元素
	fmt.Println("Delete elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("----")
	printSlice(s2)

	//删除数组头部元素
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	//删除数组尾部元素
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
