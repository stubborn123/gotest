
-  使用 panic 和 recover编写一个不包含 return 语句但能返回一个非零值的函数。





- 写一个函数读文件，要求使用 defer 机制关闭文件。




- panic，recover 和 error 的区别的联系

error 普通错误

painc 严重错误

recove 一个捕捉错误的工具


处理错误的两种模式：
（1）error 普通错误 ，可预测可处理的错误
（2）panic + recover 异常级/运行时错误，不可预测，程序运行中的崩溃或极端情况


error
```
type error interface {
    Error() string
}

```

error

panic


recover
从 panic 中恢复



error	     普通错误值	   函数返回值	           调用方检测                   if err != nil	文件未找到、网络超时
⚠️ panic	运行时异常	   panic() 或 系统异常	   程序会崩溃（除非 recover）	越界、严重逻辑错误
🧯 recover	恢复机制	   在 defer 中调用	      捕获 panic，使程序继续执行	防止服务器整体崩溃



特性	         error	                   panic	     recover
类型	         内置接口 (普通值)	        内置函数	   内置函数
控制流	         正常返回	                异常中断	   恢复 return
使用场景	     可预期出错	不可恢复错误	 防止崩溃
是否可捕获	     ✅ if err != nil	     ⚠️ 需 recover	✅ 在 defer 中有效
是否中断程序	 ❌ 否	                  ✅ 是	        ❌ 否
