/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package adapters

import (
	"github.com/codoworks/go-boilerplate/pkg/config/features"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"gorm.io/gorm"
)

var Adapters = &Adapter{
	defaultPlatform: constants.DEFAULT_DB_PLATFORM,
	currentPlatform: constants.DEFAULT_DB_PLATFORM,
	adapters:        make(map[string]IAdapter),
}

type IAdapter interface {
	SetConfig(config features.DatabaseConfig)
	GetDriver() (gorm.Dialector, error)
	GetServerDriver() (gorm.Dialector, error)
	GetDSN() (string, error)
	GetServerDSN() (string, error)
	GetDbCreateStatement() (string, error)
	GetDbDropStatement() (string, error)
	ValidateConfig() error
}

type Adapter struct {
	IAdapter
	adapters        map[string]IAdapter
	defaultPlatform string
	currentPlatform string
	config          features.DatabaseConfig
}

func (a *Adapter) SetConfig(config features.DatabaseConfig) {
	a.config = config
	a.currentPlatform = config.Platform
	for _, adapter := range a.adapters {
		adapter.SetConfig(a.config)
	}
}

func (a *Adapter) GetDriver() (gorm.Dialector, error) {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.GetDriver()
	}
	return nil, constants.ERROR_UNKNOWN_DB_PLATFORM
}

func (a *Adapter) GetServerDriver() (gorm.Dialector, error) {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.GetDriver()
	}
	return nil, constants.ERROR_UNKNOWN_DB_PLATFORM
}

func (a *Adapter) GetDSN() (string, error) {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.GetDSN()
	}
	return "", constants.ERROR_UNKNOWN_DB_PLATFORM
}

func (a *Adapter) GetServerDSN() (string, error) {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.GetServerDSN()
	}
	return "", constants.ERROR_UNKNOWN_DB_PLATFORM
}

func (a *Adapter) AppendAdapter(name string, adapter IAdapter) {
	a.adapters[name] = adapter
}

func (a *Adapter) GetDbCreateStatement() (string, error) {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.GetDbCreateStatement()
	}
	return "", constants.ERROR_UNKNOWN_DB_PLATFORM
}

func (a *Adapter) GetDbDropStatement() (string, error) {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.GetDbDropStatement()
	}
	return "", constants.ERROR_UNKNOWN_DB_PLATFORM
}

func (a *Adapter) ValidateConfig() error {
	if adapter, ok := a.adapters[a.currentPlatform]; ok {
		return adapter.ValidateConfig()
	}
	return constants.ERROR_UNKNOWN_DB_PLATFORM
}
