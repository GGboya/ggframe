记录下开发这个web框架遇到的问题

1、go run main.go 报错 undefined

解决方案：

golang 编译器 默认只加载一个main包下的一个main.go 文件，如果同一个main包下有多个 .go 的文件，其他文件默认不会加载，如下图：

![766e6583bc42b618fd8f97e7e.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/75f62dd88eba414590998a889709d13f~tplv-k3u1fbpfcp-zoom-in-crop-mark:1512:0:0:0.awebp?)

fsm.go 中定义了变量和函数，main.go 中调用。cmd 运行 go run main.go 就报错undefined.

解决办法 `go run *.go` 或者 `go run .` 文件全部加载



2、fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.

起因是挂了梯子出现问题，把梯子关闭，然后重启goland即可运行。