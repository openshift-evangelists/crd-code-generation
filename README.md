# crd-code-generation

Example repository for the blog post [Kubernetes Deep Dive: Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/).

## Getting Started

First register the custom resource definition:

```
kubectl apply -f artifacts/databases-crd.yaml
```

Then add an example of the `Database` kind:

```
kubectl apply -f artifacts/my-database.yaml
```

Finally build and run the example:

```
cd cmd/example
go build
./example -kubeconfig ~/.kube/config
```
