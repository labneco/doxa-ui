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

package web

import (
	// Loading dependencies as blank imports
	_ "github.com/labneco/doxa-ui/base"
	// Loading controllers package
	_ "github.com/labneco/doxa-ui/web/controllers"
	"github.com/labneco/doxa/doxa/server"
	"github.com/labneco/doxa/doxa/tools/logging"
)

const (
	// MODULE_NAME is the name of this module.
	// It is used by the framework to recognize a module from a simple go package.
	MODULE_NAME = "web"
)

var log *logging.Logger

func init() {
	log = logging.GetLogger("web")
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PostInit: func() {},
	})
}
