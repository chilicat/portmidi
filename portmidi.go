package portmidi

import (
	"errors"

	"github.com/xlab/portmidi/pm"
)

// Init is the library initialisation function: call this before using portmidi.
func Init() error {
	return pm.ToError(pm.Initialize())
}

// Deinit is the library termination function: call this after using portmidi.
func Deinit() error {
	return pm.ToError(pm.Terminate())
}

// GetHostError translates portmidi host error into human readable message.
func GetHostError() error {
	buf := make([]byte, pm.HostErrorMsgLen)
	pm.GetHostErrorText(buf, pm.HostErrorMsgLen)
	for i := range buf {
		if buf[i] == 0 {
			buf = buf[:i]
			break
		}
	}
	return errors.New(string(buf))
}

// CountDevices gets devices count, ids range from 0 to CountDevices()-1.
func CountDevices() int {
	return int(pm.CountDevices())
}

type DeviceID pm.DeviceID

// DefaultOutputDevice returns the default output device ID or ok=false if there are no devices.
func DefaultOutputDevice() (DeviceID, bool) {
	dev := pm.GetDefaultOutputDeviceID()
	if dev == pm.NoDevice {
		return 0, false
	}
	return DeviceID(dev), true
}

// DefaultInputDevice returns the default input device ID or ok=false if there are no devices.
func DefaultInputDevice() (DeviceID, bool) {
	dev := pm.GetDefaultInputDeviceID()
	if dev == pm.NoDevice {
		return 0, false
	}
	return DeviceID(dev), true
}

type DeviceInfo struct {
	// Interface specifies underlying MIDI API, e.g. MMSystem or DirectX.
	Interface string
	// Name is a device name, e.g. USB MidiSport 1x1
	Name string
	// Input true iff input is available.
	Input bool
	// Output true iff output is available.
	Output bool
}

// GetDeviceInfo returns device info for the provided device ID, or nil if ID is out of range.
func GetDeviceInfo(id DeviceID) *DeviceInfo {
	info := pm.GetDeviceInfo(pm.DeviceID(id))
	if info == nil {
		return nil
	}
	info.Deref()
	return &DeviceInfo{
		Interface: info.Interf,
		Name:      info.Name,
		Input:     info.Input > 0,
		Output:    info.Output > 0,
	}
}

type Event struct {
	Timestamp int32
	Message   Message
	SysExData []byte
}

// NewMessage encodes a short MIDI message into a 32-bit word. If data1
// and/or data2 are not present, use zero.
func NewMessage(status, data1, data2 byte) Message {
	return (Message(data2)<<16)&0xFF0000 |
		(Message(data1)<<8)&0xFF00 | Message(status)&0xFF
}

type Message int32

func (m Message) Status() byte {
	return byte(m & 0xFF)
}

func (m Message) Data1() byte {
	return byte((m >> 8) & 0xFF)
}

func (m Message) Data2() byte {
	return byte((m >> 16) & 0xFF)
}