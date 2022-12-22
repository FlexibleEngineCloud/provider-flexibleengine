---
title: Configuration
weight: 2
---

# FlexibleEngine provider documentation

## Install the provider

### Prerequisites

#### Upbound Up command-line

The Upbound Up command-line simplifies configuration and management of Upbound
Universal Crossplane (UXP) and interacts with the Upbound Marketplace to manage
users and accounts.

Install `up` with the command:

```shell
curl -sL "https://cli.upbound.io" | sh
```

More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

#### Upbound Universal Crossplane

UXP is the Upbound official enterprise-grade distribution of Crossplane for
self-hosted control planes. Only Upbound Universal Crossplane (UXP) supports
official providers.

Official providers aren't supported with open source Crossplane.

Install UXP into your Kubernetes cluster using the Up command-line.

```shell
up uxp install
```

Find more information in the [Upbound UXP documentation](https://docs.upbound.io/uxp/).

### Install the provider

Install the Upbound FlexibleEngine provider with the following configuration file

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
provider-flexibleengine   True        True      xpkg.upbound.io/frangipaneteam/provider-flexibleengine:v0.0.1  62s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/FrangipaneTeam/provider-flexibleengine) to view all available `Provider` options.

## Configure the provider

The FlexibleEngine provider requires credentials for authentication to Flexible Engine. This can be done in one of the following ways:

* Authenticating using a base-64 encoded service account key in a Kubernetes Secret.
* Authenticating using AK/SK

### Generate a Kubernetes Secret

Create a JSON file containing the Flexible Engine account credentials.

Here is an example key file:

```json
{
    "access_key": "your access key",
    "secret_key": "your secret key"
}
```

Use the JSON file to generate a Kubernetes secret.

```shell
kubectl create secret generic fe-creds --from-file=credentials=./<JSON file name>
```

### Create a ProviderConfig object

Apply the secret in a `ProviderConfig` Kubernetes configuration file.

```yaml
apiVersion: flexibleengine.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  region: <REGION>
  domainName: <DOMAIN_NAME>
  credentials:
    source: Secret
    secretRef:
      name: fe-creds
      namespace: upbound-system
      key: credentials
```

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl create secret generic <name>` command.
