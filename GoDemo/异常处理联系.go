package main

import "fmt"

//定义结构体
type DeError struct {
	number1 int
	number2 int
}

//继承接口Error()
func (de *DeError) Error() string {
	errFormat := `被除数不能为0.
		a: %d
		b: %d`
	return fmt.Sprintf(errFormat, de.number1, de.number2)
}

//实现接口
func Dee(n1 int, n2 int) (result int, error string) {
	if n2 == 0 {
		d := DeError{
			number1: n1,
			number2: n2,
		}
		error = d.Error()
		return
	} else {
		return n1 / n2, ""
	}
}

//main
func main() {
	if result, error := Dee(100, 10); error == "" {
		fmt.Println("number1/number2 = ", result)
	}
	if _, error := Dee(100, 0); error != "" {
		fmt.Println("number1/number2 = ", error)
	}
}
