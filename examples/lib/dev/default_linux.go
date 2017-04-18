package dev

import (
	"github.com/discordance/ble"
	"github.com/discordance/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
