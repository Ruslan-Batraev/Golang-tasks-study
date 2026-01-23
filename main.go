package main

import (
	"context"
	"fmt"
)

func main() {
	//fmt.Println("Task1")
	//person := Person{"Steve", 112}
	//fmt.Println(person.Greet())
	//
	//person.Birthday1()
	//fmt.Println(person.Birtday2())

	//account := bankAccount{120}
	//fmt.Println(account.GetBalance())
	//account.Deposit(100)
	//fmt.Println(account.GetBalance())

	//fmt.Println("Task2")
	//animal := Animal{"Dog", "Raph!"}
	//dog := Dog{}
	//animal.Speak()
	//dog.Speak()

	//car := Car{}
	//car.Drive()

	//stc := strc{}
	//stc.dog.Speak()
	//stc.cat.Speak()
	////stc.Speak()

	////fmt.Println("Task3")
	//rect := Rectangle{Width: 10, Height: 5}
	//circ := Circle{Radius: 15}
	//square := Square{}
	////fmt.Println(rect.Area())
	////fmt.Printf("%.2f\n", circ.Area())
	////PrintArea(circ)
	////PrintArea(rect)
	//
	//Describe(rect)
	//Describe(circ)
	//Describe(square)

	fmt.Println("Task4")
	new := NewPerson("Rusik", 99)
	fmt.Println(new)

	srv := NewServer(123, "NEW", true)
	fmt.Println(srv)
	///////////////////////////////
	ctx := context.Background()
	workers := []Worker{
		&FileProcessor{FilePath: "/path/to/file"},
		&NetworkFetcher{URL: "http://127.0.0.1:8080"},
	}

	for _, worker := range workers {
		err := worker.Do(ctx)

		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Всё прошло успешно!")
		}
	}
	////////////////////////////////
	Test()

	//fmt.Println("task 5")
	//list := &List{items: []int{1, 2, 3}}
	//fmt.Println(list.Len())
	//var list2 *List
	//fmt.Println(list2.Len())

	//str := &Str{19}
	//str.Get()
	//fmt.Println(str.val)
	//str.Increment()
	//str.Get()
	//fmt.Println(str.val)
	//str.Increment()
	//str.Get()
	//fmt.Println(str.val)
	//str.Set(2026)
	//str.Get()
	//fmt.Println(str.val)
	//
	////nums := []int{1,2,3}
	////PrintAll(nums)//Ошибка компиляции
	//
	//items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//interf := IntToInterface(items)
	//fmt.Println(interf)
	//
	//file := File{Name: "Text", Content: "Hello World", isOpen: true}
	//
	//var rc ReadCloser = &file
	//
	//fmt.Println(rc.Read())
	//rc.Close()
	//fmt.Println(rc.Read())

}
