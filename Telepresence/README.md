# Demo telepresence

## Requirements

- golang 1.12 or later
- [k8s](https://kubernetes.io/docs/tasks/tools/install-kubectl) 1.14 or later
  - Using [kustomize](https://kustomize.io)
  - Deploy to Amazon EKS
- [Telepresence](https://www.telepresence.io/reference/install) 0.101 or later

## Setting

### Local debugging

```shell
go run boot.go
```

### Building a container on local

```shell
docker build -t demo-telepresence .
docker run -it --rm -p 8080:8080 --name demo-telepresence demo-telepresence:latest /bin/sh
```

### Deploy to Amazon EKS

- Upload docker image to Amazon ECR

```shell
$(aws ecr get-login --no-include-email --region ap-northeast-1)
docker build -t demo-telepresence .
docker tag demo-telepresence:latest AWS_ECR_REPO/demo-telepresence:latest
docker push AWS_ECR_REPO/demo-telepresence:latest
```

- Deploy k8s components with kustomization

```shell
kubectl apply -k ./kustomize/
kubectl get -k ./kustomize/
```

## Usage

```shell
$ kubectl get svc
NAME         TYPE           CLUSTER-IP      EXTERNAL-IP                                                                    PORT(S)        AGE
demo         LoadBalancer   10.100.231.82   a679277f1b10b11e9939d06a3f13f8b1-1783984018.ap-northeast-1.elb.amazonaws.com   80:32493/TCP   3m34s
kubernetes   ClusterIP      10.100.0.1      <none>                                                                         443/TCP        50d

$ curl a679277f1b10b11e9939d06a3f13f8b1-1783984018.ap-northeast-1.elb.amazonaws.com:80/hc
Alive.
```

- Change the boot.go on local

```diff
- fmt.Fprintf(res, "Alive.")
+ fmt.Fprintf(res, "Alive with Telepresence.")
```

- Rebuilding source

```shell
$ go build -o boot .
```

- Start a proxy with Telepresence

```shell
$ sudo telepresence --swap-deployment demo --expose 8080 --run ./boot &
T: Warning: kubectl 1.15.0 may not work correctly with cluster version
T: 1.12.6-eks-d69f1b due to the version discrepancy. See
T: https://kubernetes.io/docs/setup/version-skew-policy/ for more information.

T: Starting proxy with method 'vpn-tcp', which has the following limitations: All
T: processes are affected, only one telepresence can run per machine, and you
T: can't use other VPNs. You may need to add cloud hosts and headless services
T: with --also-proxy. For a full list of method limitations see
T: https://telepresence.io/reference/methods.html
T: Volumes are rooted at $TELEPRESENCE_ROOT. See
T: https://telepresence.io/howto/volumes.html for details.
T: Starting network proxy to cluster by swapping out Deployment demo with a proxy
T: Forwarding remote port 8080 to local port 8080.

T: Setup complete. Launching your command.

$ kubectl get pods
NAME                                                     READY   STATUS        RESTARTS   AGE
demo-0e48ad3a8c4843689be31ca9d4fc423e-6b4f96b9f5-dw7jk   1/1     Running       0          8s
demo-8c8b598dc-btv9f                                     0/1     Terminating   0          18m
```

- Confirm the changed response content

```shell
$ curl a679277f1b10b11e9939d06a3f13f8b1-1783984018.ap-northeast-1.elb.amazonaws.com:80/hc
Alive with Telepresence.
```

- Stop a proxy with Telepresence

```shell
$ sudo fq
```

- Remove k8s components with kustomization

```shell
$ kubectl delete -k ./kustomize/
```
