package Part8Assertion

import "fmt
import "fmt"

/*
接口类型的变量底层保存的是：具体值 + 类型信息。
如果你想从接口变量中取出原来的类型值，就需要用 类型断言语法：

go
复制
v, ok := i.(T)
部分	含义
i	是一个接口类型的变量
T	是你想断言的目标类型
v	是从接口中提取出来的值
ok	判断断言是否成功（true/false）
*/
// ✅ 示例 1：接口断言出原始类型
func PrintValue(v interface{}) {
	// try to convert to int
	if num, ok := v.(int); ok {
		fmt.Println("It's an int", num)
	} else {
		fmt.Println("Not an int")
	}
：
}
定义类型
//
//	✅ 示例 2
// ✅ 示例 2：断言自定义类型
//	✅ 示例 2：断言自定义类型
tring
}

func WhoAre
type Person struct {
	Name string
}
func WhoAreYou(v interface{}) {
're:", p.Name)
	} else {
		fm
	if p, ok := v.(Person); ok {
		fmt.Println("You're:", p.Name)
	} else {
		fmt.Println("Not a Person")
	}
}