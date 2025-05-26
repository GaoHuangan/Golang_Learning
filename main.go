package main
import "LearningProject/Part8Assertion"


func main() {
	Part8Assertion.PrintValue(42)
	Part8Assertion.PrintValue("nickel")
	p := Part8Assertion.Person{Name: "Gao"}
	//p :=Part8Assertion.Person{Name: "Gao"}
	Part8Assertion.WhoAreYou(p)
	Part8Assertion.WhoAreYou("NotPerson")

}