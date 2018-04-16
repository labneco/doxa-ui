// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"github.com/labneco/doxa/doxa/models"
	"github.com/labneco/doxa/doxa/models/security"
	"github.com/labneco/doxa/doxa/server"
	"github.com/labneco/doxa/doxa/tools/logging"
	"github.com/labneco/doxa/pool/h"
)

const (
	// MODULE_NAME is the name of this module
	MODULE_NAME = "base"
)

var log *logging.Logger

func init() {
	log = logging.GetLogger("base")
	server.RegisterModule(&server.Module{
		Name: MODULE_NAME,
		PostInit: func() {
			err := models.ExecuteInNewEnvironment(security.SuperUserID, func(env models.Environment) {
				h.Group().NewSet(env).ReloadGroups()
			})
			if err != nil {
				log.Panic("Error while initializing", "error", err)
			}
		},
	})
}
