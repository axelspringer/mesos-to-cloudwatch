package metrics

type Mesos struct {
	MasterCpusRevocablePercent                             int     `json:"master/cpus_revocable_percent"`
	MasterValidStatusUpdateAcknowledgements                int     `json:"master/valid_status_update_acknowledgements"`
	MasterMessagesAuthenticate                             int     `json:"master/messages_authenticate"`
	RegistrarStateStoreMsP50                               float64 `json:"registrar/state_store_ms/p50"`
	MasterTasksKilled                                      int     `json:"master/tasks_killed"`
	MasterSlavesActive                                     int     `json:"master/slaves_active"`
	MasterMessagesFrameworkToExecutor                      int     `json:"master/messages_framework_to_executor"`
	MasterMessagesKillTask                                 int     `json:"master/messages_kill_task"`
	MasterMessagesDeclineOffers                            int     `json:"master/messages_decline_offers"`
	MasterValidExecutorToFrameworkMessages                 int     `json:"master/valid_executor_to_framework_messages"`
	SystemLoad5Min                                         float64 `json:"system/load_5min"`
	MasterSlaveReregistrations                             int     `json:"master/slave_reregistrations"`
	MasterMemPercent                                       float64 `json:"master/mem_percent"`
	MasterInvalidExecutorToFrameworkMessages               int     `json:"master/invalid_executor_to_framework_messages"`
	MasterSlaveRegistrations                               int     `json:"master/slave_registrations"`
	MasterSlavesConnected                                  int     `json:"master/slaves_connected"`
	RegistrarStateStoreMsP99                               float64 `json:"registrar/state_store_ms/p99"`
	MasterMessagesReregisterSlave                          int     `json:"master/messages_reregister_slave"`
	MasterMessagesExitedExecutor                           int     `json:"master/messages_exited_executor"`
	MasterMemRevocableUsed                                 int     `json:"master/mem_revocable_used"`
	MasterMessagesExecutorToFramework                      int     `json:"master/messages_executor_to_framework"`
	MasterEventQueueHTTPRequests                           int     `json:"master/event_queue_http_requests"`
	SystemLoad15Min                                        float64 `json:"system/load_15min"`
	RegistrarStateStoreMsP999                              float64 `json:"registrar/state_store_ms/p999"`
	AllocatorEventQueueDispatches                          int     `json:"allocator/event_queue_dispatches"`
	RegistrarStateFetchMs                                  float64 `json:"registrar/state_fetch_ms"`
	MasterMessagesUnregisterSlave                          int     `json:"master/messages_unregister_slave"`
	MasterMessagesRegisterFramework                        int     `json:"master/messages_register_framework"`
	MasterMessagesStatusUpdateAcknowledgement              int     `json:"master/messages_status_update_acknowledgement"`
	MasterSlaveRemovals                                    int     `json:"master/slave_removals"`
	MasterMessagesDeactivateFramework                      int     `json:"master/messages_deactivate_framework"`
	RegistrarStateStoreMsP90                               float64 `json:"registrar/state_store_ms/p90"`
	MasterMessagesLaunchTasks                              int     `json:"master/messages_launch_tasks"`
	MasterTasksError                                       int     `json:"master/tasks_error"`
	MasterMessagesResourceRequest                          int     `json:"master/messages_resource_request"`
	MasterTasksLost                                        int     `json:"master/tasks_lost"`
	MasterDiskRevocableTotal                               int     `json:"master/disk_revocable_total"`
	MasterInvalidFrameworkToExecutorMessages               int     `json:"master/invalid_framework_to_executor_messages"`
	MasterEventQueueDispatches                             int     `json:"master/event_queue_dispatches"`
	MasterOutstandingOffers                                int     `json:"master/outstanding_offers"`
	MasterCpusRevocableTotal                               int     `json:"master/cpus_revocable_total"`
	MasterUptimeSecs                                       float64 `json:"master/uptime_secs"`
	MasterTaskLostSourceSlaveReasonExecutorTerminated      int     `json:"master/task_lost/source_slave/reason_executor_terminated"`
	SystemCpusTotal                                        int     `json:"system/cpus_total"`
	MasterSlaveShutdownsCompleted                          int     `json:"master/slave_shutdowns_completed"`
	MasterTasksStaging                                     int     `json:"master/tasks_staging"`
	SystemMemFreeBytes                                     int64   `json:"system/mem_free_bytes"`
	MasterSlavesDisconnected                               int     `json:"master/slaves_disconnected"`
	MasterInvalidStatusUpdateAcknowledgements              int     `json:"master/invalid_status_update_acknowledgements"`
	MasterTasksFinished                                    int     `json:"master/tasks_finished"`
	MasterTaskFailedSourceSlaveReasonCommandExecutorFailed int     `json:"master/task_failed/source_slave/reason_command_executor_failed"`
	MasterTaskLostSourceSlaveReasonReconciliation          int     `json:"master/task_lost/source_slave/reason_reconciliation"`
	MasterSlavesInactive                                   int     `json:"master/slaves_inactive"`
	MasterMemRevocablePercent                              int     `json:"master/mem_revocable_percent"`
	MasterDiskTotal                                        int     `json:"master/disk_total"`
	MasterFrameworksDisconnected                           int     `json:"master/frameworks_disconnected"`
	MasterFrameworksInactive                               int     `json:"master/frameworks_inactive"`
	MasterMessagesReregisterFramework                      int     `json:"master/messages_reregister_framework"`
	MasterCpusTotal                                        int     `json:"master/cpus_total"`
	MasterSlaveShutdownsCanceled                           int     `json:"master/slave_shutdowns_canceled"`
	RegistrarRegistrySizeBytes                             int     `json:"registrar/registry_size_bytes"`
	RegistrarQueuedOperations                              int     `json:"registrar/queued_operations"`
	SystemLoad1Min                                         float64 `json:"system/load_1min"`
	MasterMessagesUnregisterFramework                      int     `json:"master/messages_unregister_framework"`
	MasterMessagesRegisterSlave                            int     `json:"master/messages_register_slave"`
	MasterMessagesStatusUpdate                             int     `json:"master/messages_status_update"`
	MasterMemTotal                                         int     `json:"master/mem_total"`
	MasterTasksRunning                                     int     `json:"master/tasks_running"`
	MasterTaskFailedSourceSlaveReasonMemoryLimit           int     `json:"master/task_failed/source_slave/reason_memory_limit"`
	MasterSlaveRemovalsReasonUnregistered                  int     `json:"master/slave_removals/reason_unregistered"`
	MasterFrameworksActive                                 int     `json:"master/frameworks_active"`
	RegistrarStateStoreMsP9999                             float64 `json:"registrar/state_store_ms/p9999"`
	SystemMemTotalBytes                                    int64   `json:"system/mem_total_bytes"`
	MasterCpusRevocableUsed                                int     `json:"master/cpus_revocable_used"`
	MasterSlaveRemovalsReasonRegistered                    int     `json:"master/slave_removals/reason_registered"`
	RegistrarStateStoreMsP95                               float64 `json:"registrar/state_store_ms/p95"`
	MasterDroppedMessages                                  int     `json:"master/dropped_messages"`
	MasterCpusPercent                                      float64 `json:"master/cpus_percent"`
	MasterDiskRevocableUsed                                int     `json:"master/disk_revocable_used"`
	MasterMemRevocableTotal                                int     `json:"master/mem_revocable_total"`
	MasterMemUsed                                          int     `json:"master/mem_used"`
	MasterElected                                          int     `json:"master/elected"`
	MasterValidStatusUpdates                               int     `json:"master/valid_status_updates"`
	MasterSlaveRemovalsReasonUnhealthy                     int     `json:"master/slave_removals/reason_unhealthy"`
	RegistrarStateStoreMsCount                             int     `json:"registrar/state_store_ms/count"`
	MasterRecoverySlaveRemovals                            int     `json:"master/recovery_slave_removals"`
	MasterDiskUsed                                         int     `json:"master/disk_used"`
	MasterCpusUsed                                         float64 `json:"master/cpus_used"`
	RegistrarStateStoreMs                                  float64 `json:"registrar/state_store_ms"`
	MasterEventQueueMessages                               int     `json:"master/event_queue_messages"`
	MasterTasksFailed                                      int     `json:"master/tasks_failed"`
	MasterDiskPercent                                      float64 `json:"master/disk_percent"`
	MasterDiskRevocablePercent                             int     `json:"master/disk_revocable_percent"`
	MasterMessagesUpdateSlave                              int     `json:"master/messages_update_slave"`
	MasterFrameworksConnected                              int     `json:"master/frameworks_connected"`
	MasterTaskKilledSourceSlaveReasonExecutorUnregistered  int     `json:"master/task_killed/source_slave/reason_executor_unregistered"`
	RegistrarStateStoreMsMax                               float64 `json:"registrar/state_store_ms/max"`
	MasterMessagesReconcileTasks                           int     `json:"master/messages_reconcile_tasks"`
	MasterTasksStarting                                    int     `json:"master/tasks_starting"`
	RegistrarStateStoreMsMin                               float64 `json:"registrar/state_store_ms/min"`
	MasterTaskLostSourceMasterReasonSlaveRemoved           int     `json:"master/task_lost/source_master/reason_slave_removed"`
	MasterMessagesReviveOffers                             int     `json:"master/messages_revive_offers"`
	MasterSlaveShutdownsScheduled                          int     `json:"master/slave_shutdowns_scheduled"`
	MasterValidFrameworkToExecutorMessages                 int     `json:"master/valid_framework_to_executor_messages"`
	MasterInvalidStatusUpdates                             int     `json:"master/invalid_status_updates"`
}
