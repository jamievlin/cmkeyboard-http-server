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
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/CInterface"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/Handlers"
	"net/http"
)

func initMux(prefix string) *http.ServeMux {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()
	Handlers.RegisterDeviceHandler(apiMux)

	mux.Handle(prefix+"/", http.StripPrefix(prefix, apiMux))
	return mux
}

func main() {
	defer CInterface.EnableLedControl(false, 1)

	Loggers.InfoLogger.Printf("SDK Version %d", CInterface.GetCMSDKDllVer())

	mux := initMux("/api/v1")
	err := http.ListenAndServe(":10007", mux)
	if err != nil {
		Loggers.ErrorLogger.Fatal(err)
	}
}
