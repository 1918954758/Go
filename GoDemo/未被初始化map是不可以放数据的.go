package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string //如此声明map是不可以放数据的
	//countryCapitalMap := make(map[string]string)	//如此，使用make来声明的map是可以放数据的
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	//使用键输出地图值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	//查看元素在集合中是否存在
	capital, ok := countryCapitalMap["American"] //如果确定是真实的，则存在，否则不存在
	if ok {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("American的首都不存在", capital)
	}
}
