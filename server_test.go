// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	inventory "github.com/mendersoftware/inventory/inv"
	mstore "github.com/mendersoftware/inventory/store/mocks"
)

func TestSetupApi(t *testing.T) {
	// expecting an error
	api, err := SetupAPI("foo")
	assert.Nil(t, api)
	assert.Error(t, err)

	api, err = SetupAPI(EnvDev)
	assert.NotNil(t, api)
	assert.Nil(t, err)
}

func TestMaybeWithInventory(t *testing.T) {
	db := &mstore.DataStore{}
	inv := inventory.NewInventory(db)

	conf := viper.New()

	conf.Set(SettingEnableReporting, true)
	_, err := maybeWithInventory(inv, conf)
	assert.EqualError(t, err, "reporting integration needs orchestrator address")

	conf.Set(SettingOrchestratorAddr, "http://mender-workflows:8080")
	_, err = maybeWithInventory(inv, conf)
	assert.Nil(t, err)
}
