#!/bin/sh -eu

# https://www.okteto.com/docs/cloud/okteto-cli/#built-in-tools-when-deploying-to-okteto-cloud

export ENCODED_MONGODB_URI=$(echo "${GOTAGNEWS_MONGODB_URI}" | base64 | tr -d '\n')

export ENCODED_LINE_CHANNEL_SECRET=$(echo "${GOTAGNEWS_LINE_CHANNEL_SECRET}" | base64 | tr -d '\n')

export ENCODED_LINE_CHANNEL_TOKEN=$(echo "${GOTAGNEWS_LINE_CHANNEL_TOKEN}" | base64  | tr -d '\n')

(cd k8s && envsubst < secret.yaml.template > secret.yaml)

kubectl -n gotagnews-takokun778 apply -f ./k8s
