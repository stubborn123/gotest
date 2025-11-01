package main

import (
	"fmt"
)

/**
*问题：使用两个goroutine交替打印，一个goroutine打印数字，另一个goroutine打印字母最终效果如下：
*"12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"
 */

func main() {
	//两个chan，一个是数字，一个是字母，往对应的chan放一个值就会停止阻塞继续打印
	number := make(chan bool)
	letter := make(chan bool)
	//done 是负责停止打印
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j >= 'Z' {
					done <- true
				} else {
					fmt.Print(string(j))
					j++
					number <- true
				}
			}
		}
	}()

	number <- true
	for {
		select {
		case <-done:
			return
		}
	}
}
