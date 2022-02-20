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

// Checksum is a checksum of *.c and *.h files in order to trigger a rebuild
// of package ./.. if one of those files changes.
// See https://github.com/golang/go/issues/26366#issuecomment-405683150
const Checksum = "370c4278f1aa8bc58899384584ee43f946fe21c9000a777d8c4e889138981e20"
