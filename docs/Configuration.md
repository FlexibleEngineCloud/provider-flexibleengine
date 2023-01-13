---
title: Configuration
weight: 2
---

# FlexibleEngine provider documentation

## Install the provider

Install the FlexibleEngine provider with the following configuration file

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-flexibleengine
spec:
  package: xpkg.upbound.io/frangipaneteam/provider-flexibleengine:<version>
```

Define the provider version with `spec.package`.

Install the provider with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
$ kubectl get providers
NAME                      INSTALLED   HEALTHY   PACKAGE                                                        AGE
provider-flexibleengine   True        True      xpkg.upbound.io/frangipaneteam/provider-flexibleengine:v0.1.2  62s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/FrangipaneTeam/provider-flexibleengine) to view all available `Provider` options.

## Configure the provider

The FlexibleEngine provider requires credentials for authentication to Flexible Engine. This can be done in one of the following ways:

* Authenticating using AK/SK

### Generate a Kubernetes Secret

Create a JSON file containing the Flexible Engine account credentials (AK/SK).

Here is an example key file:

```json
{
    "access_key": "your access key",
    "secret_key": "your secret key"
}
```

Use the JSON file to generate a Kubernetes secret.

```shell
kubectl -n upbound-system create secret generic fe-creds --from-file=credentials=./<JSON file name>
```

### Create a ProviderConfig object

Apply the secret in a `ProviderConfig` Kubernetes configuration file.

```yaml
apiVersion: flexibleengine.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  region: eu-west-0
  domainName: <DOMAIN_NAME>
  credentials:
    source: Secret
    secretRef:
      name: fe-creds
      namespace: upbound-system
      key: credentials
```

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl create secret generic <name>` command.

You can also use the `spec.tenantName` to specify the tenant name in the Flexible Engine account.

```yaml
apiVersion: flexibleengine.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  region: eu-west-0
  domainName: <DOMAIN_NAME>
  tenantName: <TENANT_NAME>
  credentials:
    source: Secret
    secretRef:
      name: fe-creds
      namespace: upbound-system
      key: credentials
```

If you want to use an insecure connection, you can set the `spec.insecure` to `true`.

```yaml
apiVersion: flexibleengine.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  region: eu-west-0
  domainName: <DOMAIN_NAME>
  insecure: true
  credentials:
    source: Secret
    secretRef:
      name: fe-creds
      namespace: upbound-system
      key: credentials
```
