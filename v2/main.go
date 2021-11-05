package lesson3_module

import (
	"fmt"

	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

func Hello() {
	fmt.Println("Hello, World!")
}
