package Part7Interface
import "fmt"
//关键点	说明
//Mover 是接口	只要实现了 Move() 方法的类型都算
//Drive(m Mover)	是一个接收接口参数的函数，具备多态性
//Car 和 Person	各自实现了 Move()，都可以被 Drive() 调用
//多态体现	即使变量类型不同，只要实现了接口就可以统一处理
//多态体现	即使变量类型不同，只要实现了接口就可以统一处理

// define the Mover interface
type Mover interface {
	Move() string
}

// define Car struct
type Car struct {
	Brand string
}

// Car implements Mover interface
func (c Car) Move() string {
	return fmt.Sprintf("the car %s is driving", c.Brand)
}

// define Persion struc
// define Persion struct
type Person struct {
	Name string
}

// Person implement Mover interface
func (p Person) Move() string {
	return fmt.Sprintf("%s is walking.", p.Name)
}

// Driver accepts any Mover and calls Move()
func Driver(m Mover) {
	fmt.Println(m.Move())
}