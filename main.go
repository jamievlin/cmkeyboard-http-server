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
	"flag"
	"fmt"
	Loggers "jamievlin.github.io/cmkeyboard-http-server/internal"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/CInterface"
	"jamievlin.github.io/cmkeyboard-http-server/pkg/Handlers"
	"net/http"
)

type args struct {
	LogDir  *string
	LogMode *string
	Port    *uint
}

func getArgStruct() args {
	return args{
		flag.String("logdir", "logs", "Directory to store log file. Ignored if logmode"),
		flag.String("logmode", "stderr", "Mode to log the file. Must be stderr, file or none"),
		flag.Uint("port", 10007, "Port of the server."),
	}
}

func initMux(prefix string) *http.ServeMux {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	Handlers.RegisterDeviceHandler(apiMux)
	Handlers.RegisterSysHandler(apiMux)

	mux.Handle(prefix+"/", http.StripPrefix(prefix, apiMux))
	return mux
}

func initializeLogger(arg args) {
	if *arg.LogMode == "none" {
		Loggers.InitializeLoggerDevNull()
	} else if *arg.LogMode == "stderr" {
		Loggers.InitializeLoggerToStderr()
	} else if *arg.LogMode == "file" {
		Loggers.InitializeLoggerDirectory(*arg.LogDir)
	}
}

func main() {
	defer func() {
		Loggers.InfoLogger.Print("Program exit.")
		err := CInterface.EnableLedControl(false, 1)
		if err != nil {
			Loggers.ErrorLogger.Printf("Error in disabling LED Control")
		}
	}()

	arg := getArgStruct()
	flag.Parse()

	initializeLogger(arg)

	Loggers.InfoLogger.Printf("SDK Version %d", CInterface.GetCMSDKDllVer())

	mux := initMux("/api/v1")
	addr := fmt.Sprintf("127.0.0.1:%d", *arg.Port)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		Loggers.ErrorLogger.Fatal(err)
	}
}
