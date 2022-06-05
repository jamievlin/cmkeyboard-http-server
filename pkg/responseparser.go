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
	"io"
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"net/http"
)

type ValidableResponse interface {
	Validate() bool
}

func ReadValidatedResponseOrLog[K ValidableResponse](w http.ResponseWriter, r *http.Request) (*K, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnError(w, &ErrorResponse{Message: "Cannot read message!"}, http.StatusInternalServerError)
		return nil, errors.New("message read error")
	}

	var bodyParsed K
	if json.Unmarshal(body, &bodyParsed) != nil || !bodyParsed.Validate() {
		Loggers.ErrorLogger.Print("Invalid Response!")
		ReturnError(w, &ErrorResponse{Message: "Invalid response"}, http.StatusBadRequest)
		return nil, errors.New("parse error")
	}

	return &bodyParsed, nil
}
