// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdg

import (
	"os"
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
	{"DataHome Without", "defaultDataHome", "/some/path", true, "XDG_DATA_HOME", "", MDataHome, nil, "/some/path"},
	{"DataDirs Without", "defaultDataDirs", []string{"/some/path", "/some/other/path"}, true, "XDG_DATA_DIRS", "", MDataDirs, nil, []string{"/some/path", "/some/other/path"}},
	{"ConfigHome Without", "defaultConfigHome", "/some/path", true, "XDG_CONFIG_HOME", "", MConfigHome, nil, "/some/path"},
	{"ConfigDirs Without", "defaultConfigDirs", []string{"/some/path", "/some/other/path"}, true, "XDG_CONFIG_DIRS", "", MConfigDirs, nil, []string{"/some/path", "/some/other/path"}},
	{"CacheHome Without", "defaultCacheHome", "/some/path", true, "XDG_CACHE_HOME", "", MCacheHome, nil, "/some/path"},

	{"DataHome With", "defaultDataHome", "/wrong/path", false, "XDG_DATA_HOME", "/some/path", MDataHome, nil, "/some/path"},
	{"DataDirs With", "defaultDataDirs", []string{"/wrong/path", "/some/other/wrong"}, false, "XDG_DATA_DIRS", strings.Join([]string{"/some/path", "/some/other/path"}, string(os.PathListSeparator)), MDataDirs, nil, []string{"/some/path", "/some/other/path"}},
	{"ConfigHome With", "defaultConfigHome", "/wrong/path", false, "XDG_CONFIG_HOME", "/some/path", MConfigHome, nil, "/some/path"},
	{"ConfigDirs With", "defaultConfigDirs", []string{"/wrong/path", "/some/other/wrong"}, false, "XDG_CONFIG_DIRS", strings.Join([]string{"/some/path", "/some/other/path"}, string(os.PathListSeparator)), MConfigDirs, nil, []string{"/some/path", "/some/other/path"}},
	{"CacheHome With", "defaultCacheHome", "/wrong/path", false, "XDG_CACHE_HOME", "/some/path", MCacheHome, nil, "/some/path"},

	{"DataHome App Without", "defaultDataHome", "/some/path", true, "XDG_DATA_HOME", "", MDataHome, New("OpenPeeDeeP", "XDG"), "/some/path/OpenPeeDeeP/XDG"},
	{"DataDirs App Without", "defaultDataDirs", []string{"/some/path", "/some/other/path"}, true, "XDG_DATA_DIRS", "", MDataDirs, New("OpenPeeDeeP", "XDG"), []string{"/some/path/OpenPeeDeeP/XDG", "/some/other/path/OpenPeeDeeP/XDG"}},
	{"ConfigHome App Without", "defaultConfigHome", "/some/path", true, "XDG_CONFIG_HOME", "", MConfigHome, New("OpenPeeDeeP", "XDG"), "/some/path/OpenPeeDeeP/XDG"},
	{"ConfigDirs App Without", "defaultConfigDirs", []string{"/some/path", "/some/other/path"}, true, "XDG_CONFIG_DIRS", "", MConfigDirs, New("OpenPeeDeeP", "XDG"), []string{"/some/path/OpenPeeDeeP/XDG", "/some/other/path/OpenPeeDeeP/XDG"}},
	{"CacheHome App Without", "defaultCacheHome", "/some/path", true, "XDG_CACHE_HOME", "", MCacheHome, New("OpenPeeDeeP", "XDG"), "/some/path/OpenPeeDeeP/XDG"},

	{"DataHome App With", "defaultDataHome", "/wrong/path", false, "XDG_DATA_HOME", "/some/path", MDataHome, New("OpenPeeDeeP", "XDG"), "/some/path/OpenPeeDeeP/XDG"},
	{"DataDirs App With", "defaultDataDirs", []string{"/wrong/path", "/some/other/wrong"}, false, "XDG_DATA_DIRS", strings.Join([]string{"/some/path", "/some/other/path"}, string(os.PathListSeparator)), MDataDirs, New("OpenPeeDeeP", "XDG"), []string{"/some/path/OpenPeeDeeP/XDG", "/some/other/path/OpenPeeDeeP/XDG"}},
	{"ConfigHome App With", "defaultConfigHome", "/wrong/path", false, "XDG_CONFIG_HOME", "/some/path", MConfigHome, New("OpenPeeDeeP", "XDG"), "/some/path/OpenPeeDeeP/XDG"},
	{"ConfigDirs App With", "defaultConfigDirs", []string{"/wrong/path", "/some/other/wrong"}, false, "XDG_CONFIG_DIRS", strings.Join([]string{"/some/path", "/some/other/path"}, string(os.PathListSeparator)), MConfigDirs, New("OpenPeeDeeP", "XDG"), []string{"/some/path/OpenPeeDeeP/XDG", "/some/other/path/OpenPeeDeeP/XDG"}},
	{"CacheHome App With", "defaultCacheHome", "/wrong/path", false, "XDG_CACHE_HOME", "/some/path", MCacheHome, New("OpenPeeDeeP", "XDG"), "/some/path/OpenPeeDeeP/XDG"},
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
