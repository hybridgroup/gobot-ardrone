package gobotArdrone

import (
	"github.com/hybridgroup/go-ardrone/client"
	"github.com/hybridgroup/gobot"
)

type ArdroneAdaptor struct {
	gobot.Adaptor
	drone *ardrone.Client
}

var ardroneConnect = func() (*ardrone.Client, error) {
	return ardrone.Connect(ardrone.DefaultConfig())
}

func (me *ArdroneAdaptor) Connect() bool {
	drone, err := ardroneConnect()
	if err != nil {
		panic(err)
	}
	me.drone = drone
	return true
}

func (me *ArdroneAdaptor) Reconnect() bool {
	return true
}

func (me *ArdroneAdaptor) Disconnect() bool {
	return true
}

func (me *ArdroneAdaptor) Finalize() bool {
	return true
}
