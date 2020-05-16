// +build testing

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chpvdr

import (
	"sync"

	"app/service/fabric-sdk-go-gm/pkg/common/options"
	"app/service/fabric-sdk-go-gm/pkg/common/providers/fab"
	"app/service/fabric-sdk-go-gm/pkg/fab/chconfig"
	"app/service/fabric-sdk-go-gm/pkg/util/concurrent/lazycache"
	"app/service/fabric-sdk-go-gm/pkg/util/concurrent/lazyref"
	"github.com/pkg/errors"
)

// SetChannelConfig allows setting channel configuration.
// This method is intended to enable tests and should not be called.
func SetChannelConfig(cfg ...fab.ChannelCfg) {
	provider := newMockChCfgCache(cfg...)
	cfgCacheProvider = func(opts ...options.Opt) cache {
		return provider
	}
}

type chCfgCache struct {
	cfgMap sync.Map
}

func newMockChCfgCache(cfgs ...fab.ChannelCfg) *chCfgCache {
	c := &chCfgCache{}
	for _, cfg := range cfgs {
		c.Put(cfg)
	}
	return c
}

func newChCfgRef(cfg fab.ChannelCfg) *chconfig.Ref {
	r := &chconfig.Ref{}
	r.Reference = lazyref.New(func() (interface{}, error) {
		return cfg, nil
	})
	return r
}

// Get mock channel config reference
func (m *chCfgCache) Get(k lazycache.Key, data ...interface{}) (interface{}, error) {
	channelID := k.(chconfig.CacheKey).ChannelID()
	cfg, ok := m.cfgMap.Load(channelID)
	if !ok {
		return nil, errors.Errorf("Channel config not found in cache for channel: %s", channelID)
	}
	return cfg, nil
}

// Close not implemented
func (m *chCfgCache) Close() {
}

// Put channel config reference into mock cache
func (m *chCfgCache) Put(cfg fab.ChannelCfg) {
	m.cfgMap.Store(cfg.ID(), newChCfgRef(cfg))
}
