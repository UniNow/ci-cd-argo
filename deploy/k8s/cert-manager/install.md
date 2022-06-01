## Cert Manager installation

```shell
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.yaml
```

Since the demo is running on a scaleway cluster, we also have to install the hairpin proxy to get the cert exchange working
```shell
kubectl apply -f https://raw.githubusercontent.com/compumike/hairpin-proxy/v0.2.1/deploy.yml
```