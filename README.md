## Description

Helm chart for deploying `autoversioned` application components into K8S cluster

## Homepage

<https://github.com/mtrqq/autoversioned>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Mtrqq, Inc. | <maxym.fugol@gmail.com> | <https://github.com/mtrqq> |

## Pre-requisites

- Kubernetes cluster of version 1.18 and up, 1.19 and up recommended
- Helm v3 must be installed and configured properly

## Deployment

```shell
helm repo add autoversioned https://mtrqq.github.io/autoversioned
helm install autoversioned autoversioned/autoversioned
```
