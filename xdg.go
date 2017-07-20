// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xdg impelements the XDG standard for application file locations.
package xdg

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var defaulter xdgDefaulter = new(osDefaulter)
var separator string

func init() {
	if runtime.GOOS == "windows" {
		setSeparator(";")
	} else {
		setSeparator(":")
	}
}

type xdgDefaulter interface {
	defaultDataHome() string
	defaultDataDirs() []string
	defaultConfigHome() string
	defaultConfigDirs() []string
	defaultCacheHome() string
}

type osDefaulter struct {
}

//This method is used in the testing suit
// nolint: deadcode
func setDefaulter(def xdgDefaulter) {
	defaulter = def
}

func setSeparator(sep string) {
	separator = sep
}

// XDG is information about the currently running application
type XDG struct {
	Vendor      string
	Application string
}

// New returns an instance of XDG that is used to grab files for application use
func New(vendor, application string) *XDG {
	return &XDG{
		Vendor:      vendor,
		Application: application,
	}
}

// DataHome returns the location that should be used for user specific data files for this specific application
func (x *XDG) DataHome() string {
	return filepath.Join(DataHome(), x.Vendor, x.Application)
}

// DataDirs returns a list of locations that should be used for system wide data files for this specific application
func (x *XDG) DataDirs() []string {
	dataDirs := DataDirs()
	for i, dir := range dataDirs {
		dataDirs[i] = filepath.Join(dir, x.Vendor, x.Application)
	}
	return dataDirs
}

// ConfigHome returns the location that should be used for user specific config files for this specific application
func (x *XDG) ConfigHome() string {
	return filepath.Join(ConfigHome(), x.Vendor, x.Application)
}

// ConfigDirs returns a list of locations that should be used for system wide config files for this specific application
func (x *XDG) ConfigDirs() []string {
	configDirs := ConfigDirs()
	for i, dir := range configDirs {
		configDirs[i] = filepath.Join(dir, x.Vendor, x.Application)
	}
	return configDirs
}

// CacheHome returns the location that should be used for application cache files for this specific application
func (x *XDG) CacheHome() string {
	return filepath.Join(CacheHome(), x.Vendor, x.Application)
}

// DataHome returns the location that should be used for user specific data files
func DataHome() string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = defaulter.defaultDataHome()
	}
	return dataHome
}

// DataDirs returns a list of locations that should be used for system wide data files
func DataDirs() []string {
	var dataDirs []string
	dataDirsStr := os.Getenv("XDG_DATA_DIRS")
	if dataDirsStr != "" {
		dataDirs = strings.Split(dataDirsStr, separator)
	}
	if len(dataDirs) == 0 {
		dataDirs = defaulter.defaultDataDirs()
	}
	return dataDirs
}

// ConfigHome returns the location that should be used for user specific config files
func ConfigHome() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		configHome = defaulter.defaultConfigHome()
	}
	return configHome
}

// ConfigDirs returns a list of locations that should be used for system wide config files
func ConfigDirs() []string {
	var configDirs []string
	configDirsStr := os.Getenv("XDG_CONFIG_DIRS")
	if configDirsStr != "" {
		configDirs = strings.Split(configDirsStr, separator)
	}
	if len(configDirs) == 0 {
		configDirs = defaulter.defaultConfigDirs()
	}
	return configDirs
}

// CacheHome returns the location that should be used for application cache files
func CacheHome() string {
	cacheHome := os.Getenv("XDG_CACHE_HOME")
	if cacheHome == "" {
		cacheHome = defaulter.defaultCacheHome()
	}
	return cacheHome
}
