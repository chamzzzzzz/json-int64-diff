package main

import "fmt"

//go:generate protoc --go_out=. --go_opt=module=github.com/chamzzzzzz/json-int64-diff/main json.proto

type JSONInt64 struct {
	Int64A int64 `json:"int64_a"`
	Int64B int64 `json:"int64_b"`
}

type JSONInt64StringTag struct {
	Int64A int64 `json:"int64_a,string"`
	Int64B int64 `json:"int64_b,string"`
}

func main() {
	fmt.Println("json int64 diff")
}
