package structture

import "fmt"

type Tree struct {
	Left  *Tree
	Right *Tree
	Val   interface{}
}

func PreOrder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	PreOrder(t.Left)
	PreOrder(t.Right)
}

func PreOrderTraversal(t *Tree) {
	if t == nil {
		return
	}
	st := Stack{}
	st.Push(t)
	for !st.IsEmpty() {
		v, _ := st.Pop()
		node := v.(*Tree)
		fmt.Println(node.Val)
		if node.Right != nil {
			st.Push(node.Right)
		}
		if node.Left != nil {
			st.Push(node.Left)
		}
	}
}

func InOrder(t *Tree) {
	if t == nil {
		return
	}
	InOrder(t.Left)
	fmt.Println(t.Val)
	InOrder(t.Right)
}

func InOrderTraversal(t *Tree) {
	if t == nil {
		return
	}
	st := Stack{}
	cur := t
	for cur != nil || !st.IsEmpty() {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		v, _ := st.Pop()
		node := v.(*Tree)
		fmt.Println(node.Val)
		if node.Right != nil {
			cur = node.Right
		}
	}
}

func PostOrder(t *Tree) {
	if t == nil {
		return
	}
	PostOrder(t.Left)
	PostOrder(t.Right)
	fmt.Println(t.Val)
}

func PostOrderTraversal(t *Tree) {
	if t == nil {
		return
	}
	st1 := Stack{}
	st2 := Stack{}

	st1.Push(t)
	for !st1.IsEmpty() {
		v, _ := st1.Pop()
		node := v.(*Tree)
		st2.Push(node)

		if node.Left != nil {
			st1.Push(node.Left)
		}
		if node.Right != nil {
			st1.Push(node.Right)
		}
	}

	for !st2.IsEmpty() {
		v, _ := st2.Pop()
		fmt.Println(v.(*Tree).Val)
	}
}
