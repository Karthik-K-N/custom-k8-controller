# Writing Custom Kubernetes Controller and Webhooks

### More details about this repository can be found in [Medium blog](https://medium.com/@Karthik-K-N/writing-custom-kubernetes-controller-and-webhooks-141230820e9)

## Create kind cluster and install certificate manager for testing
```text
$ kind create cluster --name custom-controller
$ kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.1/cert-manager.yaml
```

## Create and install the manifests
```text
$ make manifests
$ make install
```

## Build the controller
```text
$ make docker-build IMG=karthik/custom-controller:v1
```

## Load the image on to kind cluster
```text
$ kind load docker-image karthik/custom-controller:v1 --name custom-controller
```

## Run the controller
```text
$ make deploy IMG=karthik/custom-controller:v1
```

## Verify the pod state
```text
$ kubectl -n custom-k8-controller-system  get pods
```

## Get the logs of the pod
```text
$ kubectl -n custom-k8-controller-system logs -f <pod_name>
```

## Create a sample sum resource
```text
$ kubectl create -f config/samples/calculator_v1_sum.yaml
```
