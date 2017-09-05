// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 05 Sep 2017 19:52:40 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package pm

/*
#cgo LDFLAGS: -lportmidi
#include "portmidi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// PortMidiStream type as declared in pm/portmidi.h:163
type PortMidiStream [0]byte

// DeviceID type as declared in pm/portmidi.h:205
type DeviceID int32

// DeviceInfo as declared in pm/portmidi.h:215
type DeviceInfo struct {
	StructVersion  int32
	Interf         string
	Name           string
	Input          int32
	Output         int32
	Opened         int32
	refbc1771f0    *C.PmDeviceInfo
	allocsbc1771f0 interface{}
}

// Timestamp type as declared in pm/portmidi.h:269
type Timestamp int32

// TimeProcPtr type as declared in pm/portmidi.h:270
type TimeProcPtr func(timeInfo unsafe.Pointer) Timestamp

// Message type as declared in pm/portmidi.h:512
type Message int32

// Event as declared in pm/portmidi.h:581
type Event struct {
	Message        Message
	Timestamp      Timestamp
	ref369ef8f3    *C.PmEvent
	allocs369ef8f3 interface{}
}
