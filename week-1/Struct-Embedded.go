package main

import "fmt"

type Engine struct{}

func (Engine) start() {
	fmt.Println("Engine started")
}

type Car struct {
	Engine // Car 結構體內嵌了 Engine 結構體
}

func (c Car) driver() {
	// 由於 Car 結構體內嵌了 Engine 結構體，這是一種組合（composition）關係
	// 在 Go 語言中，內嵌的結構體的所有方法會被提升到外層結構體
	// 因此，Car 可以直接訪問 Engine 的方法，就像這些方法是 Car 自己的一樣
	c.start()
	fmt.Println("Car is being driven")
}

func main() {
	// 在 Go 中，當一個結構體被初始化時，其所有欄位（包括匿名欄位）都會被自動初始化為零值。
	// 因此 Car 在建立時可以不用顯式傳入 Engine，因為它的匿名欄位 Engine 會自動初始化為 Engine{}。
	// 這種設計讓結構體的初始化更加簡潔且直觀。
	// 如何避免直接使用零值初始化，就需要引入一個明確的初始化過程
	car := Car{}
	car.driver()
}
