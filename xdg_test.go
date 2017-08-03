// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDefaulter struct {
	mock.Mock
}

func (m *mockDefaulter) defaultDataHome() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockDefaulter) defaultDataDirs() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
func (m *mockDefaulter) defaultConfigHome() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockDefaulter) defaultConfigDirs() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
func (m *mockDefaulter) defaultCacheHome() string {
	args := m.Called()
	return args.String(0)
}

const (
	MDataHome = iota
	MDataDirs
	MConfigHome
	MConfigDirs
	MCacheHome
)

var getterTestCases = []getterTestCase{
	{"DataHome Without", "defaultDataHome", filepath.Clean("/some/path"), true, "XDG_DATA_HOME", "", MDataHome, nil, filepath.Clean("/some/path")},
	{"DataDirs Without", "defaultDataDirs", []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, true, "XDG_DATA_DIRS", "", MDataDirs, nil, []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}},
	{"ConfigHome Without", "defaultConfigHome", filepath.Clean("/some/path"), true, "XDG_CONFIG_HOME", "", MConfigHome, nil, filepath.Clean("/some/path")},
	{"ConfigDirs Without", "defaultConfigDirs", []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, true, "XDG_CONFIG_DIRS", "", MConfigDirs, nil, []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}},
	{"CacheHome Without", "defaultCacheHome", filepath.Clean("/some/path"), true, "XDG_CACHE_HOME", "", MCacheHome, nil, filepath.Clean("/some/path")},

	{"DataHome With", "defaultDataHome", filepath.Clean("/wrong/path"), false, "XDG_DATA_HOME", filepath.Clean("/some/path"), MDataHome, nil, filepath.Clean("/some/path")},
	{"DataDirs With", "defaultDataDirs", []string{filepath.Clean("/wrong/path"), filepath.Clean("/some/other/wrong")}, false, "XDG_DATA_DIRS", strings.Join([]string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, string(os.PathListSeparator)), MDataDirs, nil, []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}},
	{"ConfigHome With", "defaultConfigHome", filepath.Clean("/wrong/path"), false, "XDG_CONFIG_HOME", filepath.Clean("/some/path"), MConfigHome, nil, filepath.Clean("/some/path")},
	{"ConfigDirs With", "defaultConfigDirs", []string{filepath.Clean("/wrong/path"), filepath.Clean("/some/other/wrong")}, false, "XDG_CONFIG_DIRS", strings.Join([]string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, string(os.PathListSeparator)), MConfigDirs, nil, []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}},
	{"CacheHome With", "defaultCacheHome", filepath.Clean("/wrong/path"), false, "XDG_CACHE_HOME", filepath.Clean("/some/path"), MCacheHome, nil, filepath.Clean("/some/path")},

	{"DataHome App Without", "defaultDataHome", filepath.Clean("/some/path"), true, "XDG_DATA_HOME", "", MDataHome, New("OpenPeeDeeP", "XDG"), filepath.Clean("/some/path/OpenPeeDeeP/XDG")},
	{"DataDirs App Without", "defaultDataDirs", []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, true, "XDG_DATA_DIRS", "", MDataDirs, New("OpenPeeDeeP", "XDG"), []string{filepath.Clean("/some/path/OpenPeeDeeP/XDG"), filepath.Clean("/some/other/path/OpenPeeDeeP/XDG")}},
	{"ConfigHome App Without", "defaultConfigHome", filepath.Clean("/some/path"), true, "XDG_CONFIG_HOME", "", MConfigHome, New("OpenPeeDeeP", "XDG"), filepath.Clean("/some/path/OpenPeeDeeP/XDG")},
	{"ConfigDirs App Without", "defaultConfigDirs", []string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, true, "XDG_CONFIG_DIRS", "", MConfigDirs, New("OpenPeeDeeP", "XDG"), []string{filepath.Clean("/some/path/OpenPeeDeeP/XDG"), filepath.Clean("/some/other/path/OpenPeeDeeP/XDG")}},
	{"CacheHome App Without", "defaultCacheHome", filepath.Clean("/some/path"), true, "XDG_CACHE_HOME", "", MCacheHome, New("OpenPeeDeeP", "XDG"), filepath.Clean("/some/path/OpenPeeDeeP/XDG")},

	{"DataHome App With", "defaultDataHome", filepath.Clean("/wrong/path"), false, "XDG_DATA_HOME", filepath.Clean("/some/path"), MDataHome, New("OpenPeeDeeP", "XDG"), filepath.Clean("/some/path/OpenPeeDeeP/XDG")},
	{"DataDirs App With", "defaultDataDirs", []string{filepath.Clean("/wrong/path"), filepath.Clean("/some/other/wrong")}, false, "XDG_DATA_DIRS", strings.Join([]string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, string(os.PathListSeparator)), MDataDirs, New("OpenPeeDeeP", "XDG"), []string{filepath.Clean("/some/path/OpenPeeDeeP/XDG"), filepath.Clean("/some/other/path/OpenPeeDeeP/XDG")}},
	{"ConfigHome App With", "defaultConfigHome", filepath.Clean("/wrong/path"), false, "XDG_CONFIG_HOME", filepath.Clean("/some/path"), MConfigHome, New("OpenPeeDeeP", "XDG"), filepath.Clean("/some/path/OpenPeeDeeP/XDG")},
	{"ConfigDirs App With", "defaultConfigDirs", []string{filepath.Clean("/wrong/path"), filepath.Clean("/some/other/wrong")}, false, "XDG_CONFIG_DIRS", strings.Join([]string{filepath.Clean("/some/path"), filepath.Clean("/some/other/path")}, string(os.PathListSeparator)), MConfigDirs, New("OpenPeeDeeP", "XDG"), []string{filepath.Clean("/some/path/OpenPeeDeeP/XDG"), filepath.Clean("/some/other/path/OpenPeeDeeP/XDG")}},
	{"CacheHome App With", "defaultCacheHome", filepath.Clean("/wrong/path"), false, "XDG_CACHE_HOME", filepath.Clean("/some/path"), MCacheHome, New("OpenPeeDeeP", "XDG"), filepath.Clean("/some/path/OpenPeeDeeP/XDG")},
}

type getterTestCase struct {
	name         string
	mokedMethod  string
	mockedReturn interface{}
	calledMocked bool
	env          string
	envVal       string
	method       int
	xdgApp       *XDG
	expected     interface{}
}

func TestXDG_Getters(t *testing.T) {
	for _, tc := range getterTestCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			mockDef := new(mockDefaulter)
			mockDef.On(tc.mokedMethod).Return(tc.mockedReturn)
			setDefaulter(mockDef)
			os.Setenv(tc.env, tc.envVal) // nolint: errcheck

			actual := computeActual(tc)

			if tc.calledMocked {
				mockDef.AssertExpectations(t)
			} else {
				mockDef.AssertNotCalled(t, tc.mokedMethod)
			}
			assert.Equal(tc.expected, actual)
		})
	}
}

// nolint: gocyclo
func computeActual(tc getterTestCase) interface{} {
	var actual interface{}
	switch tc.method {
	case MDataHome:
		if tc.xdgApp != nil {
			actual = tc.xdgApp.DataHome()
		} else {
			actual = DataHome()
		}
	case MDataDirs:
		if tc.xdgApp != nil {
			actual = tc.xdgApp.DataDirs()
		} else {
			actual = DataDirs()
		}
	case MConfigHome:
		if tc.xdgApp != nil {
			actual = tc.xdgApp.ConfigHome()
		} else {
			actual = ConfigHome()
		}
	case MConfigDirs:
		if tc.xdgApp != nil {
			actual = tc.xdgApp.ConfigDirs()
		} else {
			actual = ConfigDirs()
		}
	case MCacheHome:
		if tc.xdgApp != nil {
			actual = tc.xdgApp.CacheHome()
		} else {
			actual = CacheHome()
		}
	}
	return actual
}

const (
	QData = iota
	QConfig
	QCache
)

var (
	root      = "testingFolder"
	fileTypes = []string{"data", "config", "cache"}
	fileLoc   = []string{"home", "dirs"}
)

type queryTestCase struct {
	name      string
	xdgApp    *XDG
	queryType int
	filename  string
	expected  string
}

var queryTestCases = []queryTestCase{
	{"Data Dirs", New("OpenPeeDeeP", "XDG"), QData, "XDG_DATA_DIRS.txt", "/data/dirs/OpenPeeDeeP/XDG/XDG_DATA_DIRS.txt"},
	{"Data Home", New("OpenPeeDeeP", "XDG"), QData, "XDG_DATA_HOME.txt", "/data/home/OpenPeeDeeP/XDG/XDG_DATA_HOME.txt"},
	{"Data DNE", New("OpenPeeDeeP", "XDG"), QData, "XDG_CONFIG_HOME.txt", ""},

	{"Config Dirs", New("OpenPeeDeeP", "XDG"), QConfig, "XDG_CONFIG_DIRS.txt", "/config/dirs/OpenPeeDeeP/XDG/XDG_CONFIG_DIRS.txt"},
	{"Config Home", New("OpenPeeDeeP", "XDG"), QConfig, "XDG_CONFIG_HOME.txt", "/config/home/OpenPeeDeeP/XDG/XDG_CONFIG_HOME.txt"},
	{"Config DNE", New("OpenPeeDeeP", "XDG"), QConfig, "XDG_DATA_HOME.txt", ""},

	{"Config Home", New("OpenPeeDeeP", "XDG"), QCache, "XDG_CACHE_HOME.txt", "/cache/home/OpenPeeDeeP/XDG/XDG_CACHE_HOME.txt"},
	{"Config DNE", New("OpenPeeDeeP", "XDG"), QCache, "XDG_CACHE_DIRS.txt", ""},
}

func TestXDG_Query(t *testing.T) {
	for _, tc := range queryTestCases {
		t.Run(tc.name, func(t *testing.T) {
			defer teardownQueryData() //nolint: errcheck
			standupQueryData(tc)      //nolint: errcheck
			assert := assert.New(t)
			actual := computeQuery(tc)
			assert.Equal(tc.expected, actual)
		})
	}
}

func computeQuery(tc queryTestCase) string {
	var actual string
	switch tc.queryType {
	case QData:
		actual = tc.xdgApp.QueryData(tc.filename)
	case QCache:
		actual = tc.xdgApp.QueryCache(tc.filename)
	case QConfig:
		actual = tc.xdgApp.QueryConfig(tc.filename)
	}
	rootAbs, _ := filepath.Abs(root)
	actual = strings.Replace(actual, rootAbs, "", 1)
	return actual
}

func standupQueryData(tc queryTestCase) error {
	for _, t := range fileTypes {
		for _, l := range fileLoc {
			path, err := filepath.Abs(filepath.Join(root, t, l))
			if err != nil {
				return err
			}
			if err = os.MkdirAll(filepath.Join(path, tc.xdgApp.Vendor, tc.xdgApp.Application), 0777); err != nil {
				return err
			}
			envVar := fmt.Sprintf("XDG_%s_%s", strings.ToUpper(t), strings.ToUpper(l))
			if err = os.Setenv(envVar, path); err != nil {
				return err
			}
			file, err := os.OpenFile(filepath.Join(path, tc.xdgApp.Vendor, tc.xdgApp.Application, envVar+".txt"), os.O_CREATE|os.O_RDONLY, 0666)
			if err != nil {
				return err
			}
			defer file.Close() //nolint: errcheck
		}
	}
	return nil
}

func teardownQueryData() error {
	return os.RemoveAll(root)
}
