// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta2/queue.proto

package tasks

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// State of the queue.
type Queue_State int32

const (
	// Unspecified state.
	Queue_STATE_UNSPECIFIED Queue_State = 0
	// The queue is running. Tasks can be dispatched.
	//
	// If the queue was created using Cloud Tasks and the queue has
	// had no activity (method calls or task dispatches) for 30 days,
	// the queue may take a few minutes to re-activate. Some method
	// calls may return [NOT_FOUND][google.rpc.Code.NOT_FOUND] and
	// tasks may not be dispatched for a few minutes until the queue
	// has been re-activated.
	Queue_RUNNING Queue_State = 1
	// Tasks are paused by the user. If the queue is paused then Cloud
	// Tasks will stop delivering tasks from it, but more tasks can
	// still be added to it by the user. When a pull queue is paused,
	// all [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] calls will return a
	// [FAILED_PRECONDITION][google.rpc.Code.FAILED_PRECONDITION].
	Queue_PAUSED Queue_State = 2
	// The queue is disabled.
	//
	// A queue becomes `DISABLED` when
	// [queue.yaml](https://cloud.google.com/appengine/docs/python/config/queueref) or
	// [queue.xml](https://cloud.google.com/appengine/docs/standard/java/config/queueref) is uploaded
	// which does not contain the queue. You cannot directly disable a queue.
	//
	// When a queue is disabled, tasks can still be added to a queue
	// but the tasks are not dispatched and
	// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] calls return a
	// `FAILED_PRECONDITION` error.
	//
	// To permanently delete this queue and all of its tasks, call
	// [DeleteQueue][google.cloud.tasks.v2beta2.CloudTasks.DeleteQueue].
	Queue_DISABLED Queue_State = 3
)

var Queue_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "RUNNING",
	2: "PAUSED",
	3: "DISABLED",
}

var Queue_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"RUNNING":           1,
	"PAUSED":            2,
	"DISABLED":          3,
}

func (x Queue_State) String() string {
	return proto.EnumName(Queue_State_name, int32(x))
}

func (Queue_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b86070a8ff200176, []int{0, 0}
}

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, target types, and others.
type Queue struct {
	// Caller-specified and required in [CreateQueue][google.cloud.tasks.v2beta2.CloudTasks.CreateQueue],
	// after which it becomes output only.
	//
	// The queue name.
	//
	// The queue name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the queue's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Caller-specified and required in [CreateQueue][google.cloud.tasks.v2beta2.CloudTasks.CreateQueue][],
	// after which the queue config type becomes output only, though fields within
	// the config are mutable.
	//
	// The queue's target.
	//
	// The target applies to all tasks in the queue.
	//
	// Types that are valid to be assigned to TargetType:
	//	*Queue_AppEngineHttpTarget
	//	*Queue_PullTarget
	TargetType isQueue_TargetType `protobuf_oneof:"target_type"`
	// Rate limits for task dispatches.
	//
	// [rate_limits][google.cloud.tasks.v2beta2.Queue.rate_limits] and
	// [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] are related because they both
	// control task attempts however they control how tasks are
	// attempted in different ways:
	//
	// * [rate_limits][google.cloud.tasks.v2beta2.Queue.rate_limits] controls the total rate of
	//   dispatches from a queue (i.e. all traffic dispatched from the
	//   queue, regardless of whether the dispatch is from a first
	//   attempt or a retry).
	// * [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] controls what happens to
	//   particular a task after its first attempt fails. That is,
	//   [retry_config][google.cloud.tasks.v2beta2.Queue.retry_config] controls task retries (the
	//   second attempt, third attempt, etc).
	RateLimits *RateLimits `protobuf:"bytes,5,opt,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Settings that determine the retry behavior.
	//
	// * For tasks created using Cloud Tasks: the queue-level retry settings
	//   apply to all tasks in the queue that were created using Cloud Tasks.
	//   Retry settings cannot be set on individual tasks.
	// * For tasks created using the App Engine SDK: the queue-level retry
	//   settings apply to all tasks in the queue which do not have retry settings
	//   explicitly set on the task and were created by the App Engine SDK. See
	//   [App Engine documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	RetryConfig *RetryConfig `protobuf:"bytes,6,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// Output only. The state of the queue.
	//
	// `state` can only be changed by called
	// [PauseQueue][google.cloud.tasks.v2beta2.CloudTasks.PauseQueue],
	// [ResumeQueue][google.cloud.tasks.v2beta2.CloudTasks.ResumeQueue], or uploading
	// [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	// [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] cannot be used to change `state`.
	State Queue_State `protobuf:"varint,7,opt,name=state,proto3,enum=google.cloud.tasks.v2beta2.Queue_State" json:"state,omitempty"`
	// Output only. The last time this queue was purged.
	//
	// All tasks that were [created][google.cloud.tasks.v2beta2.Task.create_time] before this time
	// were purged.
	//
	// A queue can be purged using [PurgeQueue][google.cloud.tasks.v2beta2.CloudTasks.PurgeQueue], the
	// [App Engine Task Queue SDK, or the Cloud Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	// Purge time will be truncated to the nearest microsecond. Purge
	// time will be unset if the queue has never been purged.
	PurgeTime            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Queue) Reset()         { *m = Queue{} }
func (m *Queue) String() string { return proto.CompactTextString(m) }
func (*Queue) ProtoMessage()    {}
func (*Queue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86070a8ff200176, []int{0}
}

func (m *Queue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Queue.Unmarshal(m, b)
}
func (m *Queue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Queue.Marshal(b, m, deterministic)
}
func (m *Queue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Queue.Merge(m, src)
}
func (m *Queue) XXX_Size() int {
	return xxx_messageInfo_Queue.Size(m)
}
func (m *Queue) XXX_DiscardUnknown() {
	xxx_messageInfo_Queue.DiscardUnknown(m)
}

var xxx_messageInfo_Queue proto.InternalMessageInfo

func (m *Queue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isQueue_TargetType interface {
	isQueue_TargetType()
}

type Queue_AppEngineHttpTarget struct {
	AppEngineHttpTarget *AppEngineHttpTarget `protobuf:"bytes,3,opt,name=app_engine_http_target,json=appEngineHttpTarget,proto3,oneof"`
}

type Queue_PullTarget struct {
	PullTarget *PullTarget `protobuf:"bytes,4,opt,name=pull_target,json=pullTarget,proto3,oneof"`
}

func (*Queue_AppEngineHttpTarget) isQueue_TargetType() {}

func (*Queue_PullTarget) isQueue_TargetType() {}

func (m *Queue) GetTargetType() isQueue_TargetType {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (m *Queue) GetAppEngineHttpTarget() *AppEngineHttpTarget {
	if x, ok := m.GetTargetType().(*Queue_AppEngineHttpTarget); ok {
		return x.AppEngineHttpTarget
	}
	return nil
}

func (m *Queue) GetPullTarget() *PullTarget {
	if x, ok := m.GetTargetType().(*Queue_PullTarget); ok {
		return x.PullTarget
	}
	return nil
}

func (m *Queue) GetRateLimits() *RateLimits {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *Queue) GetRetryConfig() *RetryConfig {
	if m != nil {
		return m.RetryConfig
	}
	return nil
}

func (m *Queue) GetState() Queue_State {
	if m != nil {
		return m.State
	}
	return Queue_STATE_UNSPECIFIED
}

func (m *Queue) GetPurgeTime() *timestamp.Timestamp {
	if m != nil {
		return m.PurgeTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Queue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Queue_OneofMarshaler, _Queue_OneofUnmarshaler, _Queue_OneofSizer, []interface{}{
		(*Queue_AppEngineHttpTarget)(nil),
		(*Queue_PullTarget)(nil),
	}
}

func _Queue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Queue)
	// target_type
	switch x := m.TargetType.(type) {
	case *Queue_AppEngineHttpTarget:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppEngineHttpTarget); err != nil {
			return err
		}
	case *Queue_PullTarget:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PullTarget); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Queue.TargetType has unexpected type %T", x)
	}
	return nil
}

func _Queue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Queue)
	switch tag {
	case 3: // target_type.app_engine_http_target
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AppEngineHttpTarget)
		err := b.DecodeMessage(msg)
		m.TargetType = &Queue_AppEngineHttpTarget{msg}
		return true, err
	case 4: // target_type.pull_target
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PullTarget)
		err := b.DecodeMessage(msg)
		m.TargetType = &Queue_PullTarget{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Queue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Queue)
	// target_type
	switch x := m.TargetType.(type) {
	case *Queue_AppEngineHttpTarget:
		s := proto.Size(x.AppEngineHttpTarget)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Queue_PullTarget:
		s := proto.Size(x.PullTarget)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Rate limits.
//
// This message determines the maximum rate that tasks can be dispatched by a
// queue, regardless of whether the dispatch is a first task attempt or a retry.
//
// Note: The debugging command, [RunTask][google.cloud.tasks.v2beta2.CloudTasks.RunTask], will run a task
// even if the queue has reached its [RateLimits][google.cloud.tasks.v2beta2.RateLimits].
type RateLimits struct {
	// The maximum rate at which tasks are dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// * For [App Engine queues][google.cloud.tasks.v2beta2.AppEngineHttpTarget], the maximum allowed value
	//   is 500.
	// * This field is output only   for [pull queues][google.cloud.tasks.v2beta2.PullTarget]. In addition to the
	//   `max_tasks_dispatched_per_second` limit, a maximum of 10 QPS of
	//   [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] requests are allowed per pull queue.
	//
	//
	// This field has the same meaning as
	// [rate in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
	MaxTasksDispatchedPerSecond float64 `protobuf:"fixed64,1,opt,name=max_tasks_dispatched_per_second,json=maxTasksDispatchedPerSecond,proto3" json:"max_tasks_dispatched_per_second,omitempty"`
	// Output only. The max burst size.
	//
	// Max burst size limits how fast tasks in queue are processed when
	// many tasks are in the queue and the rate is high. This field
	// allows the queue to have a high rate so processing starts shortly
	// after a task is enqueued, but still limits resource usage when
	// many tasks are enqueued in a short period of time.
	//
	// The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	// Algorithm is used to control the rate of task dispatches. Each
	// queue has a token bucket that holds tokens, up to the maximum
	// specified by `max_burst_size`. Each time a task is dispatched, a
	// token is removed from the bucket. Tasks will be dispatched until
	// the queue's bucket runs out of tokens. The bucket will be
	// continuously refilled with new tokens based on
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second].
	//
	// Cloud Tasks will pick the value of `max_burst_size` based on the
	// value of
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second].
	//
	// For App Engine queues that were created or updated using
	// `queue.yaml/xml`, `max_burst_size` is equal to
	// [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size).
	// Since `max_burst_size` is output only, if
	// [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue] is called on a queue
	// created by `queue.yaml/xml`, `max_burst_size` will be reset based
	// on the value of
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second],
	// regardless of whether
	// [max_tasks_dispatched_per_second][google.cloud.tasks.v2beta2.RateLimits.max_tasks_dispatched_per_second]
	// is updated.
	//
	MaxBurstSize int32 `protobuf:"varint,2,opt,name=max_burst_size,json=maxBurstSize,proto3" json:"max_burst_size,omitempty"`
	// The maximum number of concurrent tasks that Cloud Tasks allows
	// to be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	// The maximum allowed value is 5,000.
	//
	// This field is output only for
	// [pull queues][google.cloud.tasks.v2beta2.PullTarget] and always -1, which indicates no limit. No other
	// queue types can have `max_concurrent_tasks` set to -1.
	//
	//
	// This field has the same meaning as
	// [max_concurrent_requests in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
	MaxConcurrentTasks   int32    `protobuf:"varint,3,opt,name=max_concurrent_tasks,json=maxConcurrentTasks,proto3" json:"max_concurrent_tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimits) Reset()         { *m = RateLimits{} }
func (m *RateLimits) String() string { return proto.CompactTextString(m) }
func (*RateLimits) ProtoMessage()    {}
func (*RateLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86070a8ff200176, []int{1}
}

func (m *RateLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimits.Unmarshal(m, b)
}
func (m *RateLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimits.Marshal(b, m, deterministic)
}
func (m *RateLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimits.Merge(m, src)
}
func (m *RateLimits) XXX_Size() int {
	return xxx_messageInfo_RateLimits.Size(m)
}
func (m *RateLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimits.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimits proto.InternalMessageInfo

func (m *RateLimits) GetMaxTasksDispatchedPerSecond() float64 {
	if m != nil {
		return m.MaxTasksDispatchedPerSecond
	}
	return 0
}

func (m *RateLimits) GetMaxBurstSize() int32 {
	if m != nil {
		return m.MaxBurstSize
	}
	return 0
}

func (m *RateLimits) GetMaxConcurrentTasks() int32 {
	if m != nil {
		return m.MaxConcurrentTasks
	}
	return 0
}

// Retry config.
//
// These settings determine how a failed task attempt is retried.
type RetryConfig struct {
	// Number of attempts per task.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	//
	//
	// This field has the same meaning as
	// [task_retry_limit in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	//
	// Types that are valid to be assigned to NumAttempts:
	//	*RetryConfig_MaxAttempts
	//	*RetryConfig_UnlimitedAttempts
	NumAttempts isRetryConfig_NumAttempts `protobuf_oneof:"num_attempts"`
	// If positive, `max_retry_duration` specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once `max_retry_duration` time has passed *and* the
	// task has been attempted [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts]
	// times, no further attempts will be made and the task will be
	// deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// `max_retry_duration` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [task_age_limit in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxRetryDuration *duration.Duration `protobuf:"bytes,3,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2beta2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2beta2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// `min_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [min_backoff_seconds in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MinBackoff *duration.Duration `protobuf:"bytes,4,opt,name=min_backoff,json=minBackoff,proto3" json:"min_backoff,omitempty"`
	// A task will be [scheduled][google.cloud.tasks.v2beta2.Task.schedule_time] for retry between
	// [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] and
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] duration after it fails,
	// if the queue's [RetryConfig][google.cloud.tasks.v2beta2.RetryConfig] specifies that the task should be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// `max_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [max_backoff_seconds in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxBackoff *duration.Duration `protobuf:"bytes,5,opt,name=max_backoff,json=maxBackoff,proto3" json:"max_backoff,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A task's retry interval starts at
	// [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff], then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] up to
	// [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts] times.
	//
	// For example, if [min_backoff][google.cloud.tasks.v2beta2.RetryConfig.min_backoff] is 10s,
	// [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] is 300s, and
	// `max_doublings` is 3, then the a task will first be retried in
	// 10s. The retry interval will double three times, and then
	// increase linearly by 2^3 * 10s.  Finally, the task will retry at
	// intervals of [max_backoff][google.cloud.tasks.v2beta2.RetryConfig.max_backoff] until the
	// task has been attempted [max_attempts][google.cloud.tasks.v2beta2.RetryConfig.max_attempts]
	// times. Thus, the requests will retry at 10s, 20s, 40s, 80s, 160s,
	// 240s, 300s, 300s, ....
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field is output only for [pull queues][google.cloud.tasks.v2beta2.PullTarget].
	//
	//
	// This field has the same meaning as
	// [max_doublings in queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxDoublings         int32    `protobuf:"varint,6,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryConfig) Reset()         { *m = RetryConfig{} }
func (m *RetryConfig) String() string { return proto.CompactTextString(m) }
func (*RetryConfig) ProtoMessage()    {}
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b86070a8ff200176, []int{2}
}

func (m *RetryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryConfig.Unmarshal(m, b)
}
func (m *RetryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryConfig.Marshal(b, m, deterministic)
}
func (m *RetryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryConfig.Merge(m, src)
}
func (m *RetryConfig) XXX_Size() int {
	return xxx_messageInfo_RetryConfig.Size(m)
}
func (m *RetryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RetryConfig proto.InternalMessageInfo

type isRetryConfig_NumAttempts interface {
	isRetryConfig_NumAttempts()
}

type RetryConfig_MaxAttempts struct {
	MaxAttempts int32 `protobuf:"varint,1,opt,name=max_attempts,json=maxAttempts,proto3,oneof"`
}

type RetryConfig_UnlimitedAttempts struct {
	UnlimitedAttempts bool `protobuf:"varint,2,opt,name=unlimited_attempts,json=unlimitedAttempts,proto3,oneof"`
}

func (*RetryConfig_MaxAttempts) isRetryConfig_NumAttempts() {}

func (*RetryConfig_UnlimitedAttempts) isRetryConfig_NumAttempts() {}

func (m *RetryConfig) GetNumAttempts() isRetryConfig_NumAttempts {
	if m != nil {
		return m.NumAttempts
	}
	return nil
}

func (m *RetryConfig) GetMaxAttempts() int32 {
	if x, ok := m.GetNumAttempts().(*RetryConfig_MaxAttempts); ok {
		return x.MaxAttempts
	}
	return 0
}

func (m *RetryConfig) GetUnlimitedAttempts() bool {
	if x, ok := m.GetNumAttempts().(*RetryConfig_UnlimitedAttempts); ok {
		return x.UnlimitedAttempts
	}
	return false
}

func (m *RetryConfig) GetMaxRetryDuration() *duration.Duration {
	if m != nil {
		return m.MaxRetryDuration
	}
	return nil
}

func (m *RetryConfig) GetMinBackoff() *duration.Duration {
	if m != nil {
		return m.MinBackoff
	}
	return nil
}

func (m *RetryConfig) GetMaxBackoff() *duration.Duration {
	if m != nil {
		return m.MaxBackoff
	}
	return nil
}

func (m *RetryConfig) GetMaxDoublings() int32 {
	if m != nil {
		return m.MaxDoublings
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RetryConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RetryConfig_OneofMarshaler, _RetryConfig_OneofUnmarshaler, _RetryConfig_OneofSizer, []interface{}{
		(*RetryConfig_MaxAttempts)(nil),
		(*RetryConfig_UnlimitedAttempts)(nil),
	}
}

func _RetryConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RetryConfig)
	// num_attempts
	switch x := m.NumAttempts.(type) {
	case *RetryConfig_MaxAttempts:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MaxAttempts))
	case *RetryConfig_UnlimitedAttempts:
		t := uint64(0)
		if x.UnlimitedAttempts {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("RetryConfig.NumAttempts has unexpected type %T", x)
	}
	return nil
}

func _RetryConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RetryConfig)
	switch tag {
	case 1: // num_attempts.max_attempts
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.NumAttempts = &RetryConfig_MaxAttempts{int32(x)}
		return true, err
	case 2: // num_attempts.unlimited_attempts
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.NumAttempts = &RetryConfig_UnlimitedAttempts{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _RetryConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RetryConfig)
	// num_attempts
	switch x := m.NumAttempts.(type) {
	case *RetryConfig_MaxAttempts:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.MaxAttempts))
	case *RetryConfig_UnlimitedAttempts:
		n += 1 // tag and wire
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2beta2.Queue_State", Queue_State_name, Queue_State_value)
	proto.RegisterType((*Queue)(nil), "google.cloud.tasks.v2beta2.Queue")
	proto.RegisterType((*RateLimits)(nil), "google.cloud.tasks.v2beta2.RateLimits")
	proto.RegisterType((*RetryConfig)(nil), "google.cloud.tasks.v2beta2.RetryConfig")
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta2/queue.proto", fileDescriptor_b86070a8ff200176)
}

var fileDescriptor_b86070a8ff200176 = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0x86, 0xad, 0x24, 0xca, 0x65, 0xe4, 0x04, 0x0e, 0xcf, 0x05, 0x3e, 0x3e, 0x45, 0x62, 0x38,
	0x45, 0xe2, 0x95, 0x54, 0xa4, 0xab, 0xb6, 0x28, 0x0a, 0x3b, 0x72, 0x63, 0x17, 0x81, 0xe1, 0xca,
	0xce, 0xa6, 0x1b, 0x82, 0x96, 0x69, 0x45, 0x88, 0x44, 0xb2, 0x12, 0x55, 0x38, 0x79, 0xa4, 0x6e,
	0xfb, 0x34, 0x7d, 0x9b, 0x82, 0xd4, 0x25, 0x45, 0xd3, 0x3a, 0x3b, 0x73, 0xe6, 0xfb, 0xff, 0x19,
	0x71, 0x86, 0x86, 0xd3, 0x80, 0xf3, 0x20, 0xa2, 0x8e, 0x1f, 0xf1, 0x6c, 0xe1, 0x48, 0x92, 0xde,
	0xa6, 0xce, 0x97, 0xf3, 0x39, 0x95, 0xe4, 0xdc, 0xf9, 0x9c, 0xd1, 0x8c, 0xda, 0x22, 0xe1, 0x92,
	0xa3, 0x56, 0xce, 0xd9, 0x9a, 0xb3, 0x35, 0x67, 0x17, 0x5c, 0xeb, 0x59, 0xe1, 0x41, 0x44, 0xe8,
	0x10, 0xc6, 0xb8, 0x24, 0x32, 0xe4, 0x2c, 0xcd, 0x95, 0xad, 0xb3, 0x35, 0x15, 0x24, 0x49, 0x02,
	0x2a, 0x0b, 0xf0, 0xa8, 0x00, 0xf5, 0x69, 0x9e, 0x2d, 0x9d, 0x45, 0x96, 0x68, 0xa7, 0x22, 0x7f,
	0xfc, 0x6b, 0x5e, 0x86, 0x31, 0x4d, 0x25, 0x89, 0x45, 0x0e, 0x74, 0xbe, 0x6d, 0x81, 0xf9, 0x51,
	0xf5, 0x8c, 0x10, 0x6c, 0x31, 0x12, 0xd3, 0xa6, 0xd1, 0x36, 0xba, 0x7b, 0x9e, 0xfe, 0x8d, 0x96,
	0xf0, 0x2f, 0x11, 0x02, 0x53, 0x16, 0x84, 0x8c, 0xe2, 0x1b, 0x29, 0x05, 0xce, 0xcb, 0x37, 0x37,
	0xdb, 0x46, 0xd7, 0x3a, 0x77, 0xec, 0x3f, 0x7f, 0xa2, 0xdd, 0x13, 0x62, 0xa0, 0x85, 0x43, 0x29,
	0xc5, 0x4c, 0xcb, 0x86, 0x35, 0xef, 0x2f, 0xf2, 0x38, 0x8c, 0x46, 0x60, 0x89, 0x2c, 0x8a, 0x4a,
	0xf3, 0x2d, 0x6d, 0x7e, 0xba, 0xce, 0x7c, 0x92, 0x45, 0x51, 0xe5, 0x09, 0xa2, 0x3a, 0xa1, 0x4b,
	0xb0, 0x12, 0x22, 0x29, 0x8e, 0xc2, 0x38, 0x94, 0x69, 0xd3, 0x7c, 0xda, 0xca, 0x23, 0x92, 0x5e,
	0x69, 0xda, 0x83, 0xa4, 0xfa, 0x8d, 0x3e, 0x40, 0x3d, 0xa1, 0x32, 0xb9, 0xc3, 0x3e, 0x67, 0xcb,
	0x30, 0x68, 0x6e, 0x6b, 0xa7, 0xb3, 0xb5, 0x4e, 0x8a, 0xbf, 0xd0, 0xb8, 0x67, 0x25, 0x0f, 0x07,
	0xf4, 0x16, 0xcc, 0x54, 0x12, 0x49, 0x9b, 0x3b, 0x6d, 0xa3, 0x7b, 0xb0, 0xde, 0x44, 0x4f, 0xc3,
	0x9e, 0x2a, 0xdc, 0xcb, 0x55, 0xe8, 0x15, 0x80, 0xc8, 0x92, 0x80, 0x62, 0x35, 0xbd, 0xe6, 0xae,
	0x6e, 0xa4, 0x55, 0x7a, 0x94, 0xa3, 0xb5, 0x67, 0xe5, 0x68, 0xbd, 0x3d, 0x4d, 0xab, 0x73, 0x67,
	0x00, 0xa6, 0xb6, 0x42, 0xff, 0xc0, 0xe1, 0x74, 0xd6, 0x9b, 0x0d, 0xf0, 0xf5, 0x78, 0x3a, 0x19,
	0x5c, 0x8c, 0xde, 0x8f, 0x06, 0x6e, 0xa3, 0x86, 0x2c, 0xd8, 0xf1, 0xae, 0xc7, 0xe3, 0xd1, 0xf8,
	0xb2, 0x61, 0x20, 0x80, 0xed, 0x49, 0xef, 0x7a, 0x3a, 0x70, 0x1b, 0x1b, 0xa8, 0x0e, 0xbb, 0xee,
	0x68, 0xda, 0xeb, 0x5f, 0x0d, 0xdc, 0xc6, 0x66, 0x7f, 0x1f, 0xac, 0x7c, 0x36, 0x58, 0xde, 0x09,
	0xda, 0xf9, 0x6a, 0x00, 0x3c, 0x5c, 0x1b, 0x72, 0xe1, 0x38, 0x26, 0x2b, 0xac, 0xbf, 0x03, 0x2f,
	0xc2, 0x54, 0x10, 0xe9, 0xdf, 0xd0, 0x05, 0x16, 0x34, 0xc1, 0x29, 0xf5, 0x39, 0x5b, 0xe8, 0xad,
	0x32, 0xbc, 0xff, 0x63, 0xb2, 0x9a, 0x29, 0xca, 0xad, 0xa0, 0x09, 0x4d, 0xa6, 0x1a, 0x41, 0xcf,
	0xe1, 0x40, 0xb9, 0xcc, 0xb3, 0x24, 0x95, 0x38, 0x0d, 0xef, 0x69, 0x73, 0xa3, 0x6d, 0x74, 0x4d,
	0xaf, 0x1e, 0x93, 0x55, 0x5f, 0x05, 0xa7, 0xe1, 0x3d, 0x45, 0x2f, 0xe0, 0x6f, 0x45, 0xf9, 0x9c,
	0xf9, 0x59, 0x92, 0x50, 0x26, 0xf3, 0xb2, 0x7a, 0x21, 0x4d, 0x0f, 0xc5, 0x64, 0x75, 0x51, 0xa5,
	0x74, 0xa9, 0xce, 0xf7, 0x0d, 0xb0, 0x7e, 0x9a, 0x0c, 0x3a, 0x01, 0xe5, 0x88, 0x89, 0x94, 0x34,
	0x16, 0x32, 0xd5, 0xad, 0x99, 0xc3, 0x9a, 0x67, 0xc5, 0x64, 0xd5, 0x2b, 0x82, 0xc8, 0x01, 0x94,
	0x31, 0xbd, 0x43, 0x74, 0xf1, 0x80, 0xaa, 0x86, 0x76, 0x87, 0x35, 0xef, 0xb0, 0xca, 0x55, 0x82,
	0x4b, 0x50, 0xb5, 0x71, 0xbe, 0x32, 0xe5, 0x2b, 0x2c, 0x9e, 0xc9, 0x7f, 0x8f, 0x66, 0xe5, 0x16,
	0x80, 0xd7, 0x88, 0xc9, 0x4a, 0x37, 0x57, 0x46, 0xd0, 0x6b, 0xb0, 0xe2, 0x90, 0xe1, 0x39, 0xf1,
	0x6f, 0xf9, 0x72, 0x59, 0xbc, 0x85, 0x35, 0x0e, 0x10, 0x87, 0xac, 0x9f, 0xc3, 0x5a, 0xab, 0xae,
	0xb0, 0xd0, 0x9a, 0x4f, 0x6b, 0xc9, 0xaa, 0xd4, 0x9e, 0xc0, 0xbe, 0xd2, 0x2e, 0x78, 0x36, 0x8f,
	0x42, 0x16, 0xa4, 0x7a, 0xe1, 0xf3, 0xdb, 0x77, 0xcb, 0x58, 0xff, 0x00, 0xea, 0x2c, 0x8b, 0xab,
	0x0b, 0xe9, 0x73, 0x38, 0xf2, 0x79, 0xbc, 0x66, 0x9d, 0xfb, 0xa0, 0xf7, 0x79, 0xa2, 0x4a, 0x4f,
	0x8c, 0x4f, 0xef, 0x0a, 0x32, 0xe0, 0x11, 0x61, 0x81, 0xcd, 0x93, 0xc0, 0x09, 0x28, 0xd3, 0x8d,
	0x39, 0x79, 0x8a, 0x88, 0x30, 0xfd, 0xdd, 0xff, 0xde, 0x1b, 0x7d, 0x9a, 0x6f, 0x6b, 0xf6, 0xe5,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd1, 0xaa, 0xde, 0x84, 0x05, 0x00, 0x00,
}
