package main

import (
	"fmt"
	"math"
	"time"
)

type user struct {
	name string
	pwd  string
}

// 变量类型后置
func myadd(a int, b int) int {
	return a + b
}

// 可以根据上下文推断传参变量类型
func myadd2(a, b int) int {
	return a + b
}

// 同样返回值也可以根据上下文推断类型
func minus(n int) {
	n -= 1
}

// 指针修改传参
func minus2ptr(n *int) {
	*n -= 1
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

// 函数可以返回多个值，第一个值作为真正的返回值，其他值为返回的错误信息

// 函数可以作为参数传递给其他函数

// 函数可以作为返回值返回

// 结构体操作函数同样有指针和非指针两种用法
func checkPWD(u user, pwd string) bool {
	return u.pwd == pwd
}

func checkPWD2(u *user, pwd string) bool {
	return u.pwd == pwd
}

// 结构体方法 - 将结构体变量从传参括号中提到func关键字后面, 函数名前面
func (u user) checkPassword(pwd string) bool {
	return u.pwd == pwd
}

// 注意指针的使用
func (u *user) resetPassword(pwd string) {
	u.pwd = pwd
}

func main() {
	fmt.Println("hello world")

	// var type

	// 两种变量定义方式，有何不同？
	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"

	fmt.Println(a, b, c, d, e, f)

	fmt.Println(g)

	// 静态变量根据上下文自动确定类型

	const s string = "constant"

	const h = 500000000

	const i = 3e20 / h

	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	fmt.Println("===============")

	// 基础判断

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("===============")

	// 循环
	// 在Golang中没有while/do while循环 只有唯一的for循环

	x := 1

	// 无限循环
	for {
		fmt.Println("loop")
		break
	}

	// 有条件循环， for循环的三个参数同c++一样但任何参数都可以省略

	for j := 7; j < 9; j++ {
		fmt.Println("j =", j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n =", n)
	}

	for x <= 3 {
		fmt.Println("x =", x)
		// 以下三种自增方式都支持
		// x += 1
		// x++
		x = x + 1
	}

	fmt.Println("===============")

	// switch 分支判断

	y := 2
	// 与C++不同，每个case不需要break,每个case分支跑完之后立刻结束整个switch判断
	switch y {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	tt := time.Now()
	// golang的switch更加强大，不仅支持一般数值判断，同样支持字符串，结构体等复杂变量
	// 对于复杂判断 ，使用switch代替复杂的if else嵌套，让代码更易读
	switch { //switch中不加判断条件，直接将判断转移到每个case
	case tt.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	if tt.Hour() < 12 {
		fmt.Println("It's before noon by if")
	} else {
		fmt.Println("It's after noon by if")
	}

	// 相比来说case更易读

	fmt.Println("===============")

	// 数组-长度固定 工业级不常用，更多用切片

	var array [5]int

	array[4] = 100

	array2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array, array2)

	var twoD [2][3]int

	for di := 0; di < 2; di++ {
		for dj := 0; dj < 3; dj++ {
			twoD[di][dj] = di + dj
		}
	}
	fmt.Println("2d array:", twoD)

	fmt.Println("===============")

	// 切片slices-长度可变-类似于python列表 工业级使用最多
	// 切片的长度可以在运行时改变，但容量不可变

	myslice := make([]string, 3)
	myslice[0] = "a"
	myslice[1] = "b"
	myslice[2] = "c"
	fmt.Println("get:", myslice[2])
	fmt.Println("len:", len(myslice))

	myslice = append(myslice, "d")
	myslice = append(myslice, "e", "f")
	// 需要注意在append调用后必须赋值给原变量
	// 原因：append操作本质是原切边加上一个扩容然后返回一个指针 指向这个“新”切片，而不是修改原切片
	myslicecopy := make([]string, len(myslice)) // 需要在定义时预设容量大小
	copy(myslicecopy, myslice)
	fmt.Println("myslice copy:", myslicecopy)

	// 支持类似python的索引方式
	fmt.Println(myslice[2:5])
	fmt.Println(myslice[:5])
	fmt.Println(myslice[2:])

	// 也可以直接初始化切片内元素，从而一方面确认容量，另一方面可以省去make操作
	good := []string{"g", "o", "o", "d"}

	fmt.Println(good)

	//总结：
	//:= 中直接定义容量大小，初始化没用make的是数组
	//:= 中不定义容量容量大小直接初始化元素值，初始化要用make的是切片

	fmt.Println("===============")

	// map 类似于其他编程语言中的哈希表/字典
	// 键值对是无序的,遍历结果偏随机，不会按照字母顺序或插入顺序
	// 键值对中的key 必须是支持 == 运算符的类型
	// 值可以是任意类型

	//定义map需要确定key 和 value的变量类型
	mymap := make(map[string]int)

	mymap["one"] = 1
	mymap["two"] = 2
	fmt.Println("mymap:", mymap)
	fmt.Println("len:", len(mymap))
	fmt.Println("get:", mymap["one"])
	fmt.Println("get unknown:", mymap["unknow"])

	//可以通过接收ok变量-boolean来判断map中是否有该key存在
	r, ok := mymap["unknow"]
	fmt.Println(r, ok)

	delete(mymap, "one")

	//同样也可以通过直接赋值来初始化map
	map2 := map[string]int{"one": 1, "two": 2}
	var map3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(map2, map3)

	fmt.Println("===============")

	// range - 对数组/切片/map/字符串进行快速遍历

	numarray := []int{2, 3, 4}

	sum := 0

	for i, num := range numarray { // 返回index和value
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println("sum:", sum)

	test_map := map[string]string{"a": "A", "b": "B"}

	for k, v := range test_map { // 返回key和value
		fmt.Println("key:", k, "value:", v)
	}

	for k := range test_map {
		fmt.Println("key:", k)
	}

	//如果某个返回值不需要均可以用 _ 代替

	fmt.Println("===============")

	// 函数-函数demo见最顶部

	// 函数功能测试

	res := myadd(1, 2)
	fmt.Println("1+2=", res)

	res2 := myadd2(1, 2)
	fmt.Println("1+2=", res2)

	v, ok := exists(map[string]string{"a": "A"}, "a")

	fmt.Println(v, ok)

	fmt.Println("===============")

	// 指针 - 操作有限，基本使用方式为对函数传参的变量进行修改

	mn := 5
	minus(mn)
	fmt.Println("mn after minus:", mn) // 形参 对真实变量对象无修改

	minus2ptr(&mn)                         // 取地址
	fmt.Println("mn after minus2ptr:", mn) // 修改实参

	fmt.Println("===============")

	// 结构体 - 带类型的字段的集合

	user1 := user{name: "user1", pwd: "pwd1"}
	// 也可以直接给初始值
	user2 := user{"user2", "pwd2"}

	user3 := user{name: "user3"} // 明确写出结构体字段名可以只赋值一部分
	// 空值：字符串没初始化为空字符串， 数值就是0
	var user4 user

	// 通过.datafeild形式访问及修改结构体字段内容
	user4.name = "user4"
	user4.pwd = "pwd4"

	fmt.Println(user1, user2, user3, user4)

	// 同时结构体也可以作为函数参数及返回值

	// 同样有指针和非指针两种用法
	// 指针用法可以直接实现对结构体的修改
	// 某些情况下避免大结构体拷贝的开销

	fmt.Println(checkPWD(user1, "pwd1"))
	fmt.Println(checkPWD2(&user1, "pwd1"))

	fmt.Println("===============")

	// 结构体方法 - 类似于C++的成员函数, java的类成员函数
	// 即结构体对象可调用的方法

	test_user := user{name: "test_user", pwd: "test_pwd"}

	test_user.resetPassword("new_pwd")

	fmt.Println(test_user.checkPassword("new_pwd")) // true 修改成功

	fmt.Println("===============")

	// 错误处理

}
