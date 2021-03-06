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

package nextip

import (
	"context"
	"net"

	"go.opencensus.io/trace"
)

// NextIP Increment IP
func NextIP(ctx context.Context, stringip string) string {
	_, span := trace.StartSpan(ctx, "(*Server).NextIP")
	defer span.End()
	ip := net.ParseIP(stringip)
	// make sure it's only 4 bytes
	ip = ip.To4()
	// check ip != nil
	ip[3]++ // check for rollover
	return ip.String()
	//127.1.0.1
}
