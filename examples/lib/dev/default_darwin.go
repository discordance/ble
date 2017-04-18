package dev

import (
	"github.com/discordance/ble"
	"github.com/discordance/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
