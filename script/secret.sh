#!/bin/bash -ue

secret=$1

cat << EOF > k8s/secret/${secret}.yaml
apiVersion: v1
kind: Secret
metadata:
  name: gotagnews-secret-${secret}
type: Opaque
data:
  ${secret}: ""
EOF

touch k8s/secret/${secret}.in.txt

touch k8s/secret/${secret}.out.txt
