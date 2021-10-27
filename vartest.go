package main

import "fmt"

//一次性声明多个全局变量
var (
	n4 = 300
	n5 = 3.1415926
	n6 = "money and beauty"
)

func main() {
	//test
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	//int 默认值0，string 默认值为空串，小数默认为 0
	var i int
	fmt.Println("i=", i)
	//第二种：根据值自行判定变量类型(类型推导)
	var num = 10.11
	fmt.Println("num=", num)
	var str = "i love you"
	fmt.Println(str)
	//第三种：省略 var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	name := "lihan"
	fmt.Println(name)
	//str := "me too",这种方法是错误的，因为str在前面已经声明过了
	//第四种：多变量声明，在编程中，有时我们需要一次性声明多个变量，Golang 也提供这样的语法
	//一次性声明多个普通变量
	n1, n2, n3 := 10, 10.11, "you are the best"
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	//全局变量测试
	fmt.Println("n4=", n4, "n5=", n5, "n6=", n6)

}

/*
Go内置以下基础类型:
布尔类型:bool
整型:int8、byte、int16、int、uint、uintptr等
浮点点类型:float32、float64
复数类型:complex64、complex128
字符 :string
字符类型:rune
错误类型:error
*/
