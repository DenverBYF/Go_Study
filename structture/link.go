package structture

import "fmt"

type Link struct {
	Next *Link
	Val  interface{}
}

func (l *Link) Print() {
	cur := l
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func (l *Link) Reserve() *Link {
	if l == nil || l.Next == nil {
		return l
	}
	var prev *Link
	cur := l
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
