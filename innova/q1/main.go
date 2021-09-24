package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名欄位
	school string
	loan   float32
}

type Employee struct {
	Human   // 匿名欄位
	company string
	money   float32
}

// Human 實現 SayHi 方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 實現 Sing 方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 過載 Human 的 SayHi 方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Interface Men 被 Human, Student 和 Employee 實現
// 因為這三個型別都實現了這兩個方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}

	var i Men

	i = paul
	i.SayHi()
	i.Sing("November rain")

	i = sam
	i.SayHi()
	i.Sing("November rain")
}
