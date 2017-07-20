// Copyright (c) 2017, OpenPeeDeeP. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xdg

import "os"

func defaultDataHome() string {
	return os.Getenv("APPDATA")
}

func defaultDataDirs() []string {
	return []string{os.Getenv("PROGRAMDATA")}
}

func defaultConfigHome() string {
	return os.Getenv("APPDATA")
}

func defaultConfigDirs() []string {
	return []string{os.Getenv("PROGRAMDATA")}
}

func defaultCacheHome() string {
	return os.Getenv("LOCALAPPDATA")
}
