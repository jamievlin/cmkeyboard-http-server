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

package Loggers

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var (
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
)

func InitializeLoggerToStderr() {
	initializeLogger(os.Stderr, os.Stderr, os.Stderr)
}

func InitializeLoggerDevNull() {
	fileWriter, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Cannot initialize logger")
	}
	initializeLogger(fileWriter, fileWriter, fileWriter)
}

func InitializeLoggerDirectory(dir string) []*os.File {
	errVal := os.MkdirAll(dir, 0644)
	if errVal != nil {
		log.Fatalf("Cannot create directory %s!", dir)
	}

	createWriter := func(logName string) *os.File {
		dt := time.Now()
		dtPath := fmt.Sprintf("%02d_%02d_%04d", dt.Day(), dt.Month(), dt.Year())
		fileName := fmt.Sprintf("cmhttp_log.%s.%s.log", logName, dtPath)
		fil, err := os.OpenFile(path.Join(dir, fileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatalf("Cannot open file %s for logging!", fileName)
		}
		return fil
	}

	errWriter := createWriter("error")
	warningWriter := createWriter("warning")
	infoWriter := createWriter("info")

	initializeLogger(errWriter, warningWriter, infoWriter)

	return []*os.File{errWriter, warningWriter, infoWriter}

}

// Initializes logger to logdir directory. If null, print to stderr.
func initializeLogger(errWriter io.Writer, warnWriter io.Writer, infoWriter io.Writer) {
	ErrorLogger = log.New(errWriter, "[ERROR] ", log.LstdFlags)
	WarningLogger = log.New(warnWriter, "[INFO] ", log.LstdFlags)
	InfoLogger = log.New(infoWriter, "[INFO] ", log.LstdFlags)
}
