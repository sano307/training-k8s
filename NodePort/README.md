# Usage

```shell
kubectl apply -f nodePort.yaml
kubectl describe node | grep InternalIP
  InternalIP:  ${INTERNAL_IP}

curl ${INTERNAL_IP}:30036
```
