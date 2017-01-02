package firmata

import (
	"testing"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/gobottest"
)

var _ gobot.Adaptor = (*TCPAdaptor)(nil)

func initTestTCPAdaptor() *TCPAdaptor {
	a := NewTCPAdaptor("localhost:4567")
	return a
}

func TestBB8Driver(t *testing.T) {
	a := initTestTCPAdaptor()
	gobottest.Assert(t, a.Name(), "TCPFirmata")
}