#!/bin/bash

docker build . -t uzumlukek/tcpdump-webhook:v1

docker run -it --rm -v ${PWD}:/work -w /work debian bash

apt-get update && apt-get install -y curl &&
curl -L https://github.com/cloudflare/cfssl/releases/download/v1.5.0/cfssl_1.5.0_linux_amd64 -o /usr/local/bin/cfssl && \
curl -L https://github.com/cloudflare/cfssl/releases/download/v1.5.0/cfssljson_1.5.0_linux_amd64 -o /usr/local/bin/cfssljson && \
chmod +x /usr/local/bin/cfssl && \
chmod +x /usr/local/bin/cfssljson

#generate ca in /tmp
cfssl gencert -initca ./tls/ca-csr.json | cfssljson -bare /tmp/ca

#generate certificate in /tmp
cfssl gencert \
  -ca=/tmp/ca.pem \
  -ca-key=/tmp/ca-key.pem \
  -config=./tls/ca-config.json \
  -hostname="tcpdump-webhook,tcpdump-webhook.default.svc.cluster.local,tcpdump-webhook.default.svc,localhost,127.0.0.1" \
  -profile=default \
  ./tls/ca-csr.json | cfssljson -bare /tmp/example-webhook


#make a secret
cat <<EOF > ./manifests/webhook-secrets.yaml
apiVersion: v1
kind: Secret
metadata:
  name: tcpdump-webhook-tls
type: Opaque
data:
  tls.crt: $(cat /tmp/example-webhook.pem | base64 | tr -d '\n')
  tls.key: $(cat /tmp/example-webhook-key.pem | base64 | tr -d '\n')
EOF

#generate CA Bundle + inject into template
ca_pem_b64="$(openssl base64 -A <"/tmp/ca.pem")"

sed -e 's@${CA_PEM_B64}@'"$ca_pem_b64"'@g' <"./manifests/webhook-configuration-template.yaml" \
    > ./manifests/webhook-configuration.yaml

exit

kubectl apply -f ./manifests/webhook-secrets.yaml
kubectl apply -f ./manifests/webhook-rbac.yaml
kubectl apply -f ./manifests/webhook-deployment-svc.yaml
kubectl apply -f ./manifests/webhook-configuration.yaml

kubectl apply -f ./test/demo-pod.yaml

