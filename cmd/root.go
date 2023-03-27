// Copyright 2019 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
  "github.com/spf13/cobra"
)

var Root = New("ocip", "start registry service")

func New(use, short string) *cobra.Command {

  root := &cobra.Command{
    Use:   use,
    Short: short,
    RunE:  func(cmd *cobra.Command, _ []string) error { return cmd.Usage() },
  }

  root.AddCommand(
    newCmdRegistry(),
  )
  return root
}
