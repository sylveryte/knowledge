# Cluster

- Group of worker nodes are called **cluster** also **context**
- Current cluster
  `k config current-context`
- Switch context
  `k config use-context context_name_or_address`
- [[Pods]] within the cluster can communicate with each
  other via ips
  > [!info] get Ips of pods
  > k get pods -o wide
