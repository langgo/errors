# errors

## TODO

- 可选的记录栈信息
- 可选的记录自定义信息 WithTipMessage
- record not found。 在原有的错误的基础上追加额外信息
- 良好的日志打印。打印错误，TipMessage，栈等
- 查看一下 Go2 的错误值部分
- 参数错误（外部错误） 内部错误
- 像 logrus 一样 log = log.WithFields() errm := errors.NewM("参数错误")
- 消息覆盖等，前后调用了两次WithTip
- 并发的时候，会有问题吗？应该不会因为是由内向外返回值
- // TODO 下面的错误都是参数错误。这看起来挺好的。需要改进错误值类型，让下面错误更优雅一下。错误前缀？错误组？// err.Log() err.Log(log)
- NewF("fmt", arg1, arg2) WrapF()
- 跨系统的时候，在调试模式，需要记录错误链

https://go.googlesource.com/proposal/+/master/design/go2draft.md
https://github.com/golang/exp/tree/master/errors

错误码

- uber zap Named
- uber NewDevelopment
- uber NewProduction

