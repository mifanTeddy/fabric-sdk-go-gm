package factory

import (
	"app/service/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp"
	"app/service/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp/gm"
	"github.com/pkg/errors"
)

const (
	// GuomiBasedFactoryName is the name of the factory of the software-based BCCSP implementation
	GuomiBasedFactoryName = "GM"
)

// GMFactory is the factory of the guomi-based BCCSP.
type GMFactory struct{}

// Name returns the name of this factory
func (f *GMFactory) Name() string {
	return GuomiBasedFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *GMFactory) Get(gmOpts *GmOpts) (bccsp.BCCSP, error) {
	// Validate arguments
	if gmOpts == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	var ks bccsp.KeyStore
	switch {
	case gmOpts.Ephemeral:
		ks = gm.NewDummyKeyStore()
	case gmOpts.FileKeystore != nil:
		fks, err := gm.NewFileBasedKeyStore(nil, gmOpts.FileKeystore.KeyStorePath, false)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to initialize gm software key store")
		}
		ks = fks
	case gmOpts.InmemKeystore != nil:
		ks = gm.NewInMemoryKeyStore()
	default:
		// Default to ephemeral key store
		ks = gm.NewDummyKeyStore()
	}

	return gm.New(gmOpts.SecLevel, gmOpts.HashFamily, ks)
}

// SwOpts contains options for the SWFactory
type GmOpts struct {
	// Default algorithms when not specified (Deprecated?)
	SecLevel   int    `mapstructure:"security" json:"security" yaml:"Security"`
	HashFamily string `mapstructure:"hash" json:"hash" yaml:"Hash"`

	// Keystore Options
	Ephemeral     bool               `mapstructure:"tempkeys,omitempty" json:"tempkeys,omitempty"`
	FileKeystore  *FileKeystoreOpts  `mapstructure:"filekeystore,omitempty" json:"filekeystore,omitempty" yaml:"FileKeyStore"`
	DummyKeystore *DummyKeystoreOpts `mapstructure:"dummykeystore,omitempty" json:"dummykeystore,omitempty"`
	InmemKeystore *InmemKeystoreOpts `mapstructure:"inmemkeystore,omitempty" json:"inmemkeystore,omitempty"`
}

// Pluggable Keystores, could add JKS, P12, etc..
type FileKeystoreOpts struct {
	KeyStorePath string `mapstructure:"keystore" yaml:"KeyStore"`
}

type DummyKeystoreOpts struct{}

// InmemKeystoreOpts - empty, as there is no config for the in-memory keystore
type InmemKeystoreOpts struct{}
