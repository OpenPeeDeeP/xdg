// +build windows

// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultDataHome(t *testing.T) {
	assert := assert.New(t)
	appData := "/some/path"
	expected := appData
	os.Setenv("APPDATA", appData) // nolint: errcheck

	actual := defaulter.defaultDataHome()
	assert.Equal(expected, actual)
}

func TestDefaultDataDirs(t *testing.T) {
	assert := assert.New(t)
	programData := "/some/path"
	expected := []string{programData}
	os.Setenv("PROGRAMDATA", programData) // nolint: errcheck

	actual := defaulter.defaultDataDirs()
	assert.Equal(expected, actual)
}

func TestDefaultConfigHome(t *testing.T) {
	assert := assert.New(t)
	appData := "/some/path"
	expected := appData
	os.Setenv("APPDATA", appData) // nolint: errcheck

	actual := defaulter.defaultConfigHome()
	assert.Equal(expected, actual)
}

func TestDefaultConfigDirs(t *testing.T) {
	assert := assert.New(t)
	programData := "/some/path"
	expected := []string{programData}
	os.Setenv("PROGRAMDATA", programData) // nolint: errcheck

	actual := defaulter.defaultConfigDirs()
	assert.Equal(expected, actual)
}

func TestDefaultCacheHome(t *testing.T) {
	assert := assert.New(t)
	appData := "/some/path"
	expected := appData
	os.Setenv("LOCALAPPDATA", appData) // nolint: errcheck

	actual := defaulter.defaultCacheHome()
	assert.Equal(expected, actual)
}
