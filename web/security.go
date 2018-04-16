// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/labneco/doxa/doxa/models/security"
	"github.com/labneco/doxa/pool/h"
)

func init() {
	h.Filter().Methods().AllowAllToGroup(security.GroupEveryone)
}
