package internal

import (
	_ "github.com/dolmen-go/hid/internal/hidapi"
	_ "github.com/dolmen-go/hid/internal/hidapi/mac"
	_ "github.com/dolmen-go/hid/internal/hidapi/hidapi"
	_ "github.com/dolmen-go/hid/internal/hidapi/windows"
	_ "github.com/dolmen-go/hid/internal/hidapi/libusb"
	_ "github.com/dolmen-go/hid/internal/libusb"
	_ "github.com/dolmen-go/hid/internal/libusb/libusb"
	_ "github.com/dolmen-go/hid/internal/libusb/libusb/os"
)
