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

package Handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func createDeviceHandler() http.Handler {
	router := httprouter.New()

	devicePath := "/:device"
	router.PUT(devicePath, putDeviceLedControl)
	router.GET(devicePath, getDevicesPluggedIn)
	router.OPTIONS(devicePath, createOptionsHandler(&[]string{"GET", "PUT"}))

	deviceColorPath := "/:device/color"
	router.PUT(deviceColorPath, putDeviceColor)
	router.OPTIONS(deviceColorPath, createOptionsHandler(&[]string{"PUT"}))

	deviceColorRowColPath := "/:device/color/:row/:col"
	router.PUT(deviceColorRowColPath, putDeviceKeyColor)
	router.OPTIONS(deviceColorRowColPath, createOptionsHandler(&[]string{"PUT"}))
	return router
}

func RegisterDeviceHandler(mux *http.ServeMux) {
	deviceHandler := createDeviceHandler()
	mux.Handle("/devices/", http.StripPrefix("/devices", deviceHandler))
}
