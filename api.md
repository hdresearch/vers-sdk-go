# API

## Cluster

Params Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#CreateParam">CreateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Cluster">Cluster</a>

Methods:

- <code title="post /api/cluster">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterNewParams">APIClusterNewParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/cluster/{cluster_id}">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/cluster">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/cluster/{cluster_id}">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Cluster">Cluster</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/cluster/{cluster_id}/ssh_key">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.GetSSHKey">GetSSHKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Vm

Params Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ExecuteCommandParam">ExecuteCommandParam</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#PatchRequestParam">PatchRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ExecuteResponse">ExecuteResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>

Methods:

- <code title="get /api/vm/{vm_id}">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/vm">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/vm/{vm_id}">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmDeleteParams">APIVmDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/vm/{vm_id}/branch">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Branch">Branch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/vm/{vm_id}/commit">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Commit">Commit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/vm/{vm_id}/execute">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmExecuteParams">APIVmExecuteParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ExecuteResponse">ExecuteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/vm/{vm_id}/ssh_key">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.GetSSHKey">GetSSHKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/vm/{vm_id}">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.UpdateState">UpdateState</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmUpdateStateParams">APIVmUpdateStateParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Vm">Vm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rootfs

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#DeleteResponse">DeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ListResponse">ListResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#UploadResponse">UploadResponse</a>

Methods:

- <code title="get /api/rootfs">client.API.Rootfs.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ListResponse">ListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/rootfs/{rootfs_id}">client.API.Rootfs.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rootfsID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#DeleteResponse">DeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/rootfs/{rootfs_id}">client.API.Rootfs.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rootfsID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#UploadResponse">UploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Health

Methods:

- <code title="get /api/health">client.API.Health.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIHealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Network

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Info">Info</a>

Methods:

- <code title="get /api/network">client.API.Network.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APINetworkService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#Info">Info</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
