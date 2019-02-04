// Copyright 2016 Palantir Technologies, Inc.
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

package nobadfuncs

import (
	"encoding/json"
	"io"

	"github.com/palantir/okgo/checker"
	"github.com/palantir/okgo/okgo"
)

const (
	TypeName okgo.CheckerType     = "nobadfuncs"
	Priority okgo.CheckerPriority = 0
)

type Checker struct {
	BadFuncs map[string]string
}

func (c *Checker) Type() (okgo.CheckerType, error) {
	return TypeName, nil
}

func (c *Checker) Priority() (okgo.CheckerPriority, error) {
	return Priority, nil
}

func (c *Checker) Check(pkgPaths []string, pkgDir string, stdout io.Writer) {
	cfgJSON, err := json.Marshal(c.BadFuncs)
	if err != nil {
		okgo.WriteErrorAsIssue(err, stdout)
		return
	}

	cmd, wd := checker.AmalgomatedCheckCmd(string(TypeName), append([]string{
		"--config-json",
		string(cfgJSON),
	}, pkgPaths...), stdout)
	if cmd == nil {
		return
	}
	checker.RunCommandAndStreamOutput(cmd, func(line string) okgo.Issue {
		return okgo.NewIssueFromLine(line, wd)
	}, stdout)
}

func (c *Checker) RunCheckCmd(args []string, stdout io.Writer) {
	checker.AmalgomatedRunRawCheck(string(TypeName), args, stdout)
}
