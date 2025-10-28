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



interview
