// Copyright © 2018 Julien SENON <julien.senon@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jsenon/vpncentralmanager/internal/restapi"
	"go.opencensus.io/zpages"
)

const (
	port = ":9010"
)

// ServeRest start API Rest Server
func ServeRest() {
	fmt.Println("Start Rest z-Page Server")
	go func() {
		mux := http.DefaultServeMux
		zpages.Handle(mux, "/")
		log.Fatal(http.ListenAndServe("127.0.0.1:7777", mux))
	}()
	fmt.Println("Start Rest Server")
	fmt.Println("Listening REST on port: ", port)

	// API Part
	http.HandleFunc("/healthz", restapi.Health)
	http.HandleFunc("/.well-known", restapi.WellKnownFingerHandler)

	// http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}