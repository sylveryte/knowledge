# Pods

- Pod is _not_ a container, can be though
- Pod is **collection of containers** + other resources
  (networks or volumes etc)
- Pods run on worker nodes.
- Pods are ephemeral/expendable, don't expect long life
- Pods constently change and move around nodes

- Run a pod
  `k run nginx_thi_is_name --image=nginx`
- Get info of pod
  `k describe pod`
  `k describe pod nginx_thi_is_name`
  `k get pods -o wide`

## Pods As Code

- Create yaml out of pod
  `k get pod nginx_thi_is_name -o yaml`
  `k run nginx_yaml --image=nginx --dry-run=client -o yaml`
- Create pod from yaml
  `k create -f nginx.yaml` //create only
  `k apply -f nginx.yaml` //apply changes
- Edit pods yaml
  `k edit pod nginx_thi_is_name`

## Interact with pod

- Run command
  `k exec -it pod_name -- /bin/bash`

> [!info] k is
> kubectl
