package Part7Interface
import "fmt"

// define an interface
type Speaker interface{
	Speak() string
}
// empty interface
/*
✅ 三、空接口 interface{} —— 万能类型（实际开发中非常常用）
空接口没有方法：type Any interface{}

所有类型都实现了空接口（因为没有方法要求）

常用于 JSON 解码、日志框架、通用处理等
*/
func PrintValue(v interface{}){
	fmt.Println("Value is:",v)
}

// define a strcut
type Dog struct{}

func (d Dog) Speak() string  {
 	return "wolf"
}

