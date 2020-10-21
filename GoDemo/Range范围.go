package main

import "fmt"
func main() {
	//使用range去求一个slice的和。使用数组和这个很类似
	nums := []int{1, 2, 3}
	summ := 0
	for _, numm := range nums {
		summ += numm
	}
	fmt.Println("summ : ", summ)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用钙元素的序号，所以我们使用空白符"_"省略了。有时候我们确实需要知道他的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index : ", i)
		}
	}

	//range也可以用在map的键值对上
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s - %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符(Unicode的值)本身。
	for i, c := range "this is a string" {
		fmt.Printf("%d - %d\n", i, c)
	}
}
