// Copyright 2016 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labneco/doxa/doxa/actions"
	"github.com/labneco/doxa/doxa/models"
	"github.com/labneco/doxa/doxa/models/security"
	"github.com/labneco/doxa/doxa/models/types"
	"github.com/labneco/doxa/doxa/server"
	"github.com/labneco/doxa/pool/h"
	"github.com/labneco/doxa/pool/q"
)

// ActionLoad returns the action with the given id
func ActionLoad(c *server.Context) {
	var lang string
	if c.Session().Get("uid") != nil {
		models.ExecuteInNewEnvironment(security.SuperUserID, func(env models.Environment) {
			user := h.User().Search(env, q.User().ID().Equals(c.Session().Get("uid").(int64)))
			lang = user.ContextGet().GetString("lang")
		})
	}

	params := struct {
		ActionID          string         `json:"action_id"`
		AdditionalContext *types.Context `json:"additional_context"`
	}{}
	c.BindRPCParams(&params)
	action := *actions.Registry.MustGetById(params.ActionID)
	action.Name = action.TranslatedName(lang)
	c.RPC(http.StatusOK, action)
}

// ActionRun runs the given server action
func ActionRun(c *server.Context) {
	params := struct {
		ActionID string         `json:"action_id"`
		Context  *types.Context `json:"context"`
	}{}
	c.BindRPCParams(&params)
	action := actions.Registry.MustGetById(params.ActionID)

	// Process context ids into args
	var ids []int64
	if params.Context.Get("active_ids") != nil {
		ids = params.Context.Get("active_ids").([]int64)
	} else if params.Context.Get("active_id") != nil {
		ids = []int64{params.Context.Get("active_id").(int64)}
	}
	idsJSON, err := json.Marshal(ids)
	if err != nil {
		log.Panic("Unable to marshal ids")
	}

	// Process context into kwargs
	contextJSON, _ := json.Marshal(params.Context)
	kwargs := make(map[string]json.RawMessage)
	kwargs["context"] = contextJSON

	// Execute the function
	resAction, _ := Execute(c.Session().Get("uid").(int64), CallParams{
		Model:  action.Model,
		Method: action.Method,
		Args:   []json.RawMessage{idsJSON},
		KWArgs: kwargs,
	})

	if _, ok := resAction.(*actions.Action); ok {
		c.RPC(http.StatusOK, resAction)
	} else {
		c.RPC(http.StatusOK, false)
	}
}
