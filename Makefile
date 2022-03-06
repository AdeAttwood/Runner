
# Copyright 2022 Practically.io All rights reserved
#
# Use of this source is governed by a BSD-style
# licence that can be found in the LICENCE file or at
# https://www.practically.io/copyright

build: ## Builds the executable
	go build  -o ./bin/run -i github.com/AdeAttwood/Runner/cmd/run
