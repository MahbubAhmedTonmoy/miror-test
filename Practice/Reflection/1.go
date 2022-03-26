package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	orderId    int
	customerId int
}
type employee struct {
	name    string
	address string
	country string
	id      int
	salary  int
}

// func createQuery(q interface{}) {
// 	t := reflect.TypeOf(q)
// 	value := reflect.ValueOf(q)
// 	kind := t.Kind()

// 	fmt.Println("type", t)
// 	fmt.Println("value", value)
// 	fmt.Println("kind", kind)

// 	if reflect.ValueOf(q).Kind() == reflect.Struct {
// 		v := reflect.ValueOf(q)
// 		fmt.Println("Number of field", v.NumField())
// 		for i := 0; i < v.NumField(); i++ {
// 			fmt.Println("Field:%d type:%T value:%v\n", i, v.Type(), v.Field(i))
// 		}
// 	}
// }

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()

		query := fmt.Sprintf("insert into %s values(", t)

		v := reflect.ValueOf(q)

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	o := Order{
		orderId:    20,
		customerId: 1,
	}
	createQuery(o)
	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)

	i := 90
	createQuery(i)
}
