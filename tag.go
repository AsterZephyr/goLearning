package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name_field" binding:"required"`
	Age  int    `json:"age_in_years"`
}

func main() {
	u := User{"Alice", 30}
	t := reflect.TypeOf(u) // 1. 获取 reflect.Type

	// 检查 t 是否是 Struct 类型
	if t.Kind() != reflect.Struct {
		fmt.Println("Not a struct!")
		return
	}

	// 2. 遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // 获取第 i 个字段的 reflect.StructField

		fmt.Printf("Field Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", field.Type)

		// 3. 获取字段的 Tag (类型是 reflect.StructTag)
		tag := field.Tag

		// 4. 使用 Tag 的 Get 方法按 Key 解析
		jsonTag := tag.Get("json")
		bindingTag := tag.Get("binding")

		fmt.Printf("  JSON Tag: '%s'\n", jsonTag)
		fmt.Printf("  Binding Tag: '%s'\n", bindingTag)
		fmt.Println("---")
	}

	// 也可以按名称获取字段
	fmt.Println("\nGetting field 'Name' by name:")
	nameField, ok := t.FieldByName("Name")
	if ok {
		fmt.Printf("  JSON Tag for Name: '%s'\n", nameField.Tag.Get("json"))
	}
}
