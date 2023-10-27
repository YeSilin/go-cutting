package main

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
)

func main33() {
	hkey := hotkey.New()

	quit := make(chan bool)

	hkey.Register(hotkey.None, hotkey.F1, func() {
		fmt.Println("Quit~~~~~~~~~~~~~~~~~~~")
		//quit <- true
	})

	hkey.Register(hotkey.None, hotkey.F2, func() {
		fmt.Println("f2~~~~~~~~~~~~~~~~~~~")
		//quit <- true
	})

	fmt.Println("Start hotkey's loop")
	fmt.Println("Push Ctrl-Q to escape and quit")
	<-quit
}
