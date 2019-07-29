package tree

import "fmt"

//结构体
//Go中只有封装，没有继承和多态
type Node struct {
	Value       int
	Left, Right *Node
}

//为结构体定义方法
//值传递
func (node Node) Print() {
	fmt.Println(node.Value)
}

//值传递
//func (node Node) SetValue(Value int) {
//	node.Value = Value
//}
//编译器自动识别都地址的值或者指针
//指针作为方法的接收者
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node")
		return
	}
	node.Value = value
}

//func (node *Node) Traverse() {
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}

//自定义工厂函数
func CreateNode(testValue int) *Node {
	//返回的是局部变量的地址
	//此处对象分配在堆上内存，参与垃圾回收
	return &Node{Value: testValue}
}

func main() {
	//结构体的实例
	var root Node

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = CreateNode(2)

	//nodes := []Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)
	//root.print()
	root.Right.Left.SetValue(4)

	root.Traverse()

	//root.Right.Left.print()
	//fmt.Println()

}
