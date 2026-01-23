package main

import "fmt"

type List struct {
	items []int
}

func (list *List) Len() int {
	if list != nil {
		return len(list.items)
	}
	return 0
}

type Change interface {
	Set(int)
	Get() int
	Increment()
}

type Str struct {
	val int
}

func (str *Str) Set(val int) {
	str.val = val
}

func (str *Str) Get() int {
	return str.val
}

func (str *Str) Increment() {
	str.val += 1
}

func PrintAll(items []interface{}) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func IntToInterface(items []int) []interface{} {
	result := make([]interface{}, len(items))
	for i, v := range items {
		result[i] = v
	}
	return result
}

type Reader interface {
	Read() string
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type File struct {
	Name    string
	Content string
	isOpen  bool
}

func (f *File) Read() string {
	if !f.isOpen {
		return "Закрыто"
	}
	return f.Content
}

func (f *File) Close() error {
	if f.isOpen {
		f.isOpen = false
		fmt.Println("Закрыт")
		return nil
	}
	return fmt.Errorf("Уже закрыт")
}
