// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/v1/job/job.proto

package job

import (
	types "github.com/frantjc/sequence/api/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job      *types.Job      `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Workflow *types.Workflow `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_job_proto_rawDescGZIP(), []int{0}
}

func (x *RunJobRequest) GetJob() *types.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *RunJobRequest) GetWorkflow() *types.Workflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

var File_api_v1_job_job_proto protoreflect.FileDescriptor

var file_api_v1_job_job_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6a, 0x6f, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6a, 0x6f, 0x62, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0d, 0x52,
	0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03,
	0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03,
	0x6a, 0x6f, 0x62, 0x12, 0x34, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x32, 0x46, 0x0a, 0x03, 0x4a, 0x6f, 0x62,
	0x12, 0x3f, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x52, 0x75, 0x6e,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x30,
	0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x72, 0x61, 0x6e, 0x74, 0x6a, 0x63, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_v1_job_job_proto_rawDescOnce sync.Once
	file_api_v1_job_job_proto_rawDescData = file_api_v1_job_job_proto_rawDesc
)

func file_api_v1_job_job_proto_rawDescGZIP() []byte {
	file_api_v1_job_job_proto_rawDescOnce.Do(func() {
		file_api_v1_job_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_job_job_proto_rawDescData)
	})
	return file_api_v1_job_job_proto_rawDescData
}

var file_api_v1_job_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_job_job_proto_goTypes = []interface{}{
	(*RunJobRequest)(nil),  // 0: sequence.v1.job.RunJobRequest
	(*types.Job)(nil),      // 1: sequence.types.Job
	(*types.Workflow)(nil), // 2: sequence.types.Workflow
	(*types.Log)(nil),      // 3: sequence.types.Log
}
var file_api_v1_job_job_proto_depIdxs = []int32{
	1, // 0: sequence.v1.job.RunJobRequest.job:type_name -> sequence.types.Job
	2, // 1: sequence.v1.job.RunJobRequest.workflow:type_name -> sequence.types.Workflow
	0, // 2: sequence.v1.job.Job.RunJob:input_type -> sequence.v1.job.RunJobRequest
	3, // 3: sequence.v1.job.Job.RunJob:output_type -> sequence.types.Log
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_job_job_proto_init() }
func file_api_v1_job_job_proto_init() {
	if File_api_v1_job_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_job_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_job_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_job_job_proto_goTypes,
		DependencyIndexes: file_api_v1_job_job_proto_depIdxs,
		MessageInfos:      file_api_v1_job_job_proto_msgTypes,
	}.Build()
	File_api_v1_job_job_proto = out.File
	file_api_v1_job_job_proto_rawDesc = nil
	file_api_v1_job_job_proto_goTypes = nil
	file_api_v1_job_job_proto_depIdxs = nil
}
