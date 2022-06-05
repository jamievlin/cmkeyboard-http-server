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
	"log"
	"net/http"
)

type putDeviceLedControlBody struct {
	Enabled bool `json:"enabled"`
}

func putDeviceColor(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := pkg.GetDeviceIndexFromString(dev)
	if err != nil {
		Loggers.ErrorLogger.Printf("Device %s unknown!", dev)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("{}"))
		if err != nil {
			Loggers.ErrorLogger.Fatal("Cannot write response!", err)
		}
		return
	}

	CInterface.SetFullLedColor(255, 255, 255, devInt)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		Loggers.ErrorLogger.Fatal("Cannot write response!")
	}
}

func putDeviceLedControl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := pkg.GetDeviceIndexFromString(dev)
	if err != nil {
		Loggers.ErrorLogger.Printf("Device %s unknown!", dev)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("{}"))
		if err != nil {
			Loggers.ErrorLogger.Fatal("Cannot write response!", err)
		}
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("{}"))
		if err != nil {
			Loggers.ErrorLogger.Fatal("Cannot write response!", err)
		}
		return
	}

	var bodyParsed putDeviceLedControlBody
	if json.Unmarshal(body, &bodyParsed) != nil {
		log.Fatal("error!")
	}

	CInterface.EnableLedControl(bodyParsed.Enabled, devInt)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		Loggers.ErrorLogger.Fatal("Cannot write response!")
	}
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
