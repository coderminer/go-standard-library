### Go语言标准库之time

#### 时间的格式化和解析
  
* 格式化 `Format`
  
`Go`语言和其他语言的时间格式化的方式不同，`Go`语言格式化的方式更直观,其他的语言一般是`yyyy-mm-dd`

```
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006年01月02日 15:04:05"))
}

```

输出结果是

```
2018-09-21 10:45:502018/09/21 10:45:50
2018年09月21日 10:45:50
```

在系统中还提供了一些默认的格式

```
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
```

使用方法：

```
fmt.Println(now.Format(time.RFC3339))
```

* 解析 `Parse`

字符串解析为时间

```
func Parse(layout, value string) (Time, error)
```

```
	t1, err := time.Parse("2006-01-02 15:04:05", "2018-09-21 10:54:11")
	t2, err := time.Parse("2006/01/02 15:04:05", "2018/09/21 10:54:59")
	t3, err := time.Parse("2006年01月02日 15:04:05", "2018年09月21日 10:54:59")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
```

输出结果是

```
2018-09-21 10:54:11 +0000 UTC
2018-09-21 10:54:59 +0000 UTC
2018-09-21 10:54:59 +0000 UTC
```

#### 获取时间戳

使用`time.Now`的`time.Unix`和`time.UnixNano`方法获取对应的时间戳

```
	sec := now.Unix()      //秒
	nsec := now.UnixNano() //纳秒
	fmt.Println(sec)
	fmt.Println(nsec)
```

输出结果

```
1537498785
1537498785523262700
```

* 解析时间戳

根据时间戳获取当前时间

```
	t := time.Unix(sec, 0)
	fmt.Println(t)
    fmt.Println(t.Format("2006-01-02 15:04:05"))
```

输出结果

```
2018-09-21 11:05:30 +0800 CST
2018-09-21 11:06:43
```

* 获取当前的日期和时间

```
	year, month, day := now.Date()
	fmt.Println(year, month, day)
	fmt.Println(year, int(month), day)
	fmt.Printf("year:%d month:%d day:%d", year, month, day)

    hour, minute, second := now.Clock()
	fmt.Println(hour, minute, second)
```

输出结果

```
2018 September 21
2018 9 21
year:2018 month:9 day:21
11 11 34
```

#### 关于星期

```
	weekday := now.Weekday()
	fmt.Println(weekday)      // Friday
	fmt.Println(int(weekday)) //5
```

输出结果

```
Friday
5
```

星期日是`0`  

#### 计算已经过去了多少天

```
	days := now.YearDay()
	fmt.Println(days)
```

输出结果

```
264
```

#### 计算两个日期的时间差

```
	date1 := time.Date(2017, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	hours := date2.Sub(date1).Hours()
	fmt.Println(hours)
	between_days := hours / 24
	fmt.Println(between_days)
```

输出结果

```
8760
365
```

#### 获取月份中的天数

`time.Day()`方法是获取本月之前过去的天数，获取月份中天数,初始化时天设置为0，就是获取上月的天数，
月数+1就可以获取本月的天数

```
month_days := time.Date(now.Year(), now.Month() + 1, 0, 0, 0, 0, 0, time.UTC).Day()
fmt.Println(month_days)
```

#### 计算执行的时间

```
    start := time.Now()  //程序执行开始
    //.... 程序代码
	duration := time.Since(start) //执行结束
	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
```

输出结果

```
18.9889ms
18988900
```