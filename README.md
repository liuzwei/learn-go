## Go 语言学习笔记

---
## 1. 常用命令

- `go build` 编译
- `go run xxx.go` 编译并运行go文件
- `go env` 查看环境变量

## 2.知识点

### 2.1下划线在import中的使用

下划线在import时，导入的包，只有init函数会被执行，其他函数无法被调用！
如下：test1/hello.go中的方法不会被调用, test1会提示undefined
```go
import (
	"fmt"
	_ "learn-go/test1" 
)

// 主入口函数
func main() {
    fmt.Println("2.主函数")
    // # command-line-arguments
    //.\main.go:11:2: undefined: test1
    test1.SayHello()
}
```

### 2.2 下划线在代码中

下划线在代码中，像是占位符，不可使用。

### 2.3