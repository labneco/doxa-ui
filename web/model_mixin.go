// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/labneco/doxa/doxa/models"
	"github.com/labneco/doxa/pool/h"
	"github.com/labneco/doxa/pool/q"
)

func init() {

	h.ModelMixin().Methods().ToggleActive().DeclareMethod(
		`ToggleActive toggles the Active field of this object if it exists.`,
		func(rs h.BaseMixinSet) {
			_, exists := rs.Model().Fields().Get("active")
			if !exists {
				return
			}
			if rs.Get("Active").(bool) {
				rs.Set("Active", false)
			} else {
				rs.Set("Active", true)
			}
		})

	h.ModelMixin().Methods().Search().Extend("",
		func(rs h.ModelMixinSet, cond q.ModelMixinCondition) h.ModelMixinSet {
			activeField, exists := rs.Model().Fields().Get("active")
			activeTest := !rs.Env().Context().HasKey("active_test") || rs.Env().Context().GetBool("active_test")
			if !exists || !activeTest || cond.HasField(activeField) {
				return rs.Super().Search(cond)
			}
			activeCond := q.ModelMixinCondition{
				Condition: models.Registry.MustGet(rs.ModelName()).Field("active").Equals(true),
			}
			cond = cond.AndCond(activeCond)
			return rs.Super().Search(cond)
		})

	h.ModelMixin().Methods().SearchAll().Extend("",
		func(rs h.ModelMixinSet) h.ModelMixinSet {
			_, exists := rs.Model().Fields().Get("active")
			activeTest := !rs.Env().Context().HasKey("active_test") || !rs.Env().Context().GetBool("active_test")
			if !exists || !activeTest {
				return rs.Super().SearchAll()
			}
			activeCond := q.ModelMixinCondition{
				Condition: models.Registry.MustGet(rs.ModelName()).Field("active").Equals(true),
			}
			return rs.Search(activeCond)
		})
}
