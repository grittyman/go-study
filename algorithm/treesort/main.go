/**
 * 二叉树实现插入排序
 */
package main

import "fmt"

/**
 * 树节点结构
 */
type treeNode struct {
	value int
	left, right *treeNode
}

func main() {
	values := []int{ 5, 6, 1, 7, 0, 8, 2, 9, 4, 3 }
	values = Sort(values)
	fmt.Println(values)
}

// Sort
/**
 * 主排序函数
 */
func Sort(values []int) []int {
	var root *treeNode
	// 数组转化成二叉树
	for _, v := range values {
		root = add(root, v)
	}
	// 树展开为数组
	appendValues(values[:0], root)
	return values
}

/**
 * 向树中添加节点
 */
func add(t *treeNode, value int) *treeNode {
	// 空节点则新建节点
	if t == nil {
		t = new(treeNode)
		t.value = value
		return t
	}

	// 判断应该往哪个分支添加节点
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

/**
 * 将树有序平铺展开成数组
 */
func appendValues(values []int, t *treeNode) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
