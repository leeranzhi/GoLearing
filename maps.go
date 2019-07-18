package main

import "fmt"

//map使用哈希表，必须可以比较相等
//除了slice，map，function的内建类型都可以作为key
//Struct(自定义)类型不包含以上字段，也可作为key
func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2==empty map
	if m2 != nil {
		fmt.Println("m2!=nil")
	}

	var m3 map[string]int //m3==nil
	if m3 != nil {
		fmt.Println("m3!=nil")
	} else {
		fmt.Println("m3==nil")
	}

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	//遍历map,(HashMap 是无序的)
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	//不存在的话 拿到的是空字符串
	//causeName := m["cause"]
	//fmt.Println(causeName)
	//第二个参数来判断是否存在
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	//删除元素
	delete(m, "name")
}
