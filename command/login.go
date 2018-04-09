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
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)


func init() {
	AddCommand(&Command{
		Name: "login",
		Help: "Log in to your account",
		Handle: func(r *Request) error {
			if len(r.Args) > 0 {
				return errors.New("this does not accept any additional arguments")
			}

			validate := func(input string) error {
				if len(input) < 1 {
					return errors.New("Please enter something")
				}
				return nil
			}

			prompt := promptui.Prompt{
				Label:    "Username",
				Validate: validate,
				Default:  "",
			}

			username, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			prompt = promptui.Prompt{
				Label:    "Password",
				Validate: validate,
				Mask:     '*',
			}

			password, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return nil
			}

			fmt.Println("Trying to log in using", username, password)

			return nil
		},
	})
}