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
	{"DataHome Without", "defaultDataHome", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), true, "XDG_DATA_HOME", "", MDataHome, nil, strings.Replace("/some/path", "/", string(os.PathSeparator), -1)},
	{"DataDirs Without", "defaultDataDirs", []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, true, "XDG_DATA_DIRS", "", MDataDirs, nil, []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}},
	{"ConfigHome Without", "defaultConfigHome", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), true, "XDG_CONFIG_HOME", "", MConfigHome, nil, strings.Replace("/some/path", "/", string(os.PathSeparator), -1)},
	{"ConfigDirs Without", "defaultConfigDirs", []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, true, "XDG_CONFIG_DIRS", "", MConfigDirs, nil, []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}},
	{"CacheHome Without", "defaultCacheHome", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), true, "XDG_CACHE_HOME", "", MCacheHome, nil, strings.Replace("/some/path", "/", string(os.PathSeparator), -1)},

	{"DataHome With", "defaultDataHome", strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), false, "XDG_DATA_HOME", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), MDataHome, nil, strings.Replace("/some/path", "/", string(os.PathSeparator), -1)},
	{"DataDirs With", "defaultDataDirs", []string{strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/wrong", "/", string(os.PathSeparator), -1)}, false, "XDG_DATA_DIRS", strings.Join([]string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, string(os.PathListSeparator)), MDataDirs, nil, []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}},
	{"ConfigHome With", "defaultConfigHome", strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), false, "XDG_CONFIG_HOME", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), MConfigHome, nil, strings.Replace("/some/path", "/", string(os.PathSeparator), -1)},
	{"ConfigDirs With", "defaultConfigDirs", []string{strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/wrong", "/", string(os.PathSeparator), -1)}, false, "XDG_CONFIG_DIRS", strings.Join([]string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, string(os.PathListSeparator)), MConfigDirs, nil, []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}},
	{"CacheHome With", "defaultCacheHome", strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), false, "XDG_CACHE_HOME", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), MCacheHome, nil, strings.Replace("/some/path", "/", string(os.PathSeparator), -1)},

	{"DataHome App Without", "defaultDataHome", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), true, "XDG_DATA_HOME", "", MDataHome, New("OpenPeeDeeP", "XDG"), strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)},
	{"DataDirs App Without", "defaultDataDirs", []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, true, "XDG_DATA_DIRS", "", MDataDirs, New("OpenPeeDeeP", "XDG"), []string{strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)}},
	{"ConfigHome App Without", "defaultConfigHome", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), true, "XDG_CONFIG_HOME", "", MConfigHome, New("OpenPeeDeeP", "XDG"), strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)},
	{"ConfigDirs App Without", "defaultConfigDirs", []string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, true, "XDG_CONFIG_DIRS", "", MConfigDirs, New("OpenPeeDeeP", "XDG"), []string{strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)}},
	{"CacheHome App Without", "defaultCacheHome", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), true, "XDG_CACHE_HOME", "", MCacheHome, New("OpenPeeDeeP", "XDG"), strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)},

	{"DataHome App With", "defaultDataHome", strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), false, "XDG_DATA_HOME", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), MDataHome, New("OpenPeeDeeP", "XDG"), strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)},
	{"DataDirs App With", "defaultDataDirs", []string{strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/wrong", "/", string(os.PathSeparator), -1)}, false, "XDG_DATA_DIRS", strings.Join([]string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, string(os.PathListSeparator)), MDataDirs, New("OpenPeeDeeP", "XDG"), []string{strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)}},
	{"ConfigHome App With", "defaultConfigHome", strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), false, "XDG_CONFIG_HOME", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), MConfigHome, New("OpenPeeDeeP", "XDG"), strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)},
	{"ConfigDirs App With", "defaultConfigDirs", []string{strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/wrong", "/", string(os.PathSeparator), -1)}, false, "XDG_CONFIG_DIRS", strings.Join([]string{strings.Replace("/some/path", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path", "/", string(os.PathSeparator), -1)}, string(os.PathListSeparator)), MConfigDirs, New("OpenPeeDeeP", "XDG"), []string{strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1), strings.Replace("/some/other/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)}},
	{"CacheHome App With", "defaultCacheHome", strings.Replace("/wrong/path", "/", string(os.PathSeparator), -1), false, "XDG_CACHE_HOME", strings.Replace("/some/path", "/", string(os.PathSeparator), -1), MCacheHome, New("OpenPeeDeeP", "XDG"), strings.Replace("/some/path/OpenPeeDeeP/XDG", "/", string(os.PathSeparator), -1)},
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
