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

package pkg

import (
	"fmt"
)

type DeviceIndex uint16

const (
	MKeysL      DeviceIndex = 0
	MKeysS      DeviceIndex = 1
	MKeysLWhite DeviceIndex = 2
	MKeysMWhite DeviceIndex = 3
	MMouseL     DeviceIndex = 4
	MMouseS     DeviceIndex = 5
	MKeysM      DeviceIndex = 6
	MKeysSWhite DeviceIndex = 7
	MM520       DeviceIndex = 8
	MM530       DeviceIndex = 9
	MK750       DeviceIndex = 10
	CK372       DeviceIndex = 11
	CK550and552 DeviceIndex = 12
	CK551       DeviceIndex = 13
	Default     DeviceIndex = 0xffff
)

// from SDKDLL.h
var deviceMap = map[string]DeviceIndex{
	"MasterKeys_L":       MKeysL,
	"MasterKeys_S":       MKeysS,
	"MasterKeys_L_White": MKeysLWhite,
	"MasterKeys_M_White": MKeysMWhite,
	"MasterMouse_L":      MMouseL,
	"MasterMouse_S":      MMouseS,
	"MasterKeys_M":       MKeysM,
	"MasterKeys_S_White": MKeysSWhite,
	"MM520":              MM520,
	"MM530":              MM530,
	"MK750":              MK750,
	"CK372":              CK372,
	"CK550":              CK550and552,
	"CK552":              CK550and552,
	"CK551":              CK551,
	"Default":            Default,
}

func GetDeviceIndexFromDevName(devName string) (DeviceIndex, error) {
	if result, ok := deviceMap[devName]; ok {
		return result, nil
	} else {
		return 0, fmt.Errorf("device %s not found", devName)
	}
}
