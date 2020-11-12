package main

import "fmt"

func main() {
	//1、数据类型转换： 在需要的时候，将数据类型进行强制转换；
	// 数据类型转换的语法 : Type(value)
	var  num  int8 // num 在转换成为计算机认识的2进制数的时候，会用8个二进制位来表示。
	num = 100

	var num1 int16 // num1 在转换成为计算机认识的2进制数的时候，会用16个二进制数来表示
	num1 = 20

	// 90公斤级....100公斤级，120公斤级...
	//sum := num + num1 // Invalid operation: num + num1 (mismatched types int8 and int16)
	//sum := int16(num) + num1  // 解决方法1:
	sum := num + int8(num1) // 解决方案2
	fmt.Println(sum)

	// 注意：数据类型的转换只能在同类型当中进行转换、重点在数值型
	//age := 18 //数值型里面的整形
	//result := int8(age)
	//fmt.Println(result)

	grade := 86.54 // float
	result := int8(grade)
	fmt.Printf("数据类型是:%T\n",grade)
	fmt.Printf("result的数据类型是:%T\n",result)

	// 别名 ：int8、int16、int32、int64
	// uint别名是byte ：字节     1byte = 8个bit
	var a uint8 = 3 // 大名： 王大锤
	var b byte = 5  //小名 ：铁锤
	sum1 := a + b
	fmt.Println(sum1)

	//关于float的小数点位数
	var pai = 3.1415926 // 打印pai
	fmt.Printf("pai的数值是%f\n",pai)
	//保留2位小数
	fmt.Printf("保留两位小数的值是:%.2f\n",pai)
	//保留3位小数
	fmt.Printf("保留三位小数的数值是:%.3f\n",pai)
	//保留1位小数
	fmt.Printf("保留1位小数的数值是:%.1f\n",pai)


	// 字符串 string
	name := "yuhongwei" // 定义了一个变量 yuhongwei  9
	// 看一下这个字符串的长度, 也就是说字符串当中一共有多少个字符？
	// 计算字符串的长度: len()； length: 长度
	fmt.Println("字符串name的长度是:", len(name))

	address := "山东省泰安市泰山风景区"// 一共有11个汉字
	//汉字所占的空间长度跟字母占的空间长度不一样
	length := len(address) // 定义一个新变量length，用于接收字符串的长度
	fmt.Println("地址address字符串的长度是:",length)

	//截取字符串
	name1 := "tiechuimeimei" // 铁锤妹妹
	// 如何得到tiechui这个部分字符串
	tiechui := name1[0:7]
	fmt.Println(tiechui)

	// 截取tiechuimeimei字符串中的meimei
	xiaojiejie := name1[7: len(name1) ] //从m开始切，切到最后; m的下标是7，
	//字符串的长度是： len(name1)
	//	// 初始下标: 0
	//	// 结束的下标: len(name1) - 1
	fmt.Println(xiaojiejie)

	// 如果是切到字符串的末尾,则 可以省略不写
	xiaomeimei := name1[7:] // 省略的写法
	fmt.Println(xiaomeimei)

	//如果是从头开始截取，也可以省略不写
	// name1 := "tiechuimeimei"
	tie := name1[:7]
	fmt.Println(tie)

	// tiechuimeimei
	// 截取字符串 chuimei  锤妹
	chuimei := name1[3:10]
	fmt.Println(chuimei)
}
