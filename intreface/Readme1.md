### Go语言之Interface(二)

#### 使用指针接收器和值接收器实现接口

```
type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 39}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Beijing", "China"}
	d2 = &a
	d2.Describe()
```

#### 实现多个接口

```
type NormalSalary interface {
	DisplaySalary()
}

type LevaeSalary interface {
	CalculateLeaveLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeaveLeft() int {
	return e.totalLeaves - e.leavesTaken
}

	e := Employee{
		firstName:   "Kevin",
		lastName:    "Lee",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var ns NormalSalary = e
	ns.DisplaySalary()

	var l LevaeSalary = e
	fmt.Println("Leaves left =", l.CalculateLeaveLeft())
```

#### 接口嵌套

在`Go`语言中没有继承的概念，但是通过组合可以实现继承的效果

```
type NormalSalary interface {
	DisplaySalary()
}

type LevaeSalary interface {
	CalculateLeaveLeft() int
}

type SalaryOperator interface {
	NormalSalary
	LevaeSalary
}

	var empOp SalaryOperator = e
	empOp.DisplaySalary()
	fmt.Println("Leaves left = ", empOp.CalculateLeaveLeft())
```

#### 接口零值

零值接口是`nil`,`nil`接口中的`type`和`value`都是`nil`  

```
type Describer interface {
	Describe()
}

	var d4 Describer
	if d4 == nil {
		fmt.Printf("d4 is nil and has type %T value %v\n", d4, d4)
	}
```

输出结果是

```
d4 is nil and has type <nil> value <nil>
```

[更多精彩内容](http://coderminer.com)