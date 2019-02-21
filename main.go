package main

import (
	"log"
	"strconv"
	"syscall/js"
)

func operands(i []js.Value) (a, b int) {
	rawA := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	rawB := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	a, err := strconv.Atoi(rawA)
	if err != nil {
		log.Println(err)
		return
	}
	b, err = strconv.Atoi(rawB)
	if err != nil {
		log.Println(err)
		return
	}

	return a, b
}

func add(i []js.Value) {
	a, b := operands(i)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", a+b)
}

func subtract(i []js.Value) {
	a, b := operands(i)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", a-b)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	// register functions
	registerCallbacks()
	<-c
}
