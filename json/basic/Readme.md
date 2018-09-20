### Go语言标准库之JSON编解码

#### 基本的类型

`Go`语言中的数据类型和`JSON`的数据类型的关系 

* bool -> JSON boolean
* float64 -> JSON numbers
* string -> JSON strings
* nil -> JSON null

#### Struct to JSON

`Go`包`encoding/json`中的`json.Marshal`方法，可以将`struct`编码为`JSON`数据

```
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64  `json:"ref"`
	private string // 这个字段不会被编码
	Created time.Time
}

func main() {
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	jsonData, err := json.Marshal(basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
}
```

输出结果

```
{"Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999,"Created":"2018-09-20T11:40:05.9885387+08:00"}
```

* 只有`public`字段才能被导出
* `json:tag`可以定义导出字段的名称
* nil导出为null

#### 格式化输出

可以使用 `json.MarshalIndent`方法来格式化输出的`JSON`数据  

```
	formatData, err := json.MarshalIndent(basket, "", "    ")
	fmt.Println(string(formatData))
```

格式化输出结果是 

```
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-09-20T11:45:49.0969176+08:00"
}
```

#### JSON to Struct

`Go`包中`json.Unmarshal`方法，解析 `JSON`数据

```
	jsonStr := []byte(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999,
		"Created": "2018-09-20T11:45:49.0969176+08:00"
	}
	`)

	var basketStruct FruitBasket
	err = json.Unmarshal(jsonStr, &basketStruct)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basketStruct.Name, basketStruct.Fruit)
	fmt.Println(basketStruct.Created)
```

输出结果是

```
Standard [Apple Banana Orange]
2018-09-20 11:45:49.0969176 +0800 CST
```

#### 编码格式化任意类型的对象和数组

* `map[string]interface{}`可以编码任意类型的 `JSON` 对象
* `[]interface{}`可以编码任意类型的`JSON`数组

```
	arbitData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	var v interface{}
	err = json.Unmarshal(arbitData, &v)
	if err != nil {
		log.Println(err)
	}
	data := v.(map[string]interface{})
	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, v, "(array)")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
```

输出结果是

```
Name Eve (string)
Age 6 (float64)
Parents [Alice Bob] (array)
     0 Alice
     1 Bob
```

#### 编解码JSON文件

`Go`中`encoding/json`中的`json.Decoder`和`json.Encoder`方法可以编解码`JSON`格式文件

例子：

```
	jsonstr := `
	{"Name": "Alice", "Age": 25}
	{"Name": "Bob", "Age": 22}`
	reader := strings.NewReader(jsonstr)
	writer := os.Stdout

	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(writer)

	for {
		var m map[string]interface{}
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for k := range m {
			if k == "Age" {
				delete(m, k)
			}
		}

		if err = enc.Encode(&m); err != nil {
			log.Println(err)
		}
	}
```

输出结果是

```
{"Name":"Alice"}
{"Name":"Bob"}
```

[更多精彩内容](http://www.coderminer.com)