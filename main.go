package main
import (

"LearningProject/Part7Interface"
"fmt"
)

func main() {


	var s Part7Interface.Speaker	// Declare a variable of interface type
	s=Part7Interface.Dog{}			// Assign a struct that implements the interface
	fmt.Println(s.Speak())			// Output: Woof!

	Part7Interface.PrintValue(123)
	Part7Interface.PrintValue("abc")
	Part7Interface.PrintValue(true)


}