// Copyright 2023 The KusionStack Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math/rand"
	"time"

	"k8s.io/component-base/logs"

	"kusionstack.io/rollout/cmd/rollout/app"
	"kusionstack.io/rollout/cmd/rollout/app/options"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	opt := options.NewOptions()
	command := app.NewRolloutCommand(opt)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		panic(err.Error())
	}
}
