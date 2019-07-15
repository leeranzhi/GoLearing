package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数 返回一个值
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil

		//下划线代表不想使用
		//q, _ := div(a, b)
		//return q
	default:
		//panic(" " + op)
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//返回两个值
//13/3=4...1
//对返回值进行命名
func div(a, b int) (q, r int) {
	//建议这样return
	return a / b, a % b

	//也可以这样return
	//q = a / b
	//r = a % b
	//return
}

//函数的参数也可以是一个函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//Go中参数只有值传递
//传入指针类型的参数
//func swap(a, b *int) {
//	//*b, *a = *a, *b
//}
func swap(a, b int) (int, int) {
	//*b, *a = *a, *b
	return b, a
}

func main() {
	//fmt.Println(eval(3, 4, "1"))
	if result, err := eval(3, 4, "y"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	//fmt.Println(div(13, 3))

	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 5))

	a, b := 3, 4
	//swap(&a, &b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}
