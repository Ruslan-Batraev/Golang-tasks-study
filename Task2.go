package main

import "fmt"

type Animal struct {
	name  string
	speak string
}

func (a *Animal) Speak() {
	fmt.Println(a.speak)
}

type Dog struct {
	animal Animal
}

func (d *Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct {
	animal Animal
}

func (c *Cat) Speak() {
	fmt.Println("Meow!")
}

type Engine struct {
}

func (e *Engine) Start() {
	fmt.Println("Engine Start")
}

type Car struct {
	engine Engine
}

func (c *Car) Drive() {
	c.engine.Start()
}

type strc struct {
	dog Dog
	cat Cat
}
