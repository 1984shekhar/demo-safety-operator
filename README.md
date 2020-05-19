# demo-safety-operator
Demo operator

```
export GO111MODULE=on 
operator-sdk new demo-safety-operator --repo github.com/akoserwal/demo-safety-operator.git

// adding API
operator-sdk add api --api-version=api.redhat.com/v1alpha1 --kind=RedHatter

// adding Controller
operator-sdk add controller --api-version=api.redhat.com/v1alpha1 --kind=RedHatter

// after making changes to api code
operator-sdk generate k8s
operator-sdk generate crds

// deploy
oc apply -f deploy/role.yaml 
oc apply -f deploy/role_binding.yaml
oc apply -f deploy/service_account.yaml 
oc apply -f deploy/crds/api.redhat.com_redhatters_crd.yaml

// Creating the instance (CR):
 oc apply -f deploy/crds/api.redhat.com_v1alpha1_redhatter_cr.yaml

// running operator locally
operator-sdk run --local 
```
