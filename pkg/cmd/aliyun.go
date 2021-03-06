// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

// NewAliyunCmd returns a new aliyun command.
func NewAliyunCmd(targetReader TargetReader) *cobra.Command {
	return &cobra.Command{
		Use:          "aliyun <args>",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			target := targetReader.ReadTarget(pathTarget)
			if len(target.Stack()) < 3 {
				return errors.New("no shoot targeted")
			}

			arguments := "aliyun " + strings.Join(args[:], " ")
			operate("aliyun", arguments)

			return nil
		},
	}
}
