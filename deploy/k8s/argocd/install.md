## ArgoCD installation

For more infos see https://argo-cd.readthedocs.io/en/stable/getting_started/

``` shell
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

