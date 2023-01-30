#!/bin/bash -ue

secret=$1

key=age1a3f9l459k30wqnyg4rerhpg432xq23r085n2uk278svp2t3zjv3qdxch3u

if [[ -z "${secret}" ]]; then echo "Please specify secret"; exit 1; fi

yq -i "(del.sops)" ./k8s/secret/${secret}.yaml

yq -i ".data.${secret}=\"$(cat k8s/secret/${secret}.in.txt | base64 | tr -d '\n')\"" k8s/secret/${secret}.yaml

sops --encrypt \
	--age ${key} \
	--encrypted-regex '^(data|stringData)$$' \
	--in-place \
	k8s/secret/${secret}.yaml

yamlfmt k8s/secret/${secret}.yaml
