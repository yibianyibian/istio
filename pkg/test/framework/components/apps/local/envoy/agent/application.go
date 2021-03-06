//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package agent

import (
	"io"

	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/test/protocol"
)

// Application represents a locally running application with exposed ports.
type Application interface {
	io.Closer
	// GetPorts provides a list of ports that are actively listening for the application.
	GetPorts() model.PortList
}

// ApplicationFactory is a function that manufactures a running application.
type ApplicationFactory func(client protocol.Client) (Application, error)
