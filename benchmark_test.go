package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"reflect"
	"strconv"
	"testing"

	externalErrors "github.com/pkg/errors"
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
var err2 = errors.New("err 222")
var err3 = errors.New("err 333")

type value struct {
	num int
}

var event = struct {
	ActionID int
	AccID    string
	AppID    string
}{
	100, "sdkjfhskldfhslkd", "skjdfhsldkfhlsdkfjh",
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

func BenchmarkIncrementInt(b *testing.B) {
	a := value{0}
	for i := 0; i < b.N; i++ {
		a.num++
	}
}

func BenchmarkIncrementIntWithCastFromInterface(b *testing.B) {
	a := struct {
		num interface{}
	}{
		num: 0,
	}
	for i := 0; i < b.N; i++ {
		a.num = a.num.(int) + 1
	}
}

func BenchmarkCastIntFromInterface(b *testing.B) {
	a := struct {
		num interface{}
	}{
		num: 0,
	}
	var c int
	for i := 0; i < b.N; i++ {
		c = a.num.(int) + 1
		a.num = c
	}
}

func BenchmarkCustomHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := uint64(14695981039346656037)
		for _, c := range []byte(strconv.Itoa(event.ActionID) + event.AccID + event.AppID) {
			id *= 1099511628211
			id ^= uint64(c)
		}
	}
}

func BenchmarkMD5Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := md5.New()
		m.Write([]byte(strconv.Itoa(event.ActionID) + event.AccID + event.AppID))
		m.Sum(nil)
	}
}

func BenchmarkFnvHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := fnv.New64()
		h.Write([]byte(strconv.Itoa(event.ActionID) + event.AccID + event.AppID))
		h.Sum64()
	}
}

func BenchmarkErrof(b *testing.B) {
	var newErr error
	for i := 0; i < b.N; i++ {
		newErr = fmt.Errorf("new error %w", err2)
	}
	_ = newErr
}

func BenchmarkWrap(b *testing.B) {
	var newErr error
	for i := 0; i < b.N; i++ {
		newErr = externalErrors.Wrap(err2, "new error")
	}
	_ = newErr
}
