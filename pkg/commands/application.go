// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
package commands

type Application struct {
	Cwd       string
	Help      bool
	Arguments []string
}
