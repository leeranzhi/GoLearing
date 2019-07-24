package main

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
