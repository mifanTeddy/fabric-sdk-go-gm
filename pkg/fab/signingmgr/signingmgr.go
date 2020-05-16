/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signingmgr

import (
	"app/service/fabric-sdk-go-gm/pkg/common/logging"
	"app/service/fabric-sdk-go-gm/pkg/common/providers/core"

	"app/service/fabric-sdk-go-gm/pkg/core/cryptosuite"
	"github.com/pkg/errors"
)

var logger = logging.NewLogger("fabsdk/fab")

// SigningManager is used for signing objects with private key
type SigningManager struct {
	cryptoProvider core.CryptoSuite
	hashOpts       core.HashOpts
	signerOpts     core.SignerOpts
}

// New Constructor for a signing manager.
// @param {BCCSP} cryptoProvider - crypto provider
// @param {Config} config - configuration provider
// @returns {SigningManager} new signing manager
func New(cryptoProvider core.CryptoSuite) (*SigningManager, error) {
	// modify by bryan. The hash algorithm must same as blockchain with SM3
	return &SigningManager{cryptoProvider: cryptoProvider, hashOpts: cryptosuite.GetSM3Opts()}, nil
}

// Sign will sign the given object using provided key
func (mgr *SigningManager) Sign(object []byte, key core.Key) ([]byte, error) {

	if len(object) == 0 {
		return nil, errors.New("object (to sign) required")
	}

	if key == nil {
		return nil, errors.New("key (for signing) required")
	}

	digest, err := mgr.cryptoProvider.Hash(object, mgr.hashOpts)
	if err != nil {
		return nil, err
	}
	signature, err := mgr.cryptoProvider.Sign(key, digest, mgr.signerOpts)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
