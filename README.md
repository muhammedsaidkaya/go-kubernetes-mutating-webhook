
## K8s Admission Controller

* **Injects tcpdump container as a sidecar by object-label (tcpdump-enabled: "true") while creating PODs.**
* **Written in Golang.**

## Apply commands on Provision.sh

1. Docker Image Build (Go Server)
2. Create CA and Client Certificates (CloudFlare SSL)
3. Create secret file with client certificates (./manifests/webhook-secrets.yaml)
4. Create webhook configuration with ca bundle (./manifests/webhook-configuration.yaml)

5. Apply secret resources (./manifests/webhook-secrets.yaml)
6. Create serviceaccount (./manifests/webhook-rbac.yaml)
7. Create deployment and svc for webhook server (./manifests/webhook-deployment-svc.yaml)
8. Create mutatingwebhookconfiguration (./manifests/webhook-configuration.yaml)


* Change docker image name before STEP 1
  * Update container's image (./manifests/webhook-deployment-svc.yaml)

* Change hostname according to webhook name and namespace before STEP 2
  * Update serviceaccount's namespace (./manifests/webhook-rbac.yaml)
  * Update deplomyent's and svc's namespace and webhook name (./manifests/webhook-deployment-svc.yaml)
  * Update clientConfig's namespace and svc  (./manifests/webhook-configuration-template.yaml)
