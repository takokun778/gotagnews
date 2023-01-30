#!/bin/bash -ue

sops --decrypt k8s/secret/${secret}.yaml > k8s/secret/${secret}.tmp.yaml

yq ".data.${secret}" k8s/secret/${secret}.tmp.yaml | base64 --decode > k8s/secret/${secret}.out.txt

rm k8s/secret/${secret}.tmp.yaml
