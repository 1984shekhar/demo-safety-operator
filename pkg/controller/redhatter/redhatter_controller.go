package redhatter

import (
	"context"
	apiv1alpha1 "github.com/akoserwal/demo-safety-operator.git/pkg/apis/api/v1alpha1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_redhatter")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new RedHatter Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileRedHatter{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("redhatter-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource RedHatter
	err = c.Watch(&source.Kind{Type: &apiv1alpha1.RedHatter{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &v1.ConfigMap{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileRedHatter implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileRedHatter{}

// ReconcileRedHatter reconciles a RedHatter object
type ReconcileRedHatter struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a RedHatter object and makes changes based on the state read
// and what is in the RedHatter.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileRedHatter) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling RedHatter")

	// Fetch the RedHatter instance
	instance := &apiv1alpha1.RedHatter{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			if instance.Spec.EmployeeName == "" {
				instance.Spec.EmployeeName = "Abhishek"
			}

			instance.Spec.EmployeeStatus = "Quarantine"
			r.configMapReconcile(instance, request)

			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.

		return reconcile.Result{}, err
	}


	if instance.Spec.EmployeeName == "" {
		instance.Spec.EmployeeName = "Abhishek"
	}

	//Setting instance
	instance.Spec.IsCovidThere = true

	if instance.Spec.IsCovidThere == false {
		instance.Spec.EmployeeStatus = "Party"
	} else {
		instance.Spec.EmployeeStatus = "Quarantine"
	}

	log.Info(instance.Kind)


	//configmap
	r.configMapReconcile(instance, request)

	//update instance
	err = r.client.Update(context.TODO(), instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	configMap := r.configMapDef(instance, request)

	err = r.client.Update(context.TODO(), configMap)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileRedHatter) configMapDef(instance *apiv1alpha1.RedHatter, request reconcile.Request) *v1.ConfigMap {
	data := map[string]string{
		"name":   instance.Spec.EmployeeName,
		"status": instance.Spec.EmployeeStatus,
	}

	configMap := &v1.ConfigMap{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Data:       data,
		BinaryData: nil,
	}
	configMap.Name = instance.Name
	configMap.Namespace = instance.Namespace
	return configMap
}

func (r *ReconcileRedHatter) configMapReconcile(instance *apiv1alpha1.RedHatter, request reconcile.Request) (reconcile.Result, error) {
	configMap := r.configMapDef(instance, request)

	key := client.ObjectKey{
		Namespace: instance.Namespace,
		Name:     instance.Name ,
	}

	err := r.client.Get(context.TODO(), key, configMap)
	if err != nil {
		if errors.IsNotFound(err) {

			err = r.client.Create(context.TODO(), configMap)
			if err != nil {
				if errors.IsNotFound(err) {
					return reconcile.Result{}, nil
				}
				// Error reading the object - requeue the request.
				return reconcile.Result{}, err
			}

			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	} else {
		err = r.client.Update(context.TODO(), configMap)
		if err != nil {
			if errors.IsNotFound(err) {
				return reconcile.Result{}, nil
			}
			// Error reading the object - requeue the request.
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}
