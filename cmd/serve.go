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

package cmd

import (
	"fmt"
	"log"
	"os"

	s "github.com/jsenon/vpncentralmanager/pkg/grpc/server"
	"github.com/jsenon/vpncentralmanager/pkg/rest"
	"github.com/spf13/cobra"
)

var url string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launch VPN Server",
	Long: `Launch VPN Server 
           which manage config file generation on VPN Servers
           `,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Setenv("urldynamo", url)
		if err != nil {
			log.Fatalf("Error setenv: %v", err)
		}
		fmt.Println("Dynamo url", os.Getenv("urldynamo"))
		Start()
	},
}

func init() {
	serveCmd.PersistentFlags().StringVar(&url, "url", "http://localhost:8000", "url:port for dynamoDB")
	rootCmd.AddCommand(serveCmd)

}

// Start the server
func Start() {
	// GRPC Server
	go s.Serve()
	// REST Server
	rest.ServeRest()
}
