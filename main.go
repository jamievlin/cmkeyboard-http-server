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

package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	CmsdkInterface "jamievlin.github.io/cmkeyboard-http-server/pkg"
	"log"
	"net/http"
	"strconv"
)

func rootHandler(
	writer http.ResponseWriter,
	_ *http.Request,
	_ httprouter.Params) {
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("{}"))
	if err != nil {
		Loggers.ErrorLogger.Print("rootHandler: ", err)
	}
}

type putDeviceLedControlBody struct {
	Enabled bool `json:"enabled"`
}

func putDeviceLedControl(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := strconv.Atoi(dev)
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

	CmsdkInterface.EnableLedControl(bodyParsed.Enabled, uint(devInt))

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		Loggers.ErrorLogger.Fatal("Cannot write response!")
	}
}

func putDeviceColor(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	var dev = params.ByName("device")
	devInt, err := strconv.Atoi(dev)
	if err != nil {
		Loggers.ErrorLogger.Printf("Device %s unknown!", dev)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("{}"))
		if err != nil {
			Loggers.ErrorLogger.Fatal("Cannot write response!", err)
		}
		return
	}

	CmsdkInterface.SetFullLedColor(255, 255, 255, uint(devInt))

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		Loggers.ErrorLogger.Fatal("Cannot write response!")
	}
}

func initMux(router *httprouter.Router, prefix string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	return mux
}

func main() {
	defer CmsdkInterface.EnableLedControl(false, 1)

	Loggers.InfoLogger.Printf("SDK Version %d", CmsdkInterface.GetCMSDKDllVer())

	router := httprouter.New()
	router.GET("/hello", rootHandler)
	router.PUT("/devices/:device", putDeviceLedControl)
	router.PUT("/devices/:device/color", putDeviceColor)
	mux := initMux(router, "/api/v1")

	err := http.ListenAndServe(":10007", mux)
	if err != nil {
		Loggers.ErrorLogger.Fatal(err)
	}
}
