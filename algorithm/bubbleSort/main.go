/**
 * 冒泡排序
 */
package main

import "fmt"

// 沉底排序函数
func bubbleSort1(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-(i+1); j++ {
			// 升序
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			/*// 降序
			  if a[j] < a[j+1] {
			      a[j], a[j+1] = a[j+1], a[j]
			  }*/
		}
	}
}

// 上浮排序函数
func bubbleSort2(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			// 升序
			/*if a[i] > a[j] {
			    a[i], a[j] = a[j], a[i]
			}*/
			// 降序
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	arr1 := []int{3, 5, 1, 10, 6, 2, 8, 11, 20, 3}
	bubbleSort1(arr1)
	fmt.Printf("%v\n", arr1)
	arr2 := []int{3, 5, 1, 10, 6, 2, 8, 11, 20, 3}
	bubbleSort2(arr2)
	fmt.Printf("%v\n", arr2)
}
