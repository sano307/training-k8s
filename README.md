# Training k8s

## Requirements

- Installed Kubernetes 1.9 or later
- Installed kubectl
- Have a kubeconifg file (default location is `~/.kube/config`)
- Installed Helm

### Secret information

- Docker Hub

```shell
kubectl create secret docker-registry docker-hub-credential \
  --docker-server=https://index.docker.io/v1  \
  --docker-username=${DOCKER_HUB_USERNAME} \
  --docker-password=${DOCKER_HUB_PASSWORD} \
  --docker-email=${DOCKER_HUB_EMAIL}
```

- MySQL

```shell
kubectl create secret generic dev-mysql-read \
  --from-literal=host=${DEV_MYSQL_HOST} \
  --from-literal=db=${DEV_MYSQL_DATABASE} \
  --from-literal=username=${DEV_MYSQL_USERNAME} \
  --from-literal=password=${DEV_MYSQL_PASSWORD}
```

## Usage

### How to upload docker image to the docker-hub

```shell
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${REMOTE_IMAGE_REPOSITORY}:${REMOTE_IMAGE_TAG}
docker push ${REMOTE_IMAGE_REPOSITORY}:${REMOTE_IMAGE_TAG}
```

### How to apply helm charts

```shell
helm template charts -f charts/${ENV}.yaml | tee manifest.yaml
kubectl apply -f manifest.yaml
```
