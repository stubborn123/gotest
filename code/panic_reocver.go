package main

import "fmt"

func mayPanic() {
	panic("something went wrong")
}

func doWork() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("⚠️ 捕获 panic:", r)
		}
	}()
	fmt.Println("在panic发生前。")
	mayPanic()
	fmt.Println("这句不会执行，因为上面 panic 了")
}

func main() {
	doWork()
	fmt.Println("✅ 程序继续运行，没有崩溃")
}
