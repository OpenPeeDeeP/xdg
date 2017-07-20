// +build darwin

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
	setDefaulter(new(osDefaulter))
	assert := assert.New(t)
	homeDir := "/some/path"
	expected := homeDir + "/.local/share"
	os.Setenv("HOME", homeDir) // nolint: errcheck

	actual := defaulter.defaultDataHome()
	assert.Equal(expected, actual)
}

func TestDefaultDataDirs(t *testing.T) {
	setDefaulter(new(osDefaulter))
	assert := assert.New(t)
	expected := []string{"/Library/Application Support"}

	actual := defaulter.defaultDataDirs()
	assert.Equal(expected, actual)
}

func TestDefaultConfigHome(t *testing.T) {
	setDefaulter(new(osDefaulter))
	assert := assert.New(t)
	homeDir := "/some/path"
	expected := homeDir + "/.config"
	os.Setenv("HOME", homeDir) // nolint: errcheck

	actual := defaulter.defaultConfigHome()
	assert.Equal(expected, actual)
}

func TestDefaultConfigDirs(t *testing.T) {
	setDefaulter(new(osDefaulter))
	assert := assert.New(t)
	expected := []string{"/Library/Application Support"}

	actual := defaulter.defaultConfigDirs()
	assert.Equal(expected, actual)
}

func TestDefaultCacheHome(t *testing.T) {
	setDefaulter(new(osDefaulter))
	assert := assert.New(t)
	homeDir := "/some/path"
	expected := homeDir + "/.cache"
	os.Setenv("HOME", homeDir) // nolint: errcheck

	actual := defaulter.defaultCacheHome()
	assert.Equal(expected, actual)
}
