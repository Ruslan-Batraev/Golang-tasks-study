package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s!", p.Name)
}

func (p Person) Birthday1() {
	fmt.Println(p.Age + 1)

}

func (p *Person) Birtday2() int {
	return p.Age + 2
}

type bankAccount struct {
	balance float64
}

func (b *bankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *bankAccount) GetBalance() float64 {
	return b.balance
}
