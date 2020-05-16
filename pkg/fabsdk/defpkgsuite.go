/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabsdk

import (
	"app/service/fabric-sdk-go-gm/pkg/core/logging/api"
	"app/service/fabric-sdk-go-gm/pkg/core/logging/modlog"
	sdkApi "app/service/fabric-sdk-go-gm/pkg/fabsdk/api"
	"app/service/fabric-sdk-go-gm/pkg/fabsdk/factory/defcore"
	"app/service/fabric-sdk-go-gm/pkg/fabsdk/factory/defmsp"
	"app/service/fabric-sdk-go-gm/pkg/fabsdk/factory/defsvc"
)

type defPkgSuite struct{}

func (ps *defPkgSuite) Core() (sdkApi.CoreProviderFactory, error) {
	return defcore.NewProviderFactory(), nil
}

func (ps *defPkgSuite) MSP() (sdkApi.MSPProviderFactory, error) {
	return defmsp.NewProviderFactory(), nil
}

func (ps *defPkgSuite) Service() (sdkApi.ServiceProviderFactory, error) {
	return defsvc.NewProviderFactory(), nil
}

func (ps *defPkgSuite) Logger() (api.LoggerProvider, error) {
	return modlog.LoggerProvider(), nil
}
