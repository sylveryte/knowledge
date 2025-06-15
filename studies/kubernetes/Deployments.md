# Kubernetes Deployments

- Create Deployments
  `k create deploy test_name --image=httpd --replicas=3`
- Create Deployments Yaml
  `k create deploy test_name --image=httpd --replicas=3 --dry-run=client -o yaml`
- list
  `k get deploy<Tab>`
- edit
  `k edit deploy<Tab>`
- describe
  `k describe deploy<Tab>`
- delete
  `k delete deploy<Tab>`
- apply
  `k apply -f deploy.yaml`

## Capabilities

- Can run 10 replicas of a pod
- Can slowly replace all 10s with new 10s without halting
  the operation
  - Simply `k apply -f updated_deploy.yaml`
- You can specify how much should be available and how many
  new can be created at a time,
  > [!info] docs
  > See Kubernetes docs for maxSurge and maxUnavailable

