// Copyright 2018 mixtool authors
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

package mixer

import (
	"fmt"

	"github.com/google/go-jsonnet"
)

func evaluatePrometheusAlerts(vm *jsonnet.VM, filename string) (string, error) {
	snippet := fmt.Sprintf("(import '%s').prometheusAlerts", filename)
	return vm.EvaluateSnippet("", snippet)
}

func evaluatePrometheusRules(vm *jsonnet.VM, filename string) (string, error) {
	snippet := fmt.Sprintf("(import '%s').prometheusRules", filename)
	return vm.EvaluateSnippet("", snippet)
}

func evaluateGrafanaDashboards(vm *jsonnet.VM, filename string) (string, error) {
	snippet := fmt.Sprintf("(import '%s').grafanaDashboards", filename)
	return vm.EvaluateSnippet("", snippet)
}