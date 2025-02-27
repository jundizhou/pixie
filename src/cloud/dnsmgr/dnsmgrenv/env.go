/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package dnsmgrenv

import (
	"github.com/spf13/viper"

	"px.dev/pixie/src/shared/services/env"
)

// DNSMgrEnv is the environment used for the dnsmgr service.
type DNSMgrEnv interface {
	env.Env
}

// Impl is an implementation of the DNSMgrEnv interface
type Impl struct {
	*env.BaseEnv
}

// New creates a new dnsmgr env.
func New() *Impl {
	return &Impl{env.New(viper.GetString("domain_name"))}
}
