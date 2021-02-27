// Copyright 2021 Daniel Foehr
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

package history

import (
	"fmt"

	"github.com/danielfoehrkn/kubeswitch/pkg/store"
	"github.com/danielfoehrkn/kubeswitch/pkg/subcommands/history/util"
	setcontext "github.com/danielfoehrkn/kubeswitch/pkg/subcommands/set-context"
	"github.com/danielfoehrkn/kubeswitch/types"
	"github.com/ktr0731/go-fuzzyfinder"
)

func ListHistory(stores []store.KubeconfigStore, config *types.Config, stateDir string) error {
	history, err := util.ReadHistory()
	if err != nil {
		return err
	}

	idx, err := fuzzyfinder.Find(
		history,
		func(i int) string {
			return fmt.Sprintf("%d: %s", len(history)-i-1, history[i])
		})

	if err != nil {
		return err
	}

	return setcontext.SetContext(history[idx], stores, config, stateDir)
}

func SetPreviousContext(stores []store.KubeconfigStore, config *types.Config, stateDir string) error {
	history, err := util.ReadHistory()
	if err != nil {
		return err
	}

	if len(history) == 0 {
		return nil
	}

	var position int
	if len(history) == 1 {
		position = 0
	} else {
		position = 1
	}

	return setcontext.SetContext(history[position], stores, config, stateDir)
}
