# Autoversioned

![Version: 2.4.2](https://img.shields.io/badge/Version-2.4.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.4.2](https://img.shields.io/badge/AppVersion-2.4.2-informational?style=flat-square)

Reference project with for versioning approach demo

## Components

* Simple golang web server packaged into a docker image
* Helm chart with helmfile and skaffold wrappers on top for deployment and local development respectively

## How to run

### Install helm package

* `helm repo add autoversioned https://mtrqq.github.io/autoversioned`
* `helm install auto autoversioned/autoversioned [--version <version>]`

### Local run via skaffold

* `skaffold dev --port-forward`

### Deploy using helmfile based on the environment

* `helmfile apply [-e (dev|staging|prod)]`

### Run using docker

* `docker run -p 8080:8080 mtrqq/autoversioned:<version>`

### Run using golang

* `go mod download`
* `go run api/serve.go`

## Versioning

### Commits

Versioning is mostly based on [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification and powered by [commitizen](https://commitizen-tools.github.io/commitizen/) (note: there's a bunch of projects with conflicting names so it's better to use the direct link). The only difference for a developer when working with such kind of repository is that you have to follow certain commit naming conventions, here are some examples

* `feat(core): implement the ability to fetch cluster statistics`
* `fix: decrease memory footprint for statistics service`
* `chore: bump build-and-push action version`
* `refactor: migrate from log to zerolog`

It doesn't affect developers' life in case the team is using squash and merge strategy for PRs as the only thing to track in this case is merge commit naming. Otherwise, if you decided to go the hard way with rebase strategy - you should check each commit message which is going into remote (commit-msg `pre-commit` hooks may come in handy in this case).

### Release process

When we decide to release an application at a certain moment of development - we should manually trigger [release pipeline](.github/workflows/release.yaml), most of the time you don't have to tune any parameters as it's an automatical strategy is already pretty good, but anyway they are present in case you need them. Here's what this workflow does in brief (version 1.2.3 is used only for reference purposes):

* Creates `release/1.2.3` with already patched versions and changelog
* Creates pre-approved PR from `release/1.2.3` branch so that you can review all the changes done
* Creates draft release for tag `v1.2.3` with incremental changelog based on commits
* Builds the following artifacts:
  * Linux-based executables for API server on the most used architectures (arm, 386, amd64) and attaches them to a release
  * API server docker image pushes it to dockerhub repository, its' `imageid` is attached to release as a file
  * Helm chart package is created and added to a [GitHub Pages based repo index](https://mtrqq.github.io/autoversioned)

## Helm chart

# autoversioned

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"mtrqq/autoversioned"` |  |
| image.tag | string | `""` |  |
| imagePullSecrets | list | `[]` |  |
| ingress.annotations | object | `{}` |  |
| ingress.className | string | `""` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hosts[0].host | string | `"chart-example.local"` |  |
| ingress.hosts[0].paths[0].path | string | `"/"` |  |
| ingress.hosts[0].paths[0].pathType | string | `"ImplementationSpecific"` |  |
| ingress.tls | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources | object | `{}` |  |
| secretValue | string | `nil` |  |
| securityContext | object | `{}` |  |
| service.port | int | `80` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` |  |
| tolerations | list | `[]` |  |