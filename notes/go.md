### GMP
GMP
G goroutine 协程（用户态线程， 没有系统态就不会切换避免开销）
M machine 操作系统线程，负责执行goroutine
P processor 调度器，维护这个goroutine队列，将goroutine分配给machine来执行

以time.Sleep来看GMP：
（1）
执行time.Sleep 对应的当前goroutine运行状态变为阻塞状态，这个goroutine会被这个processor调度器从“运行队列”中移除这个go程。
（2）
这个睡眠时间，go会被定时器管理器记录唤醒时间，这个定时器管理器由一个专门的goroutine来维护。
（3）
p的任务调度，他会在自己“本地运行队列”和“全局运行队列”，寻找其他可运行的goroutine。
    （1） 有可运行的goroutine，p给当前的m去执行
    （2） 没有就去“窃取”别的p去找可运行的G给当前m执行，没有就进入休眠状态
（4）
睡眠时间到了，这个阻塞状态的goroutine，会被对应的定时器管理goroutine放到一个processor里，当p执行到这个go程就会恢复执行    


你可以理解为GMP为一组完整的执行单元，go运行的会可能由多组GMP。
一个M绑定一个P（一个系统线程绑定一个调度器）
但是一个P可以管理多个G（调度器处理管理多个go程）


### pair
go表示两个逻辑上相关联的值，被打包或一起使用

使用的场景报错就是”主要结果“和”状态信息“一并携带，这个在go的工程代码经常有这种结果和异常信息这种成对出现的情况。

pair结合泛型（go的1.18引入泛型），结合pair创建一个struct
```
type pair[T, U any] struct {
	a T
	b U
}
```


### Panic 

#### （1）panic
再Go中，panic表示程序在运行时遇到‘无法恢复的错误’
一旦发生panic：
    当前函数会立刻停止执行
    已经被延迟，defer的函数会依次执行
    然后程序会向上层调用栈传播panic
    最终如果没有被特殊恢复（recover），程序会崩溃退出
是一种运行时中断机制

```
package main

import "fmt"

func main() {
    fmt.Println("before panic")

    panic("something went wrong") // 触发 panic

    fmt.Println("after panic") // 永远不会执行
}

```

#### （2）panic和recover
panic和普通的error不一样，程序会中断，程序会崩溃。
revocer会拯救这个程序

recover怎么拯救一个panic的程序？
recover()只能defer中才有用，因为defer是最后执行的，即使panic也会执行的。

```
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

    mayPanic()
    fmt.Println("这句不会执行，因为上面 panic 了")
}

func main() {
    doWork()
    fmt.Println("✅ 程序继续运行，没有崩溃")
}

```

输出结果
```
⚠️ 捕获 panic: something went wrong
✅ 程序继续运行，没有崩溃

```
开始向上回溯函数调用栈（unwind stack）；
执行每个栈帧中注册的 defer；
如果在某个 defer 中调用了 recover()，那么：
    当前 panic 会被“捕获”；
    程序恢复正常执行；
    recover() 返回传入 panic() 的值。

### chan
#### 1 读写已关闭的chan
（1）读已经关闭的chan，永远会读到的值（chan有就返回，这个没有返回零值 + ok=false）
（2）写已经关闭的chan，会报panic


### go的属性
首先大写还是小写，表示对应的函数或者字段是公有私有的。

首字母小写表示的私有属性是“包内私有”，一个私有函数，属性是包内任意函数属性可以访问，而不是Java的“类内私有”，这有很大的不同。

```
package main

import (
    "encoding/json"
    "fmt"
)

type People struct {
    name string `json:"name"`
}

func main() {
    js := `{
        "name":"11"
    }`
    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println("people: ", p)
}
````

打印结果

```
people:  { }
```
因为name属性是私有的，json包和fmt包都访问不了这个name属性。
p的type是共有的，fmt可以访问，这个p对应可以访问，但是里面属性看不到展示一个默认的空值。





### 接口无线递归的形式

```
package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}

```
会无限循环报错

这个String方法，是一个接口interface(可以把这个String的结构体方法注释掉，就会报错，显示你还有接口方法没实现)

但是无限循环，是因为调用fmt的方法，会继续调用String方法，但是这个方法指向对应的p的实现结构体函数方法，一致不断掉用自己。

(p *People) String() -》fmt String() -》(p *People) String() 

不断循环


go 
interview
