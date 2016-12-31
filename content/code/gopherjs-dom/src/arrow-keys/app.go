package main

import "github.com/gopherjs/gopherjs/js"

const (
	LEFT  = 37
	UP    = 38
	RIGHT = 39
	DOWN  = 40
)

func main() {
	info := js.Global.Get("document").Call("getElementById", "info")

	js.Global.Call("addEventListener", "keyup", func(event *js.Object) {
		keycode := event.Get("keyCode").Int()
		if keycode == LEFT {
			info.Set("innerHTML", "left arrow key")
		}
		if keycode == UP {
			info.Set("innerHTML", "up arrow key")
		}
		if keycode == RIGHT {
			info.Set("innerHTML", "right arrow key")
		}
		if keycode == DOWN {
			info.Set("innerHTML", "down arrow key")
		}
	}, false)
}
