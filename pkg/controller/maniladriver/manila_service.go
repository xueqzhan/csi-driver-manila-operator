package maniladriver

import (
	"context"

	"github.com/banzaicloud/k8s-objectmatcher/patch"
	"github.com/go-logr/logr"
	maniladriverv1alpha1 "github.com/openshift/csi-driver-manila-operator/pkg/apis/maniladriver/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcileManilaDriver) handleManilaControllerPluginService(instance *maniladriverv1alpha1.ManilaDriver, reqLogger logr.Logger) error {
	reqLogger.Info("Reconciling Manila Controller Plugin Service")

	var labels = map[string]string{
		"app":       "openstack-manila-csi",
		"component": "controllerplugin",
	}

	// Define a new Service object
	srv := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "openstack-manila-csi-controllerplugin",
			Namespace: "manila-csi",
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "dummy",
					Port: int32(12345),
				},
			},
			Selector: labels,
		},
	}

	// Set ManilaDriver instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, srv, r.scheme); err != nil {
		return err
	}

	// Check if this Service already exists
	found := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: srv.Name, Namespace: srv.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", srv.Namespace, "Service.Name", srv.Name)
		err = r.client.Create(context.TODO(), srv)
		if err != nil {
			return err
		}

		// Service created successfully - don't requeue
		return nil
	} else if err != nil {
		return err
	}

	// Check if we need to update the object
	patchResult, err := patch.DefaultPatchMaker.Calculate(found, srv)
	if err != nil {
		return err
	}

	if !patchResult.IsEmpty() {
		reqLogger.Info("Updating Service with new changes", "Service.Namespace", found.Namespace, "Service.Name", found.Name)
		err = r.client.Update(context.TODO(), srv)
		if err != nil {
			return err
		}
	} else {
		// Service already exists - don't requeue
		reqLogger.Info("Skip reconcile: Service already exists", "Service.Namespace", found.Namespace, "Service.Name", found.Name)
	}

	return nil
}