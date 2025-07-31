# API

## Cluster

Params Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ClusterCreateRequestUnionParam">ClusterCreateRequestUnionParam</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#ClusterPatchRequestParam">ClusterPatchRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterNewResponse">APIClusterNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterGetResponse">APIClusterGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterUpdateResponse">APIClusterUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterListResponse">APIClusterListResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterDeleteResponse">APIClusterDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterGetSSHKeyResponse">APIClusterGetSSHKeyResponse</a>

Methods:

- <code title="post /api/cluster">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterNewParams">APIClusterNewParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterNewResponse">APIClusterNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/cluster/{cluster_id_or_alias}">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterGetResponse">APIClusterGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/cluster/{cluster_id_or_alias}">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterUpdateParams">APIClusterUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterUpdateResponse">APIClusterUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/cluster">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterListResponse">APIClusterListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/cluster/{cluster_id_or_alias}">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterDeleteResponse">APIClusterDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/cluster/{cluster_id_or_alias}/ssh_key">client.API.Cluster.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterService.GetSSHKey">GetSSHKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clusterIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIClusterGetSSHKeyResponse">APIClusterGetSSHKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Vm

Params Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmBranchRequestParam">VmBranchRequestParam</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmCommitRequestParam">VmCommitRequestParam</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmPatchRequestParam">VmPatchRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmDeleteResponse">VmDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmDto">VmDto</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmGetResponse">APIVmGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmUpdateResponse">APIVmUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmListResponse">APIVmListResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmDeleteResponse">APIVmDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmBranchResponse">APIVmBranchResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmCommitResponse">APIVmCommitResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmGetSSHKeyResponse">APIVmGetSSHKeyResponse</a>

Methods:

- <code title="get /api/vm/{vm_id_or_alias}">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmGetResponse">APIVmGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/vm/{vm_id_or_alias}">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmUpdateParams">APIVmUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmUpdateResponse">APIVmUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/vm">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmListResponse">APIVmListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/vm/{vm_id_or_alias}">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmDeleteParams">APIVmDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmDeleteResponse">APIVmDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/vm/{vm_id_or_alias}/branch">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Branch">Branch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmBranchParams">APIVmBranchParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmBranchResponse">APIVmBranchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/vm/{vm_id_or_alias}/commit">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.Commit">Commit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmCommitParams">APIVmCommitParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmCommitResponse">APIVmCommitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/vm/{vm_id_or_alias}/ssh_key">client.API.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmService.GetSSHKey">GetSSHKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIVmGetSSHKeyResponse">APIVmGetSSHKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rootfs

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfListResponse">APIRootfListResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfDeleteResponse">APIRootfDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfUploadResponse">APIRootfUploadResponse</a>

Methods:

- <code title="get /api/rootfs">client.API.Rootfs.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfListResponse">APIRootfListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/rootfs/{rootfs_id}">client.API.Rootfs.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rootfsID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfDeleteResponse">APIRootfDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/rootfs/{rootfs_id}">client.API.Rootfs.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, rootfsID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfUploadParams">APIRootfUploadParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIRootfUploadResponse">APIRootfUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Health

Methods:

- <code title="get /api/health">client.API.Health.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APIHealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Telemetry

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#TelemetryDto">TelemetryDto</a>

Methods:

- <code title="get /api/telemetry">client.API.Telemetry.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#APITelemetryService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#TelemetryDto">TelemetryDto</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
