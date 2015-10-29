//12. Equivalent Binary Trees 练习
//待简化代码。。。。
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		} else {
			walk(t.Left)
			ch <- t.Value
			walk(t.Right)
		}
	}
	walk(t)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		fmt.Println(v1, v2)
		if !(v1 == v2 && ok1 == ok2) {
			return false
		} else if ok1 == false && ok2 == false {
			return true
		}
	}
}

func main() {

	t1 := tree.New(1)
	t2 := tree.New(2)
	t3 := tree.New(1)
	//New（）生成的树为二叉搜索树，这个方法只能解决二叉搜索树的等价分析问题。
	fmt.Println(Same(t1, t3))
	fmt.Println(Same(t1, t2))
}
