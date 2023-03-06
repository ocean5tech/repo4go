package main

import "fmt"

func main() {

	fmt.Println("hello go!")
	fmt.Println(PI)
	fmt.Println(ERR1, ERR2, ERR3, ERR4, ERR5)

	//格式化输出
	uname := "bobby"
	age := []int{18, 19, 20}
	address := []string{"China", "BeiJing", "256"}
	fmt.Printf("用户名： %s, 年龄： %v\r\n", uname, age)                         //只管打印，不返回字符串
	infor := fmt.Sprintf("用户名： %T, 年龄： %#v, 地址: %v", uname, age, address) //返回字符串  %v %s有啥差别？ %#v 带类型输出 %T 打印类型
	fmt.Println(infor)
	infor2 := fmt.Sprintf("用户名： %s, 年龄： %d, 地址: %s", uname, age, address) //返回字符串
	fmt.Println(infor2)

	// 用stringbuilder进行拼接，性能好
}
