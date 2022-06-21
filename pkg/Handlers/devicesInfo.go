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

type devicesPluggedInInfo struct {
	PluggedIn bool `json:"plugged_in"`
}

func getDevicesPluggedIn(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	pkg.InitResponse(w)

	var dev = params.ByName("device")
	devInt, err := pkg.RetrieveDeviceIndexOrLog(dev, w)
	if err != nil {
		return
	}

	retVal := devicesPluggedInInfo{PluggedIn: CInterface.IsDevicePlug(devInt)}

	jsonData, err := json.Marshal(retVal)
	if err != nil {
		Loggers.ErrorLogger.Printf("Json marshall error for IsDevicePlug with device %d", devInt)
		pkg.ReturnError(w, &pkg.ErrorResponse{Message: "Json marshall error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	pkg.WriteOutputMsg(w, jsonData)
}
