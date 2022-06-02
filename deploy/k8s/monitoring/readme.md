# Prometheus monitoring stack

````shell
kubectl apply --server-side -f manifests/setup
````

````shell
until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
kubectl apply -f manifests/
````