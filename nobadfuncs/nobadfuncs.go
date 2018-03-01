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

	"github.com/go-yaml/yaml"
	"github.com/palantir/okgo/checker"
	"github.com/palantir/okgo/okgo"
	"github.com/pkg/errors"
)

const (
	TypeName okgo.CheckerType     = "nobadfuncs"
	Priority okgo.CheckerPriority = 0
)

func Creator() checker.Creator {
	return checker.NewCreator(
		TypeName,
		Priority,
		func(cfgYML []byte) (okgo.Checker, error) {
			var cfg nobadfuncsCheckCfg
			if err := yaml.Unmarshal(cfgYML, &cfg); err != nil {
				return nil, errors.Wrapf(err, "failed to unmarshal configuration YAML %q", string(cfgYML))
			}
			return &nobadfuncsCheck{
				BadFuncs: cfg.BadFuncs,
			}, nil
		},
	)
}

type nobadfuncsCheck struct {
	BadFuncs map[string]string
}

type nobadfuncsCheckCfg struct {
	BadFuncs map[string]string `yaml:"bad-funcs"`
}

func (c *nobadfuncsCheck) Type() (okgo.CheckerType, error) {
	return TypeName, nil
}

func (c *nobadfuncsCheck) Priority() (okgo.CheckerPriority, error) {
	return Priority, nil
}

func (c *nobadfuncsCheck) Check(pkgPaths []string, pkgDir string, stdout io.Writer) {
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

func (c *nobadfuncsCheck) RunCheckCmd(args []string, stdout io.Writer) {
	checker.AmalgomatedRunRawCheck(string(TypeName), args, stdout)
}
