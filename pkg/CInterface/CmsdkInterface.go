/*
 *    Copyright 2022 Supakorn 'Jamie' Rassameemasmuang
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package CInterface

// #cgo CFLAGS: -I${SRCDIR}/../../third_party/cmsdk/include/
// #cgo LDFLAGS: -L${SRCDIR}/../../third_party/cmsdk/lib/ -lSDKDLL64
// #include "SDKDLL.h"
import "C"
import (
	"fmt"
	"jamievlin.github.io/cmkeyboard-http-server/pkg"
)

func GetCMSDKDllVer() int {
	return int(C.GetCM_SDK_DllVer())
}

func GetRamUsage() uint32 {
	return uint32(C.GetRamUsage())
}

func EnableLedControl(enabled bool, deviceIndex pkg.DeviceIndex) error {
	ret := bool(C.EnableLedControl(C.bool(enabled), C.DEVICE_INDEX(deviceIndex)))
	if ret {
		return nil
	} else {
		return fmt.Errorf("setting EnableLedControl failed on device %d", deviceIndex)
	}
}

func SetFullLedColor(r byte, g byte, b byte, deviceIndex pkg.DeviceIndex) error {
	ret := bool(C.SetFullLedColor(
		C.BYTE(r),
		C.BYTE(g),
		C.BYTE(b),
		C.DEVICE_INDEX(deviceIndex)),
	)

	if ret {
		return nil
	} else {
		return fmt.Errorf(
			"setting SetFullLedColor to (%d, %d, %d) failed on device %d",
			r, g, b,
			deviceIndex)
	}
}

func SetLedColor(row int, column int, r byte, g byte, b byte, deviceIndex pkg.DeviceIndex) error {
	ret := bool(C.SetLedColor(
		C.int(row), C.int(column),
		C.BYTE(r), C.BYTE(g), C.BYTE(b),
		C.DEVICE_INDEX(deviceIndex),
	))

	if ret {
		return nil
	} else {
		return fmt.Errorf(
			"SetLedColor on key (%d, %d) with color (%d, %d, %d) failed on device (%d)",
			row, column,
			r, g, b,
			deviceIndex)
	}
}

type CmKeyColor struct {
	Red   byte
	Green byte
	Blue  byte
}

type CmColorMatrix [pkg.MaxLedRow][pkg.MaxLedColumn]CmKeyColor

func (km *CmColorMatrix) CreateKeyColor() *C.COLOR_MATRIX {
	var newKm C.COLOR_MATRIX

	for i, kmCol := range km {
		for j, kmEntry := range kmCol {
			entry := &newKm.KeyColor[i][j]
			entry.r = C.BYTE(kmEntry.Red)
			entry.g = C.BYTE(kmEntry.Green)
			entry.b = C.BYTE(kmEntry.Blue)
		}
	}
	return &newKm
}

func SetAllLedColor(colorMatrix *CmColorMatrix, deviceIndex pkg.DeviceIndex) error {
	ret := bool(C.SetAllLedColor(
		*colorMatrix.CreateKeyColor(),
		C.DEVICE_INDEX(deviceIndex),
	))

	if ret {
		return nil
	} else {
		return fmt.Errorf("SetAllLedColor on deviceIndex %d failed", deviceIndex)
	}
}
