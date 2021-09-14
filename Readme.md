````
minikube image build -t "example.com/go-webhook:latest" . -f Containerfile
````

The main structure and the test setup is all stolen proudly from what the
[Kubebuilder](https://kubebuilder.io) kickstarted.

# Debugging

* You can enable `attachControlPlaneOutput` in [kuttl-test.yaml](./test/kuttl-test.yaml)
to run the tests with `kube-apiserver` logs (it's quite verbose).