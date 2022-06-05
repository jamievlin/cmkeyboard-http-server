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
