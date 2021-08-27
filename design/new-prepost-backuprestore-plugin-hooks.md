# PreBackup, PostBackup, PreBackup and PostBackup Action Plugins

## Abstract

Velero should provide a way to trigger actions before and after each Velero backup and each Velero restore.
Important: this is fundamentally different from the existent plugin hooks of BackupItemAction and RestoreItemAction, which are triggered per item during backup and restore. The proposed new plugins are executed only once.

These plugins would be invoked:

- PreBackupAction: plugin is executed right after the backup object is created but the backup is not yet in progress (before runBackup).
- PostBackupAction: plugin is executed once the backup object finishes executing (after runBackup) and after receiving its final status.
- PreRestoreAction: plugin is executed when a restore object is created but not yet in progress.
- PostRestoreAction: plugin is executed when a restore object is finished.

## Background

More and more, Velero is employed for workload migrations across different Kubernetes clusters.
Using Velero for migrations is an atomic operation involving a Velero backup on a source cluster sequenced by a Velero restore on a destination cluster.

It is common during these migrations to perform many actions inside and outside Kubernetes clusters.
Attention: these actions are not necessarily per resource item, but they are actions to be executed *once* before and/or after the migration itself (remember, migration in this context is Velero Backup + Velero Restore).

One important use case driving this proposal is migrating stateful workloads in scale across different clusters/storage backends.
Today, Velero's Restic integration is the response for such use cases, but there are some limitations:

- Quiesce/unquiesce workloads: Podhooks are useful for quiesce/unquiesce workloads, but we constantly observe that platform engineers do not have the luxury/visibility/time/knowledge to go to each pod add specific commands to quiesce/unquiesce workloads.
- Orphan PVC/PV pairs:  PVC/PV that do not have a running pod are not backed up and consequently not migrated.
  
Aiming to address these two limitations, we want to write a velero plugin to be executed *once* (not per resource item) prior backup to scale down the applications setting .spec.replicas=0 to all deployments, statefulsets, daemonsets, replicasets, etc. and starting a small footprint staging pod to mount all PVC/PV pairs. Similarly, we want to write another plugin to unquiesce the applications (killing the staging pod and reinstating original .spec.replicas values) after the Velero restore is completed.

Other examples of the usability of these proposed PreBackupAction, PostBackupAction, PreRestoreAction, and PostRestoreAction:

PostBackupAction: triggering a Velero Restore after a successful Velero backup (so we can complete the migration operation).
PreRestoreAction: pre-expanding the cluster's capacity via Cluster API to avoid starvation of cluster resources before the restore.
PostRestoreAction: call some actions to be performed outside Kubernetes clusters, such as configure a global load balancer (GLB) to enable the new cluster.

This design seeks to provide the missing extension points. This proposal's scope is only adding the new plugin hooks, not the plugins themselves.

## Goals

- Provide a PreBackupAction, PostBackupAction, PreRestoreAction and PostRestoreAction API for plugins to implement
- Update Velero backup and restore creation logic to invoke registered PreBackupAction and PreRestoreAction plugins before processing the backup and restore respectively.
- Update Velero backup and restore complete logic to invoke registered PostBackupAction and PostRestoreAction plugins after flagging the objects as completed.

## Non-Goals

- Specific implementations of the PreBackupAction, PostBackupAction, PreRestoreAction and PostRestoreAction API beyond test cases.

## High-Level Design

The PreBackupAction plugin API will resemble the BackupItemAction plugin design, but with the fundamental difference that will receive only the Velero `Backup` Go struct created. It will not receive any resource list item because the backup is not yet running at that stage. In addition, the `PreBackupAction` interface will only have `Execute()` since the plugin will be executed once per Backup creation, not per item.

The Velero backup controller will be modified so that if there are any PreBackupAction plugins registered, it will be executed as the last step of backup validation. Then, each PreBackupAction plugin will be executed.
If any PreBackupAction returns an error, the backup object will not be valid, and consequently, the backup controller will never call `c.runBackup`. In other words, the backup will not be executed if the PreBackup plugin returns an error.

PreBackupAction plugins will be run in alphanumeric order based on their registered names.

The PostBackupAction plugin API will resemble the BackupItemAction plugin design, but with the fundamental difference that will receive only the Velero `Backup` Go struct without any resource list item stage, the backup was already executed. The `PostBackupAction` interface will only have `Execute()` since the plugin will be executed only once per Backup, not per item.

The Velero backup controller package will be modified if there are any PostBackupAction plugins registered; the plugins will be executed as the last step of the Backup execution `c.runBackup` function immediately after the gzip creation. If any PostBackupAction returns an error or a warning, they will be shown on the final backup status.

PostBackupAction plugins will be run in alphanumeric order based on their registered names.

The PreRestoreAction plugin API will resemble the RestoreItemAction plugin design, but with the fundamental difference that will receive only the Velero `Restore` Go struct created. It will not receive any resource list item because the restore has not yet been running at that stage. In addition, the `PreRestoreAction` interface will only have `Execute()` since the plugin will be executed only once per Restore creation, not per item.

The Velero restore controller will be modified so that if there are any PreRestoreAction plugins registered, it will be executed as the last step of restore validation. Each PreBackupAction plugin will be executed, if any PreRestoreAction returns an error, it will increment the `restore.Status.ValidationErrors` value. If there is an error on the restore object, the restore controller will never call `r.runValidatedRestore,` and in other words, the restore will not be executed.

PreRestoreAction plugins will be run in alphanumeric order based on their registered names.

The PostRestoreAction plugin API will resemble the RestoreItemAction plugin design, but with the fundamental difference that will receive only the Velero `Restore` Go struct without any resource list. At this stage, the restore was already executed. The `PostRestoreAction` interface will only have `Execute()` since the plugin will be executed only once per Restore, not per item.

The Velero restore controller package will be modified. If any PostBackupAction plugins are registered, they will be executed after `c.restorer.Restore`  independently of the restore request warning and errors. If any PreRestoreAction returns errors or warnings, they will be counted as restore errors or warnings, respectively. The PreRestoreAction will be run before the restore logs are uploaded on the object storage.

PostBackupAction plugins will be run in alphanumeric order based on their registered names.

## Detailed Design

### New types

#### PreBackupAction

The `PreBackupAction` interface is as follows:

```go
// PreBackupAction is an actor that performs an action based on a backup created and validated.
type PreBackupAction interface {
	// Execute allows the PreBackupAction to perform arbitrary logic with the backup object before its execution.
    Execute(PreBackupActionInput) error
}
```

The `PreBackupActionInput` type is defined as follows:

```go
type PreBackupActionInput struct {
	// Backup is the representation of the backup resource processed by Velero.
	Backup *api.Backup
}
```

Both `PreBackupAction` and `PreBackupActionInput` will be defined in `pkg/plugin/velero/pre_backup_action.go`.

#### PostBackupAction

The `PostBackupAction` interface is as follows:

```go
// PostBackupAction is an actor that performs an action based on a backup was executed.
type PostBackupAction interface {
	// Execute allows the PostBackupAction to perform arbitrary logic with the backup object after its execution.
    Execute(PostBackupActionInput) error
}
```

The `PostBackupActionInput` type is defined as follows:

```go
type PostBackupActionInput struct {
	// Backup is the representation of the backup resource processed by Velero.
	Backup *api.Backup
}
```

Both `PostBackupAction` and `PostBackupActionInput` will be defined in `pkg/plugin/velero/post_backup_action.go`.

#### PreRestoreAction

The `PreRestoreAction` interface is as follows:

```go
// PreRestoreAction is an actor that performs an action based on a restore was created and it is being validated.
type PreRestoreAction interface {
	// Execute allows the PreRestoreAction to perform arbitrary logic with the restore object before its execution.
    Execute(PreRestoreActionInput) error
}
```

The `PreRestoreActionInput` type is defined as follows:

```go
type PreRestoreActionInput struct {
	// Restore is the representation of the restore resource processed by Velero.
	Restore *api.Restore
}
```

Both `PreRestoreAction` and `PreRestoreActionInput` will be defined in `pkg/plugin/velero/pre_restore_action.go`.

#### PostRestoreAction

The `PostRestoreAction` interface is as follows:

```go
// PostRestoreAction is an actor that performs an action based on a restore was executed.
type PostRestoreAction interface {
	// Execute allows the PostRestorection to perform arbitrary logic with the restore object after its execution.
    Execute(PostRestoreActionInput) error
}
```

The `PostRestoreActionInput` type is defined as follows:

```go
type PostRestoreActionInput struct {
	// Restore is the representation of the restore resource processed by Velero.
	Restore *api.Restore
}
```

Both `PostRestoreAction` and `PostRestoreActionInput` will be defined in `pkg/plugin/velero/post_restore_action.go`.

### Generate protobuf definitions and client/servers

In `pkg/plugin/proto`, add:

1. Protobuf definitions will be necessary for PreBackupAction, file `PreBackupAction.proto`

```protobuf
message PreBackupActionExecuteRequest {
	...
}

message PreBackupActionExecuteResponse {
	...
}


service PreBackupAction {
	rpc Execute(PreBackupActionExecuteRequest) returns (PreBackupActionExecuteResponse)
}
```

Once these are written, then a client and server implementation can be written in `pkg/plugin/framework/pre_backup_item_action_client.go` and `pkg/plugin/framework/pre_backup_item_action_server.go`, respectively.

1. Protobuf definitions will be necessary for PostBackupAction, file `PostBackupAction.proto`

```protobuf
message PostBackupActionExecuteRequest {
	...
}

message PostBackupActionExecuteResponse {
	...
}


service PostBackupAction {
	rpc Execute(PostBackupActionExecuteRequest) returns (PostBackupActionExecuteResponse)
}
```

Once these are written, then a client and server implementation can be written in `pkg/plugin/framework/post_backup_item_action_client.go` and `pkg/plugin/framework/post_backup_item_action_server.go`, respectively.

1. Protobuf definitions will be necessary for PreRestoreAction, file `PreRestoreAction.proto`

```protobuf
message PreRestoreActionExecuteRequest {
	...
}

message PreRestoreActionExecuteResponse {
	...
}


service PreRestoreAction {
	rpc Execute(PreRestoreActionExecuteRequest) returns (PreRestoreActionExecuteResponse)
}
```

Once these are written, then a client and server implementation can be written in `pkg/plugin/framework/pre_restore_item_action_client.go` and `pkg/plugin/framework/pre_restore_item_action_server.go`, respectively.

1. Protobuf definitions will be necessary for PostRestoreAction, file `PostRestoreAction.proto`

```protobuf
message PostRestoreActionExecuteRequest {
	...
}

message PostRestoreActionExecuteResponse {
	...
}


service PostRestoreAction {
	rpc Execute(PostRestoreActionExecuteRequest) returns (PostRestoreActionExecuteResponse)
}
```

Once these are written, then a client and server implementation can be written in `pkg/plugin/framework/post_restore_item_action_client.go` and `pkg/plugin/framework/post_restore_item_action_server.go`, respectively.

### Restartable delete plugins

Similar to `RestoreItemAction` and `BackupItemAction` plugins, restartable processes will need to be implemented (with the difference there is no `AppliedTo()`).

In `pkg/plugin/clientmgmt`, add

1. `restartable_pre_backup_item_action.go`, creating the following unexported type:

```go
type restartablePreBackupAction struct {
	key kindAndName
	sharedPluginProcess RestartableProcess
	config map[string]string
}

// newRestartablePreBackupAction returns a new restartablePreBackupAction.
func newRestartablePreBackupAction(name string, sharedPluginProcess RestartableProcess) *restartablePreBackupAction {
	// ...
}

// getPreBackupAction returns the delete item action for this restartablePreBackupAction. It does *not* restart the
// plugin process.
func (r *restartablePreBackupAction) getPreBackupAction() (velero.PreBackupAction, error) {
	// ...
}

// getDelegate restarts the plugin process (if needed) and returns the prebackup action for this restartablePreBackupAction.
func (r *restartablePreBackupAction) getDelegate() (velero.PreBackupAction, error) {
	// ...
}

// Execute restarts the plugin's process if needed, then delegates the call.
func (r *restartablePreBackupAction) Execute(input *velero.PreBackupActionInput) (error) {
	// ...
}
```

1. `restartable_post_backup_item_action.go`, creating the following unexported type:

```go
type restartablePostBackupAction struct {
	key kindAndName
	sharedPluginProcess RestartableProcess
	config map[string]string
}

// newRestartablePostBackupAction returns a new restartablePostBackupAction.
func newRestartablePostBackupAction(name string, sharedPluginProcess RestartableProcess) *restartablePostBackupAction {
	// ...
}

// getPostBackupAction returns the delete item action for this restartablePostBackupAction. It does *not* restart the
// plugin process.
func (r *restartablePostBackupAction) getPostBackupAction() (velero.PostBackupAction, error) {
	// ...
}

// getDelegate restarts the plugin process (if needed) and returns the postbackup action for this restartablePostBackupAction.
func (r *restartablePostBackupAction) getDelegate() (velero.PostBackupAction, error) {
	// ...
}

// Execute restarts the plugin's process if needed, then delegates the call.
func (r *restartablePostBackupAction) Execute(input *velero.PostBackupActionInput) (error) {
	// ...
}
```

1. `restartable_pre_restore_item_action.go`, creating the following unexported type:

```go
type restartablePreRestoreAction struct {
	key kindAndName
	sharedPluginProcess RestartableProcess
	config map[string]string
}

// newRestartablePreRestoreAction returns a new restartablePreRestoreAction.
func newRestartablePreRestoreAction(name string, sharedPluginProcess RestartableProcess) *restartablePreRestoreAction {
	// ...
}

// getPreRestoreAction returns the delete item action for this restartablePreRestoreAction. It does *not* restart the
// plugin process.
func (r *restartablePreRestoreAction) getPreRestoreAction() (velero.PreRestoreAction, error) {
	// ...
}

// getDelegate restarts the plugin process (if needed) and returns the prerestore action for this restartablePreRestoreAction.
func (r *restartablePreRestoreAction) getDelegate() (velero.PreRestoreAction, error) {
	// ...
}

// Execute restarts the plugin's process if needed, then delegates the call.
func (r *restartablePreRestoreAction) Execute(input *velero.PreRestoreActionInput) (error) {
	// ...
}
```

1. `restartable_post_restore_item_action.go`, creating the following unexported type:

```go
type restartablePostRestoreAction struct {
	key kindAndName
	sharedPluginProcess RestartableProcess
	config map[string]string
}

// newRestartablePostRestoreAction returns a new restartablePostRestoreAction.
func newRestartablePostRestoreAction(name string, sharedPluginProcess RestartableProcess) *restartablePostRestoreAction {
	// ...
}

// getPostRestoreAction returns the delete item action for this restartablePostRestoreAction. It does *not* restart the
// plugin process.
func (r *restartablePostRestoreAction) getPostRestoreAction() (velero.PostRestoreAction, error) {
	// ...
}

// getDelegate restarts the plugin process (if needed) and returns the postrestore action for this restartablePostRestoreAction.
func (r *restartablePostRestoreAction) getDelegate() (velero.PostRestoreAction, error) {
	// ...
}

// Execute restarts the plugin's process if needed, then delegates the call.
func (r *restartablePostRestoreAction) Execute(input *velero.PostRestoreActionInput) (error) {
	// ...
}
```

### Plugin manager changes

Add the following methods to `pkg/plugin/clientmgmt/manager.go`'s `Manager` interface:

```go
type Manager interface {
	...
	// Get PreBackupAction returns a PreBackupAction plugin for name.
	GetPreBackupAction(name string) (PreBackupAction, error)

	// Get PreBackupActions returns the all PreBackupAction plugins.
	GetPreBackupActions() ([]PreBackupAction, error)

	// Get PostBackupAction returns a PostBackupAction plugin for name.
	GetPostBackupAction(name string) (PostBackupAction, error)

	// GetPostBackupActions returns the all PostBackupAction plugins.
	GetPostBackupActions() ([]PostBackupAction, error)

	// Get PreRestoreAction returns a PreRestoreAction plugin for name.
	GetPreRestoreAction(name string) (PreRestoreAction, error)

	// Get PreRestoreActions returns the all PreRestoreAction plugins.
	GetPreRestoreActions() ([]PreRestoreAction, error)

	// Get PostRestoreAction returns a PostRestoreAction plugin for name.
	GetPostRestoreAction(name string) (PostRestoreAction, error)

	// GetPostRestoreActions returns the all PostRestoreAction plugins.
	GetPostRestoreActions() ([]PostRestoreAction, error)

}
```

`GetPreBackupAction` and `GetPreBackupActions` will invoke the `restartablePreBackupAction` implementations.
`GetPostBackupAction` and `GetPostBackupActions` will invoke the `restartablePostBackupAction` implementations.
`GetPreRestoreAction` and `GetPreRestoreActions` will invoke the `restartablePreRestoreAction` implementations.
`GetPostRestoreAction` and `GetPostRestoreActions` will invoke the `restartablePostRestoreAction` implementations.

## Alternatives Considered

The alternative of these plugin hooks is to implement all the pre/post logic outside Velero. In this case, one would need to write an external controller similarly as Konveyor Crane does today.
We find this a viable way, but philosophically, we think better Velero provides better migrations.
We think that Velero users can benefit from more capability on Velero when they opt to write or load plugins extensions and not rely on an external component.

## Security Considerations

The plugins will only be invoked if loaded at per user's discretion.
It is recommended to check security vulnerabilities before execution.

## Compatibility

In terms of backward compatibility, this design should stay compatible with most Velero installations that are upgrading.
If not plugins are present, then the backup/restore process should proceed the same way it worked before their inclusion.

## Implementation

The implementation dependencies are roughly in the order as they are described in the [Detailed Design](#detailed-design) section.

## Open Issues
