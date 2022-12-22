---
title: Quickstart
weight: 1
---

# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and install the FlexibleEngine provider.

To use this provider, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create a *managed resource* in Flexible Engine via Kubernetes.

## Install the Up command-line

Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.0.1
```

*Note*: providers only support `up` command-line versions v0.13.0 or later.

## Install Universal Crossplane

Install Upbound Universal Crossplane with the Up command-line.

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                        READY   STATUS    RESTARTS      AGE
crossplane-7fdfbd897c-pmrml                 1/1     Running   0             68m
crossplane-rbac-manager-7d6867bc4d-v7wpb    1/1     Running   0             68m
upbound-bootstrapper-5f47977d54-t8kvk       1/1     Running   0             68m
xgql-7c4b74c458-5bf2q                       1/1     Running   3 (67m ago)   68m
```

## Install the FlexibleEngine provider

Install the provider into the Kubernetes cluster with a Kubernetes configuration file.

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-flexibleengine
spec:
  package: xpkg.upbound.io/frangipaneteam/provider-flexibleengine:<version>
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.

```shell
$ kubectl get providers
NAME                      INSTALLED   HEALTHY   PACKAGE                                                        AGE
provider-flexibleengine   True        True      xpkg.upbound.io/frangipaneteam/provider-flexibleengine:v0.0.1  62s
```

It may take up to 5 minutes to report `HEALTHY`.

## Create a Kubernetes secret

The provider requires credentials to create and manage GCP resources.

### Generate a GCP JSON key file

Create a JSON key file containing the GCP account credentials. GCP provides documentation on [how to create a key file](https://cloud.google.com/iam/docs/creating-managing-service-account-keys).

Here is an example key file:  

```json
{
  "type": "service_account",
  "project_id": "caramel-goat-354919",
  "private_key_id": "e97e40a4a27661f12345678f4bd92139324dbf46",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCwA+6MWRhmcPB3\nF/irb5MDPYAT6BWr7Vu/16U8FbCHk7xtsAWYjKXKHu5mGzum4F781sM0aMCeitlv\n+jr2y7Ny23S9uP5W2kfnD/lfj0EjCdfoaN3m7W0j4DrriJviV6ESeSdb0Ehg+iEW\ngNrkb/ljigYgsSLMuemby5lvJVINUazXJtGUEZew+iAOnI4/j/IrDXPCYVNo5z+b\neiMsDYWfccenWGOQf1hkbVWyKqzsInxu8NQef3tNhoUXNOn+/kgarOA5VTYvFUPr\n2l1P9TxzrcYuL8XK++HjVj5mcNaWXNN+jnFpxjMIJOiDJOZoAo0X7tuCJFXtAZbH\n9P61GjhbAgMBAAECggEARXo31kRw4jbgZFIdASa4hAXpoXHx4/x8Q9yOR4pUNR/2\nt+FMRCv4YTEWb01+nV9hfzISuYRDzBEIxS+jyLkda0/+48i69HOTAD0I9VRppLgE\ne97e40a4a27661f12345678f4bd92139324dbf46+2H7ulQDtbEgfcWpNMQcL2JiFq+WS\neh3H0gHSFFIWGnAM/xofrlhGsN64palZmbt2YiKXcHPT+WgLbD45mT5j9oMYxBJf\nPkUUX5QibSSBQyvNqCgRKHSnsY9yAkoNTbPnEV0clQ4FmSccogyS9uPEocQDefuY\nY7gpwSzjXpaw7tP5scK3NtWmmssi+dwDadfLrKF7oQKBgQDjIZ+jwAggCp7AYB/S\n6dznl5/G28Mw6CIM6kPgFnJ8P/C/Yi2y/OPKFKhMs2ecQI8lJfcvvpU/z+kZizcG\nr/7iRMR/SX8n1eqS8XfWKeBzIdwQmiKyRg2AKelGKljuVtI8sXKv9t6cm8RkWKuZ\n9uVroTCPWGpIrh2EMxLeOrlm0QKBgQDGYxoBvl5GfrOzjhYOa5GBgGYYPdE7kNny\nhpHE9CrPZFIcb5nGMlBCOfV+bqA9ALCXKFCr0eHhTjk9HjHfloxuxDmz34vC0xXG\ncegqfV9GNKZPDctysAlCWW/dMYw4+tzAgoG9Qm13Iyfi2Ikll7vfeMX7fH1cnJs0\nnYpN9LYPawKBgQCwMi09QoMLGDH+2pLVc0ZDAoSYJ3NMRUfk7Paqp784VAHW9bqt\n1zB+W3gTyDjgJdTl5IXVK+tsDUWu4yhUr8LylJY6iDF0HaZTR67HHMVZizLETk4M\nLfvbKKgmHkPO4NtG6gEmMESRCOVZUtAMKFPhIrIhAV2x9CBBpb1FWBjrgQKBgQCj\nkP3WRjDQipJ7DkEdLo9PaJ/EiOND60/m6BCzhGTvjVUt4M22XbFSiRrhXTB8W189\noZ2xrGBCNQ54V7bjE+tBQEQbC8rdnNAtR6kVrzyoU6xzLXp6Wq2nqLnUc4+bQypT\nBscVVfmO6stt+v5Iomvh+l+x05hAjVZh8Sog0AxzdQKBgQCMgMTXt0ZBs0ScrG9v\np5CGa18KC+S3oUOjK/qyACmCqhtd+hKHIxHx3/FQPBWb4rDJRsZHH7C6URR1pHzJ\nmhCWgKGsvYrXkNxtiyPXwnU7PNP9JNuCWa45dr/vE/uxcbccK4JnWJ8+Kk/9LEX0\nmjtDm7wtLVlTswYhP6AP69RoMQ==\n-----END PRIVATE KEY-----\n",
  "client_email": "my-sa-313@caramel-goat-354919.iam.gserviceaccount.com",
  "client_id": "103735491955093092925",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/my-sa-313%40caramel-goat-354919.iam.gserviceaccount.com"
}
```

Save this JSON file as `gcp-credentials.json`.

### Create a Kubernetes secret with GCP credentials

Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the Universal Crossplane cluster.

`kubectl create secret generic gcp-secret -n upbound-system --from-file=creds=./gcp-credentials.json`

View the secret with `kubectl describe secret`

```shell
$ kubectl describe secret gcp-secret -n upbound-system
Name:         gcp-secret
Namespace:    upbound-system
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
creds:  2334 bytes
```

## Create a ProviderConfig

Create a `ProviderConfig` Kubernetes configuration file to attach the GCP credentials to the installed official provider.

**Note:** the `ProviderConfig` must contain the correct GCP project ID. The project ID must match the `project_id` from the JSON key file.

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

Apply this configuration with `kubectl apply -f`.

**Note:** the `ProviderConfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.secretRef.key` must match the value in the `Data` section of the secret.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`.

```yaml
$ kubectl describe providerconfigs
Name:         default
API Version:  flexibleengine.upbound.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        credentials
      Name:       fe-secret
      Namespace:  upbound-system
    Source:       Secret
  Domain Name:    <DOMAIN_NAME>
  Region:         <REGION>
```

## Create a managed resource

Create a managed resource to verify the provider is functioning.

This example creates a FlexibleEngine OBS storage bucket, which requires a globally unique name.

Generate a unique bucket name from the command line.

`echo "flexibleengine-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 10)`

For example

```shell
$ echo "flexibleengine-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 10)
flexibleengine-bucket-4a917c947
```

Use this bucket name for `spec.forProvider.bucket` value.

Create a `Bucket` configuration file. Replace `<BUCKET NAME>` with the `flexibleengine-bucket-` generated name.

```yaml
apiVersion: oss.flexibleengine.upbound.io/v1beta1
kind: OBSBucket
metadata:
  name: example-obsbucket
spec:
  forProvider:
    acl: private
    bucket: <BUCKET NAME>
    storageClass: STANDARD
  providerConfigRef:
    name: default
```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get obsbucket` to verify bucket creation.

```shell
$ kubectl get obsbucket
NAME                READY   SYNCED   EXTERNAL-NAME                     AGE
example-obsbucket   True    True     flexibleengine-bucket-4a917c947   90s
```

Provider created the bucket when the values `READY` and `SYNCED` are `True`.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.

Here is an example of a failure because the `Bucket` `spec.providerConfigRef.name` value doesn't match the `ProviderConfig` `metadata.name`.

```shell
$ kubectl describe obsbucket
Name:         example-obsbucket
Labels:       <none>
Annotations:  crossplane.io/external-name: flexibleengine-bucket-4a917c947
API Version:  oss.flexibleengine.upbound.io/v1beta1
Kind:         OBSBucket
# Output truncated
Spec:
  Deletion Policy:  Delete
  For Provider:
    Acl:            private
    Bucket:         flexibleengine-bucket-4a917c947
    Region:         eu-west-0
    Storage Class:  STANDARD
  Provider Config Ref:
    Name:  default
Status:
  At Provider:
  Conditions:
    Last Transition Time:  2022-07-26T19:29:35Z
    Message:               connect failed: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.flexibleengine.upbound.io "default" not found
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age               From                                                 Message
  ----     ------                   ----              ----                                                 -------
  Warning  CannotConnectToProvider  7s (x5 over 20s)  managed/oss.flexibleengine.upbound.io/v1beta1, kind=obsbucket  cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.flexibleengine.upbound.io "default" not found
```

The output indicates the `OBSBucket` is using `ProviderConfig` named `default`. The applied `ProviderConfig` is `my-config`.

```shell
$ kubectl get providerconfigs
NAME        AGE
my-config   56s
```

## Delete the managed resource

Remove the managed resource by using `kubectl delete -f` with the same `OBSBucket` object file. Verify the removal of the bucket with `kubectl get obsbucket`.

```shell
$ kubectl get obsbucket
No resources found
```
