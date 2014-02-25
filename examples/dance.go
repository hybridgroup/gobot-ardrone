package main

import (
	"github"
)

func main() {
	drone, err := ardrone.Connect(ardrone.DefaultConfig())
	if err != nil {
		panic(err)
	}

	var flying = make(chan bool)

	go func() {
		flying <- drone.Takeoff()
	}()

	if <-flying {
		drone.Right(0.1)
		time.Sleep(500 * time.Millisecond)
		drone.Hover()
		time.Sleep(500 * time.Millisecond)
		drone.Left(0.1)
		time.Sleep(500 * time.Millisecond)
		drone.Land()
	}
}
