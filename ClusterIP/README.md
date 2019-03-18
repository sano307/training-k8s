# Usage

```shell
kubectl apply -f clusterIP.yaml
kubectl proxy --port=8080
curl http://localhost:8080/api/v1/proxy/namespaces/default/services/nginx:80
```
