/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"encoding/hex"
	"path/filepath"
	"strings"

	"app/service/fabric-sdk-go-gm/pkg/common/providers/core"
	"app/service/fabric-sdk-go-gm/pkg/common/providers/msp"
	"app/service/fabric-sdk-go-gm/pkg/fab/keyvaluestore"
	"github.com/pkg/errors"
)

// NewFileKeyStore ...
func NewFileKeyStore(cryptoConfigMSPPath string) (core.KVStore, error) {
	opts := &keyvaluestore.FileKeyValueStoreOptions{
		Path: cryptoConfigMSPPath,
		KeySerializer: func(key interface{}) (string, error) {
			pkk, ok := key.(*msp.PrivKeyKey)
			if !ok {
				return "", errors.New("converting key to PrivKeyKey failed")
			}
			if pkk == nil || pkk.MSPID == "" || pkk.ID == "" || pkk.SKI == nil {
				return "", errors.New("invalid key")
			}

			// TODO: refactor to case insensitive or remove eventually.
			r := strings.NewReplacer("{userName}", pkk.ID, "{username}", pkk.ID)
			keyDir := filepath.Join(r.Replace(cryptoConfigMSPPath), "keystore")

			return filepath.Join(keyDir, hex.EncodeToString(pkk.SKI)+"_sk"), nil
		},
	}
	return keyvaluestore.New(opts)
}
