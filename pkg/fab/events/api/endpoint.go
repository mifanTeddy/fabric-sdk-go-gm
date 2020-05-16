/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"app/service/fabric-sdk-go-gm/pkg/common/options"
	"app/service/fabric-sdk-go-gm/pkg/common/providers/fab"
)

// EventEndpoint extends a Peer endpoint and provides the
// event URL, which may or may not be the same as the Peer URL
type EventEndpoint interface {
	fab.Peer

	// Opts returns additional options for the connection
	Opts() []options.Opt
}
