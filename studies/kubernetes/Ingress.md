# Ingress

- Ingress is URLs
- Exposes HTTP & HTTPS routes from outside cluster to
  services within the cluster
- SSL/TSL Termination
- External URL provider
- Path based routing
  - Route to service based on path
    - example.com/meal -> mealie service
    - example.com/book -> hat service
- Ingress controllers
  - nginx
  - traefik
  - cilium
  - etc
- Flow
  - ==DNS -> LoadBalancer -> Ingress -> Service==

---
