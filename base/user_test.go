// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package base

import (
	"testing"

	"github.com/hexya-erp/hexya/hexya/models"
	"github.com/hexya-erp/hexya/hexya/models/security"
	"github.com/hexya-erp/hexya/hexya/tests"
	"github.com/hexya-erp/hexya/pool/h"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	tests.RunTests(m, "base")
}

func TestUserAuthentication(t *testing.T) {
	Convey("Testing User Authentication", t, func() {
		models.SimulateInNewEnvironment(security.SuperUserID, func(env models.Environment) {
			userJohn := h.User().Create(env, &h.UserData{
				Name:     "John Smith",
				Login:    "jsmith",
				Password: "secret",
			})
			Convey("Correct user authentication", func() {
				uid, err := h.User().NewSet(env).Authenticate("jsmith", "secret")
				So(uid, ShouldEqual, userJohn.ID())
				So(err, ShouldBeNil)
			})
			Convey("Invalid credentials authentication", func() {
				uid, err := h.User().NewSet(env).Authenticate("jsmith", "wrong-secret")
				So(uid, ShouldEqual, 0)
				So(err, ShouldHaveSameTypeAs, security.InvalidCredentialsError(""))
			})
			Convey("Unknown user authentication", func() {
				uid, err := h.User().NewSet(env).Authenticate("jsmith2", "wrong-secret")
				So(uid, ShouldEqual, 0)
				So(err, ShouldHaveSameTypeAs, security.UserNotFoundError(""))
			})
		})
	})
}
