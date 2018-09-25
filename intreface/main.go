package main

import (
	"fmt"
)

type VowelsFinder interface {
	FindVowels() []rune
}

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
	fmt.Printf("Totla Expense per Month $%d\n", expense)
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

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

func descrip(i interface{}) {
	fmt.Printf("Type = %T,value = %v\n", i, i)
}

func asstert(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

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

func main() {
	name := MyString("Sam Andreson")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 5000, 30}
	cemp1 := Contract{3, 3000}

	employes := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employes)

	var t Test
	f := MyFloat(89.4)
	t = f
	describe(t)
	t.Tester()

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

	var s1 interface{} = "hello"
	asstert(s1)
	var s2 interface{} = 56
	asstert(s2)

	FindType("CoderMiner")
	FindType(77)
	FindType(69.3)
}
