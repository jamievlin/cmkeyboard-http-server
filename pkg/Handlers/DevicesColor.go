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
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"jamievlin.github.io/cmkeyboard-http-server/pkg"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/CInterface"
	"net/http"
)

func inByteRange(val int) bool {
	return val >= 0 && val <= 255
}

// methods+types for putDeviceColor

type RGBColor struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

func (body RGBColor) Validate() bool {
	return inByteRange(body.Red) && inByteRange(body.Green) && inByteRange(body.Blue)
}

func (body RGBColor) toBytes() (byte, byte, byte) {
	return byte(body.Red), byte(body.Green), byte(body.Blue)
}

func createRGBColor(data *interface{}) (*RGBColor, bool) {
	var ret RGBColor
	res, ok := (*data).(map[string]interface{})
	if !ok {
		return nil, false
	}

	vr, okr := res["red"]
	vg, okg := res["green"]
	vb, okb := res["blue"]

	if !(okr && okg && okb) {
		return nil, false
	}

	rv, okrv := vr.(float64)
	rg, okrg := vg.(float64)
	rb, okrb := vb.(float64)

	if !(okrv && okrg && okrb) {
		return nil, false
	}

	ret.Red = int(rv)
	ret.Green = int(rg)
	ret.Blue = int(rb)
	return &ret, true
}

func putDeviceColor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := pkg.RetrieveDeviceIndexOrLog(dev, w)
	if err != nil {
		return
	}

	bodyParsed, err := pkg.ReadResponseOrLog[map[string]interface{}](w, r)
	if err != nil {
		return
	}

	mode, ok := (*bodyParsed)["mode"]
	rawBody, okb := (*bodyParsed)["body"]
	if !(ok && okb) {
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Mode is required"}, http.StatusBadRequest)
		return
	}

	result := false

	if mode == "all" {
		body, res := createRGBColor(&rawBody)
		if !res {
			pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Cannot parse body!"}, http.StatusInternalServerError)
			return
		}
		result = true
		red, green, blue := body.toBytes()

		if CInterface.SetFullLedColor(red, green, blue, devInt) != nil {
			pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Cannot set Full LED"}, http.StatusInternalServerError)
			return
		}
	}

	if !result {
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Mode must be all or matrix!"}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, []byte("{}"))
}

// method for putDeviceLedControl

type putDeviceLedControlBody struct {
	Enabled bool `json:"enabled"`
}

func (body putDeviceLedControlBody) Validate() bool {
	return true
}

func putDeviceLedControl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := pkg.RetrieveDeviceIndexOrLog(dev, w)
	if err != nil {
		return
	}

	bodyParsed, err := pkg.ReadValidatedResponseOrLog[putDeviceLedControlBody](w, r)
	if err != nil {
		return
	}

	if CInterface.EnableLedControl(bodyParsed.Enabled, devInt) != nil {
		Loggers.ErrorLogger.Print("Cannot set LED Control")
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Cannot set LED Control"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, []byte("{}"))
}

// register + main functions

func putDeviceKeyColor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	devInt, err := pkg.RetrieveDeviceIndexOrLog(
		params.ByName("device"), w)
	if err != nil {
		return
	}

	row, col, err := pkg.ParseRowColOrLog(
		params.ByName("row"),
		params.ByName("col"),
		w)
	if err != nil {
		return
	}

	bodyParsed, err := pkg.ReadValidatedResponseOrLog[RGBColor](w, r)
	if err != nil {
		return
	}

	red, green, blue := bodyParsed.toBytes()

	if CInterface.SetLedColor(
		row, col,
		red, green, blue,
		devInt) != nil {
		Loggers.ErrorLogger.Print("Cannot set LED Control for key")
		pkg.ReturnError(
			w,
			&pkg.ErrorResponse{Message: "Cannot set LED Control for key"},
			http.StatusInternalServerError)
		return
	}
}

func createDeviceHandler() http.Handler {
	router := httprouter.New()

	router.PUT("/:device", putDeviceLedControl)
	router.PUT("/:device/color", putDeviceColor)
	router.PUT("/:device/color/:row/:col", putDeviceKeyColor)
	return router
}

func RegisterDeviceHandler(mux *http.ServeMux) {
	deviceHandler := createDeviceHandler()
	mux.Handle("/devices/", http.StripPrefix("/devices", deviceHandler))
}
