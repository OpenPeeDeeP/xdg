// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdg

import (
	"os"
	"path/filepath"
)

func defaultDataHome() string {
	return filepath.Join(os.Getenv("HOME"), ".local", "share")
}

func defaultDataDirs() []string {
	return []string{"/usr/local/share/", "/usr/share/"}
}

func defaultConfigHome() string {
	return filepath.Join(os.Getenv("HOME"), ".config")
}

func defaultConfigDirs() []string {
	return []string{"/etc/xdg"}
}

func defaultCacheHome() string {
	return filepath.Join(os.Getenv("HOME"), ".cache")
}
