# Namespaces

- Logical groups of resources
- Set namespace
  `k config set-context --curent --namespace=default`
- Can use to close off network access using network policies
- Create n all other see [[Deployments]]
  `k create namespace lp`
- Can test bunch of things within namespace and delete it
  all together using the namespace
- get pods only for namespace, similarily specify namespace
  for any of the commands
  `k get pods --namespace lp`
  `k get pods -n lp`
- In deployment.yaml can specify namespace under meta
