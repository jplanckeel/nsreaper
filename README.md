# nsreaper
 Find previews namespace and delete after a time.


## Purpose

This tool call k8s list previews namespace and deletes them if the ttl is exceeded. If you deploy with Helm3, all elements is in namespace, delete namespace remove all k8s object in this namespace.

Warning: if you deploy with helm2, remove release before.

## Requierments

For detecte namespaces previews we need annotations annotations on namespace

|name|values|mandatory|description|
|--|--|--|--|
|nsreaper/type| preview | required | annotation for detect previews namespace
|nsreaper/repository| project-a | optional | name of project for logs
|nsreaper/pull_request_id| "10" | optional | number of PR for logs
|nsreaper/ttl| 14 | optional | specify a different TTL 


## Usage

```shell
$ nsreaper --help
a cli for clean previews namespace after a ttl, by default 10 day:
namespace clean --dry-run true --ttl 10

Usage:
  nsreaper [command]

Available Commands:
  clean       Clean Preview Namespace on 
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help      help for preview-reaper
  -v, --version   version for preview-reaper

Use "nsreaper [command] --help" for more information about a command.


## Dry run mode

$nsreaper clean --dry-run true

Namespace: project-a-pr10, Creationdate: 2022-01-20 17:19:33 +0100 CET, Repo project-a, Pr:10 will be delete 
Namespace: project-b-pr55, Creationdate: 2022-01-20 17:19:33 +0100 CET, Repo project-b, Pr:55 will be delete 

```

## Build binnary 


```shell
cd src 
go build -ldflags "-s -w"

```
