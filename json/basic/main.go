package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	formatData, err := json.MarshalIndent(basket, "", "    ")
	fmt.Println(string(formatData))

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

}
