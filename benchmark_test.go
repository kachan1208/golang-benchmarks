package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

type testInterface struct {
	Message interface{}
	Error   error
}

type testStruct struct {
	Message msg
	Error   error
}

type msg struct {
	body string
	var1 int
	var2 int
	var3 int
}

var body = `123456789012345678901234567890sdlfksdlf;h sdlkfjhsdflkjhsdflk jsdflkajh sdflkj 
            ahsdlfkjahsd flkjasdh flaksjdhf lkasjdhflsakdjhflaksjdhfalskdjhfalsjkdhflaskdj
            hfalsdkjfhlaskdjfhlsakdjfhalskdjfhsalkdjhf`
var message = msg{body, 1, 2, 3}
var err error

type value struct {
	num int
}

func BenchmarkIterateMapWithoutPointer(b *testing.B) {
	data := make(map[int]value)

	for i := 0; i < b.N; i++ {
		data[i] = value{
			num: i,
		}
	}

	var c int
	for k, v := range data {
		c = k
		c = v.num
		_ = c
	}
}

func BenchmarkIterateMapWithPointer(b *testing.B) {
	data := make(map[int]*value)
	for i := 0; i < b.N; i++ {
		data[i] = &value{
			num: i,
		}
	}

	var c int
	for k, v := range data {
		c = k
		c = v.num
		_ = c
	}
}

func BenchmarkStructEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(testStruct{message, err})
	}
}

func BenchmarkInterfaceEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(testInterface{message, err})
	}
}

func BenchmarkPassStructIntoEncodeFunc(b *testing.B) {
	encode := func(response msg, err error) {
		json.Marshal(testStruct{response, err})
	}

	for i := 0; i < b.N; i++ {
		encode(message, err)
	}
}

func BenchmarkPassInterfaceIntoEncodeFunc(b *testing.B) {
	encode := func(response interface{}, err error) {
		json.Marshal(testInterface{response, err})
	}

	for i := 0; i < b.N; i++ {
		encode(message, err)
	}
}

func BenchmarkArrayEncode(b *testing.B) {
	var arr [1]string
	arr[0] = body
	for i := 0; i < b.N; i++ {
		json.Marshal(arr)
	}
}

func BenchmarkSliceEncode(b *testing.B) {
	var arr []string
	arr = append(arr, body)
	for i := 0; i < b.N; i++ {
		json.Marshal(arr)
	}
}

func BenchmarkPassArrayIntoEncodeFunc(b *testing.B) {
	encode := func(data [1]string) {
		json.Marshal(data)
	}

	var arr [1]string
	arr[0] = body
	for i := 0; i < b.N; i++ {
		encode(arr)
	}
}

func BenchmarkPassSliceIntoEncodeFunc(b *testing.B) {
	encode := func(data []string) {
		json.Marshal(data)
	}

	var arr []string
	arr = append(arr, body)
	for i := 0; i < b.N; i++ {
		encode(arr)
	}
}

func BenchmarkPassArrayByPointerIntoEncodeFunc(b *testing.B) {
	encode := func(data *[1]string) {
		json.Marshal(data)
	}

	var arr [1]string
	arr[0] = body
	for i := 0; i < b.N; i++ {
		encode(&arr)
	}
}

//Yep, I know, Slice is already pointer
func BenchmarkPassSliceByPointerIntoEncodeFunc(b *testing.B) {
	encode := func(data *[]string) {
		json.Marshal(data)
	}

	var arr []string
	arr = append(arr, body)
	for i := 0; i < b.N; i++ {
		encode(&arr)
	}
}

func BenchmarkMapGetKeysReflect(b *testing.B) {
	testMap := make(map[int]int)

	for i := 0; i < b.N; i++ {
		testMap[i] = i
	}

	keys := reflect.ValueOf(testMap).MapKeys()
	_ = keys
}

func BenchmarkMapGetKeysForeach(b *testing.B) {
	testMap := make(map[int]int)

	for i := 0; i < b.N; i++ {
		testMap[i] = i
	}

	var keys []int
	for key := range testMap {
		keys = append(keys, key)
	}

	_ = keys
}
