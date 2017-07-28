// Copyright 2017 Istio Authors
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

package na

import (
	"google.golang.org/grpc/credentials"
	pb "istio.io/auth/proto"
	"istio.io/auth/utils"
)

type onPremPlatfromImpl struct {
}

func (na *onPremPlatfromImpl) getTransportCredentials(cfg *Config) credentials.TransportCredentials {
	return utils.GetTLSCredentials(*cfg.NodeIdentityCertFile,
		*cfg.NodeIdentityPrivateKeyFile,
		*cfg.RootCACertFile, true /* isClient */)
}

func (na *onPremPlatfromImpl) getNodeAgentCredentials(cfg *Config) *pb.NodeAgentCredentials {
	return nil
}