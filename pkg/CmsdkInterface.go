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

package CmsdkInterface

// #cgo CFLAGS: -I${SRCDIR}/../third_party/cmsdk/include/
// #cgo LDFLAGS: -L${SRCDIR}/../third_party/cmsdk/lib/ -lSDKDLL64
// #include "SDKDLL.h"
import "C"
import (
	"fmt"
)

func GetCMSDKDllVer() int {
	return int(C.GetCM_SDK_DllVer())
}

func GetRamUsage() uint32 {
	return uint32(C.GetRamUsage())
}

func EnableLedControl(enabled bool, deviceIndex uint) error {
	ret := bool(C.EnableLedControl(C.bool(enabled), C.DEVICE_INDEX(deviceIndex)))
	if ret {
		return nil
	} else {
		return fmt.Errorf("setting EnableLedControl failed on device %d", deviceIndex)
	}
}

func SetFullLedColor(r byte, g byte, b byte, deviceIndex uint) error {
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
