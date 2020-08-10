// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

// model task redis pb

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Task struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 TaskType `protobuf:"varint,4,opt,name=type,proto3,enum=pb.TaskType" json:"type,omitempty"`
	Times                int32    `protobuf:"varint,5,opt,name=times,proto3" json:"times,omitempty"`
	RetryTimes           int32    `protobuf:"varint,6,opt,name=retry_times,json=retryTimes,proto3" json:"retry_times,omitempty"`
	ExecTime             int64    `protobuf:"varint,7,opt,name=exec_time,json=execTime,proto3" json:"exec_time,omitempty"`
	NextExecTime         int64    `protobuf:"varint,8,opt,name=next_exec_time,json=nextExecTime,proto3" json:"next_exec_time,omitempty"`
	Schema               string   `protobuf:"bytes,9,opt,name=schema,proto3" json:"schema,omitempty"`
	Address              string   `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Path                 string   `protobuf:"bytes,11,opt,name=path,proto3" json:"path,omitempty"`
	CreatedAt            int64    `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetType() TaskType {
	if m != nil {
		return m.Type
	}
	return TaskType_Ignore
}

func (m *Task) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

func (m *Task) GetRetryTimes() int32 {
	if m != nil {
		return m.RetryTimes
	}
	return 0
}

func (m *Task) GetExecTime() int64 {
	if m != nil {
		return m.ExecTime
	}
	return 0
}

func (m *Task) GetNextExecTime() int64 {
	if m != nil {
		return m.NextExecTime
	}
	return 0
}

func (m *Task) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Task) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Task) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Task) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Task) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Task)(nil), "pb.Task")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcf, 0x6a, 0x84, 0x30,
	0x10, 0xc6, 0x89, 0xff, 0x76, 0x9d, 0x75, 0xa5, 0x84, 0x52, 0x42, 0x4b, 0x69, 0x28, 0x3d, 0x78,
	0xf2, 0xd0, 0x3e, 0xc1, 0x1e, 0xfa, 0x02, 0xe2, 0x5d, 0xa2, 0x99, 0xb2, 0xb2, 0xe8, 0x06, 0x93,
	0x85, 0xf5, 0x5d, 0xfb, 0x30, 0x25, 0xa3, 0xd2, 0xdb, 0x7c, 0xbf, 0xdf, 0x47, 0x92, 0x09, 0x80,
	0x53, 0xf6, 0x52, 0x9a, 0xe9, 0xea, 0xae, 0x3c, 0x30, 0xed, 0x73, 0xa6, 0xf1, 0xa7, 0x1f, 0x71,
	0x21, 0xef, 0xbf, 0x01, 0x44, 0xb5, 0xb2, 0x17, 0x9e, 0x43, 0xd0, 0x6b, 0xc1, 0x24, 0x2b, 0x8e,
	0x55, 0xd0, 0x6b, 0xfe, 0x00, 0xe1, 0xad, 0xd7, 0x22, 0x90, 0xac, 0x48, 0x2b, 0x3f, 0x72, 0x0e,
	0xd1, 0xa8, 0x06, 0x14, 0x21, 0x21, 0x9a, 0xb9, 0x84, 0xc8, 0xcd, 0x06, 0x45, 0x24, 0x59, 0x91,
	0x7f, 0x66, 0xa5, 0x69, 0x4b, 0x7f, 0x5a, 0x3d, 0x1b, 0xac, 0xc8, 0xf0, 0x47, 0x88, 0x5d, 0x3f,
	0xa0, 0x15, 0xb1, 0x64, 0x45, 0x5c, 0x2d, 0x81, 0xbf, 0xc1, 0x61, 0x42, 0x37, 0xcd, 0xcd, 0xe2,
	0x12, 0x72, 0x40, 0xa8, 0xa6, 0xc2, 0x0b, 0xa4, 0x78, 0xc7, 0x8e, 0xbc, 0xd8, 0x49, 0x56, 0x84,
	0xd5, 0xde, 0x03, 0x6f, 0xf9, 0x07, 0xe4, 0x23, 0xde, 0x5d, 0xf3, 0xdf, 0xd8, 0x53, 0x23, 0xf3,
	0xf4, 0x7b, 0x6b, 0x3d, 0x41, 0x62, 0xbb, 0x33, 0x0e, 0x4a, 0xa4, 0xf4, 0xe2, 0x35, 0x71, 0x01,
	0x3b, 0xa5, 0xf5, 0x84, 0xd6, 0x0a, 0x20, 0xb1, 0x45, 0xbf, 0xa1, 0x51, 0xee, 0x2c, 0x0e, 0xcb,
	0x86, 0x7e, 0xe6, 0xaf, 0x00, 0xdd, 0x84, 0xca, 0xa1, 0x6e, 0x94, 0x13, 0x19, 0xdd, 0x93, 0xae,
	0xe4, 0xe4, 0xbc, 0xbe, 0x19, 0xbd, 0xe9, 0xe3, 0xa2, 0x57, 0x72, 0x72, 0x6d, 0x42, 0xbf, 0xfc,
	0xf5, 0x17, 0x00, 0x00, 0xff, 0xff, 0x47, 0x3a, 0xdc, 0x71, 0x85, 0x01, 0x00, 0x00,
}