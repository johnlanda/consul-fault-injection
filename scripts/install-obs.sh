#!/bin/zsh

# Install cert manager (required jaeger prerequisite)
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.6.3/cert-manager.yaml
kubectl rollout status deployment -n cert-manager cert-manager
kubectl rollout status deployment -n cert-manager cert-manager-cainjector
kubectl rollout status deployment -n cert-manager cert-manager-webhook

# Install the jaeger operator and CRDs
kubectl create namespace observability
kubectl apply -f https://github.com/jaegertracing/jaeger-operator/releases/download/v1.51.0/jaeger-operator.yaml -n observability
kubectl rollout status deployment -n observability jaeger-operator

# Install jaeger all in one
kubectl apply -n observability -f ./charts-obs/jaeger.yaml
kubectl rollout status deployment -n observability jaeger