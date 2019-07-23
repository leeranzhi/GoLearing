package main

import "fmt"

//结构体
//Go中只有封装，没有继承和多态
type treeNode struct {
	value       int
	left, right *treeNode
}

//值传递
func (node treeNode) print() {
	fmt.Println(node.value)
}

//值传递
//func (node treeNode) setValue(value int) {
//	node.value = value
//}
//编译器自动识别都地址的值或者指针
//指针接收者
func (node *treeNode) setValue(value int) {
	node.value = value
}


//自定义工厂函数
func createNode(testValue int) *treeNode {
	//返回的是局部变量的地址
	//此处对象分配在堆上内存，参与垃圾回收
	return &treeNode{value: testValue}
}

func main() {
	//结构体的实例
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)
	root.print()
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

}
