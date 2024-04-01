/*
Copyright 2023 KDP(Kubernetes Data Platform).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	bdcctrl "kdp-oam-operator/pkg/controllers/bdc"
	"kdp-oam-operator/pkg/controllers/bdc/v1alpha1/application"
	"kdp-oam-operator/pkg/controllers/bdc/v1alpha1/bigdatacluster"
	"kdp-oam-operator/pkg/controllers/bdc/v1alpha1/contextsecret"
	"kdp-oam-operator/pkg/controllers/bdc/v1alpha1/contextsetting"
	"kdp-oam-operator/pkg/controllers/bdc/v1alpha1/xdefinitions"

	ctrl "sigs.k8s.io/controller-runtime"
)

func Setup(mgr ctrl.Manager, args bdcctrl.Args) error {
	for _, setup := range []func(ctrl.Manager, bdcctrl.Args) error{
		bigdatacluster.Setup,
		contextsecret.Setup,
		contextsetting.Setup,
		xdefinitions.Setup,
		application.Setup,
	} {
		if err := setup(mgr, args); err != nil {
			return err
		}
	}
	return nil
}
