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
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"jamievlin.github.io/cmkeyboard-http-server/pkg"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/CInterface"
	"net/http"
)

type cmsdkVersionResponse struct {
	Version int `json:"version"`
}

func getCmsdkVersion(w http.ResponseWriter, _ *http.Request) {
	pkg.InitResponse(w)

	retBody := cmsdkVersionResponse{Version: CInterface.GetCMSDKDllVer()}

	jsonData, err := json.Marshal(retBody)
	if err != nil {
		Loggers.ErrorLogger.Printf("Json marshall error for getCmsdkVersion")
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Json marshall error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, jsonData)
}

type cmsdkVolumeData struct {
	PeakVolume float32 `json:"peak_volume"`
}

func getPeakVolume(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	pkg.InitResponse(w)

	retBody := cmsdkVolumeData{PeakVolume: CInterface.GetNowVolumePeekValue()}

	jsonData, err := json.Marshal(retBody)
	if err != nil {
		Loggers.ErrorLogger.Printf("Json marshall error for GetNowVolumePeekValue")
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Json marshall error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, jsonData)
}

type cmsdkRamData struct {
	RamPercentage uint32 `json:"ram_percentage"`
}

func getRamData(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	pkg.InitResponse(w)

	retBody := cmsdkRamData{RamPercentage: CInterface.GetRamUsage()}

	jsonData, err := json.Marshal(retBody)
	if err != nil {
		Loggers.ErrorLogger.Printf("Json marshall error for GetNowVolumePeekValue")
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Json marshall error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, jsonData)
}

func createSysHandler() http.Handler {
	router := httprouter.New()

	router.GET("/peakvolume", getPeakVolume)
	router.GET("/ram", getRamData)
	return router
}

// RegisterSysHandler registers /sys/ and /sdkversion HTTP endpointd
func RegisterSysHandler(mux *http.ServeMux) {
	sysHandler := createSysHandler()
	mux.HandleFunc("/sdkversion", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCmsdkVersion(w, r)
		} else {
			pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Method not allowed"}, http.StatusMethodNotAllowed)
		}

	})
	mux.Handle("/sys/", http.StripPrefix("/sys", sysHandler))
}
