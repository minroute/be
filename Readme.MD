# 主题
1. 类型转换
2. 数据结构扩展




# 内容：
   [ ] 基本类型,如 be.String()，be.Int....
   [ ] Null类型,如 be.NullString()
   [ ] 数组,如 be.Array()
[ ] 结构体,如 be.Struct()
[ ] 切片,如 be.Slice()
[ ] Set,如 be.Set







# 实现方式和目的
通过对现有知名框架代码的学习和代码重组，达到进一步学习掌握的目的。

# 本次学习框架
1. goframe v2 版本，极其优秀的一款框架！


# 如何使用

```go
import "github.com/xiiapp/be"

func main()  {
    be.String("Hello World!")
	be.String(1233)
	be.String(true)
	be.String(nil)
	be.String([]int{1,2,3})
	be.String(map[string]string{"a": "b"})
	be.String(struct{}{})
	
}
```


# 学习成果
1. 断言除了诊断变量的类型外， 还可以诊断变量实现了某个接口。参见 be_string.go 文件里的istring部分
2. 断言+断无可断的情况下，再使用反射，反射不行再用json序列化和fmt包。
3. 断言性能优于反射。
4. switch语句中fallthrough的用法:
   fallthrough：Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。