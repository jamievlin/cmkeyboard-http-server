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

import "strconv"

type DeviceIndex uint16

const (
	MKeysL  DeviceIndex = 0
	MKeysS  DeviceIndex = 1
	MK750   DeviceIndex = 10
	Default DeviceIndex = 0xffff
)

func GetDeviceIndexFromString(input string) (DeviceIndex, error) {
	devInt, err := strconv.Atoi(input)
	return DeviceIndex(devInt), err
}
