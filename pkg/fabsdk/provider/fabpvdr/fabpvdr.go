/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabpvdr

import (
	"app/service/fabric-sdk-go-gm/pkg/common/logging"
	"app/service/fabric-sdk-go-gm/pkg/common/providers/context"
	"app/service/fabric-sdk-go-gm/pkg/common/providers/fab"
	"app/service/fabric-sdk-go-gm/pkg/fab/comm"
	"app/service/fabric-sdk-go-gm/pkg/fab/orderer"
	peerImpl "app/service/fabric-sdk-go-gm/pkg/fab/peer"
	"github.com/pkg/errors"
)

var logger = logging.NewLogger("fabsdk")

// InfraProvider represents the default implementation of Fabric objects.
type InfraProvider struct {
	providerContext context.Providers
	commManager     *comm.CachingConnector
}

// New creates a InfraProvider enabling access to core Fabric objects and functionality.
func New(config fab.EndpointConfig) *InfraProvider {
	idleTime := config.Timeout(fab.ConnectionIdle)
	sweepTime := config.Timeout(fab.CacheSweepInterval)

	return &InfraProvider{
		commManager: comm.NewCachingConnector(sweepTime, idleTime),
	}
}

// Initialize sets the provider context
func (f *InfraProvider) Initialize(providers context.Providers) error {
	f.providerContext = providers
	return nil
}

// Close frees resources and caches.
func (f *InfraProvider) Close() {
	logger.Debug("Closing comm manager...")
	f.commManager.Close()
}

// CommManager provides comm support such as GRPC onnections
func (f *InfraProvider) CommManager() fab.CommManager {
	return f.commManager
}

// CreatePeerFromConfig returns a new default implementation of Peer based configuration
func (f *InfraProvider) CreatePeerFromConfig(peerCfg *fab.NetworkPeer) (fab.Peer, error) {
	return peerImpl.New(f.providerContext.EndpointConfig(), peerImpl.FromPeerConfig(peerCfg))
}

// CreateOrdererFromConfig creates a default implementation of Orderer based on configuration.
func (f *InfraProvider) CreateOrdererFromConfig(cfg *fab.OrdererConfig) (fab.Orderer, error) {
	newOrderer, err := orderer.New(f.providerContext.EndpointConfig(), orderer.FromOrdererConfig(cfg))
	if err != nil {
		return nil, errors.WithMessage(err, "creating orderer failed")
	}
	return newOrderer, nil
}
