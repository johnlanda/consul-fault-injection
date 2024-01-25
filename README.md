# Consul Fault Injection Demo

## Create local kind cluster and registry

```bash
./scripts/create-kind-cluster-with-registry.sh
```

## Install observability stack

```bash
./scripts/install-obs.sh
```

## Build Consul Enterprise docker image

```bash
cd $CONSUL_ENTERPRISE_REPO
make dev-docker
docker tag 'hashicorp/consul-enterprise:local' 'localhost:5001/consul-enterprise:local'
docker push 'localhost:5001/consul-enterprise:local'
```

## Create the Consul namespace and license secret

```bash
kubectl create ns consul
# Given consul license in CONSUL_LICENSE env var
kubectl create secret generic consul-ent-license --from-literal="key=${CONSUL_LICENSE}" -n consul
```

## Install Consul in the cluster

```bash
brew tap hashicorp/tap
brew install hashicorp/tap/consul-k8s
consul-k8s install -config-file=./config/values.yaml
```

## Configure CLI to interact with Consul

```bash
export CONSUL_HTTP_TOKEN=$(kubectl get --namespace consul secrets/consul-bootstrap-acl-token --template={{.data.token}} | base64 -d)
export CONSUL_HTTP_ADDR=https://127.0.0.1:8501
export CONSUL_HTTP_SSL_VERIFY=false
```

In a new window, port forward to allow the CLI to connect to consul.
```bash
kubectl port-forward svc/consul-ui --namespace consul 8501:443
```

## Verify the license
```bash
consul license get
```

## Deploy the services

```bash
kubectl apply -f ./config/counting.yaml && kubectl apply -f ./config/dashboard.yaml
```

## Create the proxy, mesh, and service defaults

```bash
consul config write ./config/defaults/mesh-defaults.hcl
consul config write ./config/defaults/proxy-defaults.hcl
consul config write ./config/defaults/dashboard-defaults.hcl
consul config write ./config/defaults/counting-defaults.hcl
```

## Test the demo application

```bash
kubectl port-forward svc/dashboard --namespace default 9002:9002
```

Open http://localhost:9002 in your browser, and you should see that the counting service is not available. The dashboard
will show a count of `-1`.

## Create the intentions to allow the services to communicate

```bash
kubectl apply -f ./config/intentions.yaml
```

## Re-test the demo application

Ensure that the dashboard port-forward is still running, and refresh the page. The dashboard should now show an increasing count.

## View trace data of functioning services

```bash
kubectl port-forward -n observability <jaeger-pod> 16686:16686
```

Navigate to http://localhost:16686 and view the traces.

## Create the fault injection filters

```bash

```

## Observe failures in tracing

Ensure that the jaeger port-forward is still running, and navigate to http://localhost:16686. You should now be able to
see requests with injected failures or delays based on the configured fault injection filters.

## Cleanup

```bash
kind delete cluster -n dc1
```
