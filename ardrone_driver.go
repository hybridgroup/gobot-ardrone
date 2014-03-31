package gobotArdrone

import (
	"github.com/hybridgroup/go-ardrone/client"
	"github.com/hybridgroup/gobot"
)

type ArdroneDriver struct {
	gobot.Driver
	ArdroneAdaptor *ArdroneAdaptor
	Drone          *ardrone.Client
}

type ArdroneInterface interface {
}

func (me *ArdroneDriver) Start() bool {
	return true
}
func (me *ArdroneDriver) Halt() bool {
	return true
}
func (me *ArdroneDriver) Init() bool {
	return true
}

func NewArdrone(adaptor *ArdroneAdaptor) *ArdroneDriver {
	d := new(ArdroneDriver)
	d.Events = make(map[string]chan interface{})
	d.ArdroneAdaptor = adaptor
	d.Commands = []string{}
	return d
}

func (me *ArdroneDriver) TakeOff() {
	me.Events["Flying"] = make(chan interface{}, 1)
	gobot.Publish(me.Events["Flying"], me.ArdroneAdaptor.drone.Takeoff())
}
func (me *ArdroneDriver) Land() {
	me.ArdroneAdaptor.drone.Land()
}
func (me *ArdroneDriver) Up(a float64) {
	me.ArdroneAdaptor.drone.Up(a)
}
func (me *ArdroneDriver) Down(a float64) {
	me.ArdroneAdaptor.drone.Down(a)
}
func (me *ArdroneDriver) Left(a float64) {
	me.ArdroneAdaptor.drone.Left(a)
}
func (me *ArdroneDriver) Right(a float64) {
	me.ArdroneAdaptor.drone.Right(a)
}
func (me *ArdroneDriver) Forward(a float64) {
	me.ArdroneAdaptor.drone.Forward(a)
}
func (me *ArdroneDriver) Backward(a float64) {
	me.ArdroneAdaptor.drone.Backward(a)
}
func (me *ArdroneDriver) Clockwise(a float64) {
	me.ArdroneAdaptor.drone.Clockwise(a)
}
func (me *ArdroneDriver) CounterClockwise(a float64) {
	me.ArdroneAdaptor.drone.Counterclockwise(a)
}
func (me *ArdroneDriver) Hover() {
	me.ArdroneAdaptor.drone.Hover()
}
