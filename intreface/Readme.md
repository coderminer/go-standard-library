### Go语言之Interface(一)

#### 什么是interface

在面向对象语言中接口是：接口定义了一个对象的行为，但在`Go`中接口就是方法签名的集合，当一个类型提供了这个接口中的所有的方法，就可以说这个类型实现了这个接口

#### 接口的声明和实现

```
package main

import (
	"fmt"
)

// 接口的声明
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// 接口的实现
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Andreson")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}

```

* 声明一个接口
```
// 接口的声明
type VowelsFinder interface {
	FindVowels() []rune
}
```

* 实现接口

```
type MyString string

// 接口的实现
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
```

#### 练习使用接口

```
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Totla Expense per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 5000, 30}
	cemp1 := Contract{3, 3000}

	employes := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employes)
}
```

#### 接口的内部表示

接口内部可以看着是一个元组 `(type,value)`,`type`表示接口的具体类型，`value`表示这个具体类型的值

```
type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(89.4)
	t = f
	describe(t)
	t.Tester()
}
```

输出结果

```
Interface type main.MyFloat value 89.4
89.4
```

#### 空接口

如果一个接口没有包含方法，那么这个接口就是一个空接口 `interface{}`,由于空接口中没有方法，所以所有的类型都可以实现空接口

```
func descrip(i interface{}) {
	fmt.Printf("Type = %T,value = %v\n", i, i)
}

func main() {
	s := "Hello,World"
	descrip(s)
	i := 55
	descrip(i)

	strt := struct {
		name string
	}{
		name: "CoderMiner",
	}
	descrip(strt)
}

```

输出结果是

```
Type = string,value = Hello,World
Type = int,value = 55
Type = struct { name string },value = {CoderMiner}
```

#### 类型断言(类型转换)

类型转换的语法 `i.(T)`,把`i`转换为类型`T`

如转换为`int`类型

```
func asstert(i interface{}) {
	s := i.(int)
	fmt.Println(s)
}

var s1 interface{} = 56
asstert(s1)
```

但如果在类型转换时，如果是不同类型转换会发生什么呢？

如`string`转换为`int`时

```
func asstert(i interface{}) {
	s := i.(int)
	fmt.Println(s)
}
var s1 interface{} = "hello"
asstert(s1)
```

输出结果是

```
panic: interface conversion: interface {} is string, not int
```
会触发`panic`,类型转换失败，如果要避免出现`panic`,可以使用下面的方法 

```
v, ok := i.(T)
```
如果能够转换成功，`ok`的值是`true`，转换失败就是`false`

```
func asstert(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}
var s1 interface{} = "hello"
asstert(s1)
var s2 interface{} = 56
asstert(s2)
```

输出结果

```
0 false
56 true
```

#### 类型判断

语法是

```
i.(type) 
```

```
func FindType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("this is the string type and the value is %s\n", i.(string))
	case int:
		fmt.Printf("this is the int type and the value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

FindType("CoderMiner")
FindType(77)
FindType(69.3)
```

[更多精彩内容](http://coderminer.com)