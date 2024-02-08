# Consul Fault Injection Demo

## Create local kind cluster and registry

```bash
./scripts/create-kind-cluster-with-registry.sh
```

## Build Consul Enterprise docker image

```bash
cd $CONSUL_ENTERPRISE_REPO
make dev-docker
docker tag 'hashicorp/consul-enterprise:local' 'localhost:5001/consul-enterprise:local'
docker push 'localhost:5001/consul-enterprise:local'
```

## Build and publish test applications

```bash
cd $CONSUL_FAULT_INJECTION_DEMO_REPO
(cd services && make all)
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

## Install the Consul CLI

```bash
brew install consul
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
kubectl apply -f ./config/heartbeat.yaml && kubectl apply -f ./config/dashboard.yaml
```

## Test the demo application

```bash
kubectl port-forward svc/dashboard --namespace default 9002:9002
```

Open http://localhost:9002 in your browser, and you should that there are no current requests reaching the heartbeat service.

## Create the intentions to allow the services to communicate

```bash
kubectl apply -f ./config/intentions.yaml
```

## Re-test the demo application

Ensure that the dashboard port-forward is still running, and refresh the page. The dashboard should now show successful requests.

## Create the proxy, mesh, and service defaults

Apply the defaults via either `kubectl` or the `consul config write` command.

```bash
kubectl apply -f ./config/defaults/yaml/mesh-defaults.yaml
kubectl apply -f ./config/defaults/yaml/proxy-defaults.yaml
kubectl apply -f ./config/defaults/yaml/dashboard-defaults.yaml
kubectl apply -f ./config/defaults/yaml/heartbeat-defaults.yaml
```

Or alternatively:

```bash
consul config write ./config/defaults/hcl/mesh-defaults.hcl
consul config write ./config/defaults/hcl/proxy-defaults.hcl
consul config write ./config/defaults/hcl/dashboard-defaults.hcl
consul config write ./config/defaults/hcl/heartbeat-defaults.hcl
```

Note that if you create the defaults via `consul config write`, you will not be able to overwrite them later via `kubectl apply` as the reconciler will fail.

The dashboard should now show all requests going through the proxies.

## Create the fault injection filters

```bash
kubectl apply -f ./config/fault-injection/yaml/heartbeat-fault-injection.yaml
```

Or alternatively:

```bash
consul config write ./config/fault-injection/hcl/heartbeat-fault-injection.hcl
```

Note that 50% of the requests will now have a 500 status code injected from the filter.

## Change to a delay injection filter

```bash
kubectl apply -f ./config/fault-injection/yaml/heartbeat-delay-injection.yaml
```

Or alternatively:

```bash
consul config write ./config/fault-injection/hcl/heartbeat-delay-injection.hcl
```

Note that statuses are all 200 now, however, 50% of requests are delayed by 500 milliseconds.

## Cleanup

```bash
kind delete cluster -n dc1
```
