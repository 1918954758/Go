package main

import "fmt"

//定义一个DivideError结构体
type DivideError struct {
	dividee int
	divider int
}

//实现`error`接口
func (de *DivideError) Error() string {
	strFormat := `Cannot proceed, the divider is zero. 
        dividee: %d
        divider: %d`
	return fmt.Sprintf(strFormat, de.dividee, de.divider)
}

//定义`int`类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
func main() {
	//正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	//异常情况
	if result, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("100/0 = ", result, errorMsg)
	}
}
