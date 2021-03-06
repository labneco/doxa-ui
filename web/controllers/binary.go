// Copyright 2016 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labneco/doxa/doxa/server"
	"github.com/labneco/doxa/doxa/tools/generate"
)

// CompanyLogo serves the logo of the company
func CompanyLogo(c *server.Context) {
	c.File(filepath.Join(generate.DoxaDir, "doxa", "server", "static", "web", "src", "img", "logo.png"))
}

// Image serves the image stored in the database (base64 encoded)
// in the given model and given field
func Image(c *server.Context) {
	model := c.Query("model")
	field := c.Query("field")
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	uid := c.Session().Get("uid").(int64)
	img, gErr := getFieldValue(uid, id, model, field)
	if gErr != nil {
		c.Error(fmt.Errorf("Unable to fetch image: %s", gErr))
		return
	}
	if img.(string) == "" {
		c.File(filepath.Join(generate.DoxaDir, "doxa", "server", "static", "web", "src", "img", "placeholder.png"))
		return
	}
	res, err := base64.StdEncoding.DecodeString(img.(string))
	if err != nil {
		c.Error(fmt.Errorf("Unable to convert image: %s", err))
		return
	}
	c.Data(http.StatusOK, "image/png", res)
}
