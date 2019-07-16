package main

import "fmt"

//参数传入指针类型  进行修改
//Go中一般不直接使用数组，而是使用切片
func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	//多维数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//遍历数组
	//range关键字
	for i, v := range arr3 {
		//fmt.Println(arr3[i])
		fmt.Println(i, v)
	}

	fmt.Println("---arr3---")
	printArray(&arr3)
	fmt.Println("---arr1---")
	printArray(&arr1)

	fmt.Println("---arr1,arr3---")
	fmt.Println(arr1, arr3)

}
