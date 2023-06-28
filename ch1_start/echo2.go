package main

import "fmt"

func main() {

	// 提示请输入姓名，年龄，薪水，是否通过考试，并将其显示在终端
	var name string
	var age int
	var sal float64
	var isPass bool
	fmt.Println("请输入姓名：")
	if _, err := fmt.Scanln(&name); err != nil {
		fmt.Println("姓名输入有误，请重新输入")
	}
	fmt.Println("请输入年龄：")
	if _, err := fmt.Scanln(&age); err != nil {
		fmt.Println("年龄输入有误，请重新输入")
	}
	fmt.Println("请输入薪水：")
	if _, err := fmt.Scanln(&sal); err != nil {
		fmt.Println("薪水输入有误，请重新输入")
	}
	fmt.Println("请输入是否通过考试：")
	if _, err := fmt.Scanln(&isPass); err != nil {
		fmt.Println("是否通过考试输入有误，请重新输入")
	}
	fmt.Printf("名字是：%v\n年龄是：%v\n薪水是：%v\n是否通过考试：%v\n", name, age, sal, isPass)

}
