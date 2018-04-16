// Copyright 2016 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package controllers

import (
	"net/http"

	"github.com/labneco/doxa/doxa/actions"
	"github.com/labneco/doxa/doxa/server"
)

// CallKW executes the given method of the given model
func CallKW(c *server.Context) {
	uid := c.Session().Get("uid").(int64)
	var params CallParams
	c.BindRPCParams(&params)
	res, err := Execute(uid, params)
	c.RPC(http.StatusOK, res, err)
}

// CallButton executes the given method of the given model
// and returns the result only if it is an action
func CallButton(c *server.Context) {
	uid := c.Session().Get("uid").(int64)
	var params CallParams
	c.BindRPCParams(&params)
	res, err := Execute(uid, params)
	_, isAction := res.(actions.Action)
	_, isActionPtr := res.(*actions.Action)
	if !isAction && !isActionPtr {
		res = false
	}
	c.RPC(http.StatusOK, res, err)
}

// SearchRead returns Records from the database
func SearchRead(c *server.Context) {
	uid := c.Session().Get("uid").(int64)
	var params searchReadParams
	c.BindRPCParams(&params)
	res, err := searchRead(uid, params)
	c.RPC(http.StatusOK, res, err)
}
