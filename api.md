# Orchestrator

## Vm

Params Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#NewRootRequestParam">NewRootRequestParam</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmFromCommitRequestParam">VmFromCommitRequestParam</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmUpdateStateRequestParam">VmUpdateStateRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#NewVmResponse">NewVmResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmCommitResponse">VmCommitResponse</a>
- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmDeleteResponse">VmDeleteResponse</a>

Methods:

- <code title="delete /vm/{vm_id}">client.Orchestrator.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmDeleteResponse">VmDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /vm/{vm_id}/branch">client.Orchestrator.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmService.Branch">Branch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#NewVmResponse">NewVmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /vm/{vm_id}/commit">client.Orchestrator.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmService.Commit">Commit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmCommitResponse">VmCommitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /vm/new_root">client.Orchestrator.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmService.NewRoot">NewRoot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmNewRootParams">OrchestratorVmNewRootParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#NewVmResponse">NewVmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /vm/from_commit">client.Orchestrator.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmService.RestoreFromCommit">RestoreFromCommit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmRestoreFromCommitParams">OrchestratorVmRestoreFromCommitParams</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#NewVmResponse">NewVmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /vm/{vm_id}/state">client.Orchestrator.Vm.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmService.UpdateState">UpdateState</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vmID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorVmUpdateStateParams">OrchestratorVmUpdateStateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Node

Response Types:

- <a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmListAllResponse">VmListAllResponse</a>

Methods:

- <code title="get /node/{node_id}/vms">client.Orchestrator.Node.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#OrchestratorNodeService.ListVms">ListVms</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go">vers</a>.<a href="https://pkg.go.dev/github.com/hdresearch/vers-sdk-go#VmListAllResponse">VmListAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
