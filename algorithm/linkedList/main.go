/**
 * 数组转链表
 */
package main

import "fmt"

// 链表节点类型
type node struct {
	data int
	next *node
}

func main() {
	var head *node = build([]int{6, 2, 10, 1, 3})
	fmt.Printf("%v\n", head)
}

func build(arr []int) (h *node) {
	var n *node
	for k, v := range arr {
		if n == nil {
			h = new(node)
			n = h
		} else {
			n = n.next
		}
		n.data = v
		if k < len(arr)-1 {
			n.next = new(node)
		}
	}
	return
}
