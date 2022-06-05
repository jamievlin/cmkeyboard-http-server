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
	"encoding/json"
	"errors"
	"fmt"
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteOutputMsg(writer http.ResponseWriter, message []byte) {
	n, err := writer.Write(message)
	if err != nil {
		Loggers.ErrorLogger.Fatal("Cannot write response!")
	} else {
		Loggers.InfoLogger.Printf("Wrote %d byte for message", n)
	}
}

func ReturnError(writer http.ResponseWriter, response *ErrorResponse, status int) {
	writer.WriteHeader(status)

	val, err := json.Marshal(response)
	if err != nil {
		Loggers.ErrorLogger.Print("Error marshaling JSON response!")
		return
	}

	WriteOutputMsg(writer, val)
}

func RetrieveDeviceIndexOrLog(dev string, writer http.ResponseWriter) (DeviceIndex, error) {
	devInt, err := GetDeviceIndexFromString(dev)
	if err != nil {
		var errorMsg = fmt.Sprintf("Device %s unknown", dev)
		Loggers.ErrorLogger.Print(errorMsg)
		ReturnError(writer, &ErrorResponse{Message: errorMsg}, http.StatusBadRequest)
		return 0, errors.New("cannot retrieve index")
	}

	return devInt, nil
}
