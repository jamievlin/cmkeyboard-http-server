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
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"jamievlin.github.io/cmkeyboard-http-server/pkg"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/CInterface"
	"net/http"
)

type putDeviceLedControlBody struct {
	Enabled bool `json:"enabled"`
}

func putDeviceColor(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := pkg.RetrieveDeviceIndexOrLog(dev, w)
	if err != nil {
		return
	}

	if CInterface.SetFullLedColor(255, 255, 255, devInt) != nil {
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Cannot set Full LED"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, []byte("{}"))
}

func putDeviceLedControl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := pkg.RetrieveDeviceIndexOrLog(dev, w)
	if err != nil {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: ""}, http.StatusInternalServerError)
		return
	}

	var bodyParsed putDeviceLedControlBody
	if json.Unmarshal(body, &bodyParsed) != nil {
		Loggers.ErrorLogger.Print("Cannot unmarshal response")
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Cannot unmarshal response"}, http.StatusBadRequest)
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

func createDeviceHandler() http.Handler {
	router := httprouter.New()

	router.PUT("/:device", putDeviceLedControl)
	router.PUT("/:device/color", putDeviceColor)
	return router
}

func RegisterDeviceHandler(mux *http.ServeMux) {
	deviceHandler := createDeviceHandler()
	mux.Handle("/devices/", http.StripPrefix("/devices", deviceHandler))
}
