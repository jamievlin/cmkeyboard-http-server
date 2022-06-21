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
	"strconv"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func InitResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Content-Encoding", "UTF-8")
}

func InitResponseWithBody(w *http.ResponseWriter) {
	InitResponse(w)
	(*w).Header().Set("Accept", "application/json")
	(*w).Header().Set("Accept-Encoding", "UTF-8")
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

const (
	MaxLedRow    int = 7
	MaxLedColumn int = 24
)

func validateRowColumn(row int, column int) bool {
	rowValid := row >= 0 && row < MaxLedRow
	colValid := column >= 0 && column < MaxLedColumn

	return rowValid && colValid
}

func ParseRowColOrLog(row string, column string, writer http.ResponseWriter) (int, int, error) {
	rowInt, err := strconv.Atoi(row)
	colInt, errCol := strconv.Atoi(column)

	if err != nil || errCol != nil {
		var errorMsg = fmt.Sprintf("cannot parse (row, index) = (%s, %s)", row, column)
		Loggers.ErrorLogger.Print(errorMsg)
		ReturnError(writer, &ErrorResponse{Message: errorMsg}, http.StatusBadRequest)
		return 0, 0, errors.New("cannot retrieve index")
	}

	if !validateRowColumn(rowInt, colInt) {
		var errorMsg = fmt.Sprintf("(row, index) of (%d, %d) is not valid!", rowInt, colInt)
		Loggers.ErrorLogger.Print(errorMsg)
		ReturnError(writer, &ErrorResponse{Message: errorMsg}, http.StatusBadRequest)
		return 0, 0, errors.New("invalid row/column index")
	}

	return rowInt, colInt, nil
}
