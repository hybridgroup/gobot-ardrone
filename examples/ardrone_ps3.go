package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-ardrone"
	"github.com/hybridgroup/gobot-joystick"
	"math"
)

type pair struct {
	x float64
	y float64
}

func main() {
	joystickAdaptor := new(gobotJoystick.JoystickAdaptor)
	joystickAdaptor.Name = "ps3"
	joystickAdaptor.Params = map[string]interface{}{
		"config": "./examples/dualshock3.json",
	}

	joystick := gobotJoystick.NewJoystick(joystickAdaptor)
	joystick.Name = "ps3"

	ardroneAdaptor := new(gobotArdrone.ArdroneAdaptor)
	ardroneAdaptor.Name = "Drone"

	drone := gobotArdrone.NewArdrone(ardroneAdaptor)
	drone.Name = "Drone"

	work := func() {

		offset := 32767.0
		right_stick := pair{x: 0, y: 0}
		left_stick := pair{x: 0, y: 0}

		gobot.On(joystick.Events["square"], func(data interface{}) {
			drone.TakeOff()
		})
		gobot.On(joystick.Events["triangle"], func(data interface{}) {
			drone.Hover()
		})
		gobot.On(joystick.Events["x"], func(data interface{}) {
			drone.Land()
		})
		gobot.On(joystick.Events["left_x"], func(data interface{}) {
			left_stick.x = float64(data.(int16))
		})
		gobot.On(joystick.Events["left_y"], func(data interface{}) {
			left_stick.y = float64(data.(int16))
		})
		gobot.On(joystick.Events["right_x"], func(data interface{}) {
			right_stick.x = float64(data.(int16))
		})
		gobot.On(joystick.Events["right_y"], func(data interface{}) {
			right_stick.y = float64(data.(int16))
		})

		gobot.Every("0.01s", func() {
			pair := left_stick
			if pair.y < 0 {
				drone.Forward(validatePitch(pair.y, offset))
			} else if pair.y > 0 {
				drone.Backward(validatePitch(pair.y, offset))
			}

			if pair.x > 0 {
				drone.Right(validatePitch(pair.x, offset))
			} else if pair.x < 0 {
				drone.Left(validatePitch(pair.x, offset))
			}
		})

		gobot.Every("0.01s", func() {
			pair := right_stick
			if pair.y < 0 {
				drone.Up(validatePitch(pair.y, offset))
			} else if pair.y > 0 {
				drone.Down(validatePitch(pair.y, offset))
			}

			if pair.x > 20 {
				drone.Clockwise(validatePitch(pair.x, offset))
			} else if pair.x < -20 {
				drone.CounterClockwise(validatePitch(pair.x, offset))
			}
		})

		gobot.Every("0.1s", func() {
			drone.Hover()
		})

	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{joystickAdaptor, ardroneAdaptor},
		Devices:     []gobot.Device{joystick, drone},
		Work:        work,
	}

	robot.Start()
}

func validatePitch(data float64, offset float64) float64 {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return float64(int(value*100)) / 100
		} else {
			return 1.0
		}
	} else {
		return 0.0
	}
}
