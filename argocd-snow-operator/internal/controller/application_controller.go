/*
Copyright 2025.

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

package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ApplicationReconciler reconciles an Application object
type ApplicationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=argoproj.io,resources=applications,verbs=get;list;watch

func (r *ApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the Application resource
	app := &unstructured.Unstructured{}
	app.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "argoproj.io",
		Version: "v1alpha1",
		Kind:    "Application",
	})

	err := r.Get(ctx, req.NamespacedName, app)
	if err != nil {
		logger.Error(err, "unable to fetch Application")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Application", "appObject", app.Object)

	// Extract sync status
	syncStatus, found, err := unstructured.NestedString(app.Object, "status", "operationState", "phase")
	if err != nil || !found {
		logger.Info("Sync status not found in Application")
		return ctrl.Result{}, nil
	}

	logger.Info("Application", "app", app)

	//logger.Info("Application sync status", "status", syncStatus)
	//annotationValue, err := getAnnotationValue(app, "foo.bar.com/snow-data")

	if err != nil {
		logger.Error(err, "Error fetching annotation")
		return ctrl.Result{}, nil
	}

	//logger.Info("Annotation value", "value", annotationValue)

	// Handle different sync phases
	switch syncStatus {
	case "Running":
		logger.Info("Sync Phase: Is Running")
	case "Syncing":
		logger.Info("Sync Phase: Is Syncing")
		//startChange(ctx, app)
		// Add logic to handle sync start
	case "Succeeded":
		logger.Info("Application sync succeeded")
		//startChange(ctx, app)
		//endChange(ctx, app)
		//closeChange(ctx, app)
		//addEvidence(ctx, app)
		// Add logic to handle sync success
	case "Failed":
		logger.Info("Application sync failed")
		// Add logic to handle sync failure
	default:
		logger.Info("Unhandled sync status", "status", syncStatus)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "argoproj.io/v1alpha1",
				"kind":       "Application",
			},
		}).
		Complete(r)
}

func startChange(ctx context.Context, app *unstructured.Unstructured) {
	logger := log.FromContext(ctx)
	logger.Info("Create change", "name", app.GetName())
	logger.Info("Start change")
}

func endChange(ctx context.Context, app *unstructured.Unstructured) {
	logger := log.FromContext(ctx)
	logger.Info("End change", "name", app.GetName())
}

func closeChange(ctx context.Context, app *unstructured.Unstructured) {
	logger := log.FromContext(ctx)
	logger.Info("Close change", "name", app.GetName())
}

func addEvidence(ctx context.Context, app *unstructured.Unstructured) {
	logger := log.FromContext(ctx)
	logger.Info("Add evidence", "name", app.GetName())
}
