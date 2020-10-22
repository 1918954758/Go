package main

import "fmt"

func main() {
	//创建map
	countryCapitalMap1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	//打印地图
	for country := range countryCapitalMap1 {
		fmt.Println(country, "首都是", countryCapitalMap1[country])
	}
	//删除元素
	delete(countryCapitalMap1, "France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除元素后地图")
	//打印地图
	for country := range countryCapitalMap1 {
		fmt.Println(country, "首都是", countryCapitalMap1[country])
	}
}
