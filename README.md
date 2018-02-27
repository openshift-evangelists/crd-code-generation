# crd-code-generation

Example repository for the blog post [Kubernetes Deep Dive: Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/).

## Installation

Checkout this repo under this folder and set the GOPATH
export GOPATH=~/go
go get github.com/openshift-evangelists/crd-code-generation

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
cd cd ~/go/src/github.com/openshift-evangelists/crd-code-generation/cmd/example/
go build
./example -kubeconfig ~/.kube/config
```
