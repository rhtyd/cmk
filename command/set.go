// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package command

import (
	"fmt"
	"strings"
)

// FIXME: set apikey and secretkeys?
func init() {
	AddCommand(&Command{
		Name:        "set",
		Help:        "Configures options for cmk",
		SubCommands: []string{"block", "display", "profile"},
		Handle: func(r *Request) error {
			if len(r.Args) < 1 {
				fmt.Println("Please select one of the sub-commands: ", strings.Join(r.Command.SubCommands, ", "))
				return nil
			}
			subCommand := r.Args[0]
			value := strings.Join(r.Args[1:], " ")
			r.Config.UpdateGlobalConfig(subCommand, value)

			if subCommand == "profile" && r.Shell != nil {
				r.Shell.SetPrompt(r.Config.GetPrompt())
			}
			return nil
		},
	})
}
