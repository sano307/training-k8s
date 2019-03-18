# Usage

```shell
kubectl apply -f loadBalancer.yaml
kubectl get svc nginx
NAME    TYPE           CLUSTER-IP      EXTERNAL-IP      PORT(S)                                 AGE
nginx   LoadBalancer   ${CLUSTER_IP}   ${EXTERNAL_IP}   ${EXTERNAL_PORT}:${INTERNAL_PORT}/TCP   12s

curl ${EXTERNAL_IP}:${EXTERNAL_PORT}
```
