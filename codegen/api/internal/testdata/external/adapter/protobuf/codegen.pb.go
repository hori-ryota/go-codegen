// Code generated by protoc-gen-go. DO NOT EDIT.
// source: codegen.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type StringEnum int32

const (
	StringEnum_UNKNOWN_STRING_ENUM StringEnum = 0
	StringEnum_STRING_A            StringEnum = 1
	StringEnum_STRING_B            StringEnum = 2
	StringEnum_STRING_C            StringEnum = 3
)

var StringEnum_name = map[int32]string{
	0: "UNKNOWN_STRING_ENUM",
	1: "STRING_A",
	2: "STRING_B",
	3: "STRING_C",
}

var StringEnum_value = map[string]int32{
	"UNKNOWN_STRING_ENUM": 0,
	"STRING_A":            1,
	"STRING_B":            2,
	"STRING_C":            3,
}

func (x StringEnum) String() string {
	return proto.EnumName(StringEnum_name, int32(x))
}

func (StringEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{0}
}

type IntEnum int32

const (
	IntEnum_UNKNOWN_INT_ENUM IntEnum = 0
	IntEnum_INT_ONE          IntEnum = 1
	IntEnum_INT_THREE        IntEnum = 2
	IntEnum_INT_TWO          IntEnum = 3
)

var IntEnum_name = map[int32]string{
	0: "UNKNOWN_INT_ENUM",
	1: "INT_ONE",
	2: "INT_THREE",
	3: "INT_TWO",
}

var IntEnum_value = map[string]int32{
	"UNKNOWN_INT_ENUM": 0,
	"INT_ONE":          1,
	"INT_THREE":        2,
	"INT_TWO":          3,
}

func (x IntEnum) String() string {
	return proto.EnumName(IntEnum_name, int32(x))
}

func (IntEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{1}
}

type DoingSomethingWithOutputAndActorUsecaseInput struct {
	StringParam                     string                                                                          `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	IntParam                        int64                                                                           `protobuf:"varint,2,opt,name=int_param,json=intParam,proto3" json:"int_param,omitempty"`
	Int64Param                      int64                                                                           `protobuf:"varint,3,opt,name=int64_param,json=int64Param,proto3" json:"int64_param,omitempty"`
	UintParam                       uint64                                                                          `protobuf:"varint,4,opt,name=uint_param,json=uintParam,proto3" json:"uint_param,omitempty"`
	Uint64Param                     uint64                                                                          `protobuf:"varint,5,opt,name=uint64_param,json=uint64Param,proto3" json:"uint64_param,omitempty"`
	Float32Param                    float32                                                                         `protobuf:"fixed32,6,opt,name=float32_param,json=float32Param,proto3" json:"float32_param,omitempty"`
	Float64Param                    float64                                                                         `protobuf:"fixed64,7,opt,name=float64_param,json=float64Param,proto3" json:"float64_param,omitempty"`
	BytesParam                      []byte                                                                          `protobuf:"bytes,8,opt,name=bytes_param,json=bytesParam,proto3" json:"bytes_param,omitempty"`
	AnonymousNestedStructParam      *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam        `protobuf:"bytes,9,opt,name=anonymous_nested_struct_param,json=anonymousNestedStructParam,proto3" json:"anonymous_nested_struct_param,omitempty"`
	NamedNestedStructParam          *NamedSomeType                                                                  `protobuf:"bytes,10,opt,name=named_nested_struct_param,json=namedNestedStructParam,proto3" json:"named_nested_struct_param,omitempty"`
	StringEnumParam                 StringEnum                                                                      `protobuf:"varint,11,opt,name=string_enum_param,json=stringEnumParam,proto3,enum=codegen.StringEnum" json:"string_enum_param,omitempty"`
	IntEnumParam                    IntEnum                                                                         `protobuf:"varint,12,opt,name=int_enum_param,json=intEnumParam,proto3,enum=codegen.IntEnum" json:"int_enum_param,omitempty"`
	StringSliceParam                []string                                                                        `protobuf:"bytes,13,rep,name=string_slice_param,json=stringSliceParam,proto3" json:"string_slice_param,omitempty"`
	IntSliceParam                   []int64                                                                         `protobuf:"varint,14,rep,packed,name=int_slice_param,json=intSliceParam,proto3" json:"int_slice_param,omitempty"`
	Int64SliceParam                 []int64                                                                         `protobuf:"varint,15,rep,packed,name=int64_slice_param,json=int64SliceParam,proto3" json:"int64_slice_param,omitempty"`
	UintSliceParam                  []uint64                                                                        `protobuf:"varint,16,rep,packed,name=uint_slice_param,json=uintSliceParam,proto3" json:"uint_slice_param,omitempty"`
	Uint64SliceParam                []uint64                                                                        `protobuf:"varint,17,rep,packed,name=uint64_slice_param,json=uint64SliceParam,proto3" json:"uint64_slice_param,omitempty"`
	Float32SliceParam               []float32                                                                       `protobuf:"fixed32,18,rep,packed,name=float32_slice_param,json=float32SliceParam,proto3" json:"float32_slice_param,omitempty"`
	Float64SliceParam               []float64                                                                       `protobuf:"fixed64,19,rep,packed,name=float64_slice_param,json=float64SliceParam,proto3" json:"float64_slice_param,omitempty"`
	BytesSliceParam                 [][]byte                                                                        `protobuf:"bytes,20,rep,name=bytes_slice_param,json=bytesSliceParam,proto3" json:"bytes_slice_param,omitempty"`
	AnonymousNestedStructSliceParam []*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam `protobuf:"bytes,21,rep,name=anonymous_nested_struct_slice_param,json=anonymousNestedStructSliceParam,proto3" json:"anonymous_nested_struct_slice_param,omitempty"`
	NamedNestedStructSliceParam     []*NamedSomeType                                                                `protobuf:"bytes,22,rep,name=named_nested_struct_slice_param,json=namedNestedStructSliceParam,proto3" json:"named_nested_struct_slice_param,omitempty"`
	StringEnumSliceParam            []StringEnum                                                                    `protobuf:"varint,23,rep,packed,name=string_enum_slice_param,json=stringEnumSliceParam,proto3,enum=codegen.StringEnum" json:"string_enum_slice_param,omitempty"`
	IntEnumSliceParam               []IntEnum                                                                       `protobuf:"varint,24,rep,packed,name=int_enum_slice_param,json=intEnumSliceParam,proto3,enum=codegen.IntEnum" json:"int_enum_slice_param,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                                                                        `json:"-"`
	XXX_unrecognized                []byte                                                                          `json:"-"`
	XXX_sizecache                   int32                                                                           `json:"-"`
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) Reset() {
	*m = DoingSomethingWithOutputAndActorUsecaseInput{}
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithOutputAndActorUsecaseInput) ProtoMessage() {}
func (*DoingSomethingWithOutputAndActorUsecaseInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{0}
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput.Unmarshal(m, b)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput.Merge(m, src)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput.Size(m)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput proto.InternalMessageInfo

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetIntParam() int64 {
	if m != nil {
		return m.IntParam
	}
	return 0
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetInt64Param() int64 {
	if m != nil {
		return m.Int64Param
	}
	return 0
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetUintParam() uint64 {
	if m != nil {
		return m.UintParam
	}
	return 0
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetUint64Param() uint64 {
	if m != nil {
		return m.Uint64Param
	}
	return 0
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetFloat32Param() float32 {
	if m != nil {
		return m.Float32Param
	}
	return 0
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetFloat64Param() float64 {
	if m != nil {
		return m.Float64Param
	}
	return 0
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetBytesParam() []byte {
	if m != nil {
		return m.BytesParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetAnonymousNestedStructParam() *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam {
	if m != nil {
		return m.AnonymousNestedStructParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetNamedNestedStructParam() *NamedSomeType {
	if m != nil {
		return m.NamedNestedStructParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetStringEnumParam() StringEnum {
	if m != nil {
		return m.StringEnumParam
	}
	return StringEnum_UNKNOWN_STRING_ENUM
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetIntEnumParam() IntEnum {
	if m != nil {
		return m.IntEnumParam
	}
	return IntEnum_UNKNOWN_INT_ENUM
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetStringSliceParam() []string {
	if m != nil {
		return m.StringSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetIntSliceParam() []int64 {
	if m != nil {
		return m.IntSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetInt64SliceParam() []int64 {
	if m != nil {
		return m.Int64SliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetUintSliceParam() []uint64 {
	if m != nil {
		return m.UintSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetUint64SliceParam() []uint64 {
	if m != nil {
		return m.Uint64SliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetFloat32SliceParam() []float32 {
	if m != nil {
		return m.Float32SliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetFloat64SliceParam() []float64 {
	if m != nil {
		return m.Float64SliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetBytesSliceParam() [][]byte {
	if m != nil {
		return m.BytesSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetAnonymousNestedStructSliceParam() []*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam {
	if m != nil {
		return m.AnonymousNestedStructSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetNamedNestedStructSliceParam() []*NamedSomeType {
	if m != nil {
		return m.NamedNestedStructSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetStringEnumSliceParam() []StringEnum {
	if m != nil {
		return m.StringEnumSliceParam
	}
	return nil
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput) GetIntEnumSliceParam() []IntEnum {
	if m != nil {
		return m.IntEnumSliceParam
	}
	return nil
}

type DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) Reset() {
	*m = DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam{}
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) ProtoMessage() {}
func (*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{0, 0}
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam.Unmarshal(m, b)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam.Merge(m, src)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam.Size(m)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam proto.InternalMessageInfo

func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) Reset() {
	*m = DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam{}
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) ProtoMessage() {}
func (*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{0, 1}
}

func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam.Unmarshal(m, b)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam.Merge(m, src)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam.Size(m)
}
func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam proto.InternalMessageInfo

func (m *DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type NamedSomeType struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedSomeType) Reset()         { *m = NamedSomeType{} }
func (m *NamedSomeType) String() string { return proto.CompactTextString(m) }
func (*NamedSomeType) ProtoMessage()    {}
func (*NamedSomeType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{1}
}

func (m *NamedSomeType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedSomeType.Unmarshal(m, b)
}
func (m *NamedSomeType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedSomeType.Marshal(b, m, deterministic)
}
func (m *NamedSomeType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedSomeType.Merge(m, src)
}
func (m *NamedSomeType) XXX_Size() int {
	return xxx_messageInfo_NamedSomeType.Size(m)
}
func (m *NamedSomeType) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedSomeType.DiscardUnknown(m)
}

var xxx_messageInfo_NamedSomeType proto.InternalMessageInfo

func (m *NamedSomeType) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type DoingSomethingWithOutputAndActorUsecaseOutput struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithOutputAndActorUsecaseOutput) Reset() {
	*m = DoingSomethingWithOutputAndActorUsecaseOutput{}
}
func (m *DoingSomethingWithOutputAndActorUsecaseOutput) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithOutputAndActorUsecaseOutput) ProtoMessage() {}
func (*DoingSomethingWithOutputAndActorUsecaseOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{2}
}

func (m *DoingSomethingWithOutputAndActorUsecaseOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseOutput.Unmarshal(m, b)
}
func (m *DoingSomethingWithOutputAndActorUsecaseOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseOutput.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithOutputAndActorUsecaseOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseOutput.Merge(m, src)
}
func (m *DoingSomethingWithOutputAndActorUsecaseOutput) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseOutput.Size(m)
}
func (m *DoingSomethingWithOutputAndActorUsecaseOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithOutputAndActorUsecaseOutput proto.InternalMessageInfo

func (m *DoingSomethingWithOutputAndActorUsecaseOutput) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type DoingSomethingWithOutputWithoutActorUsecaseInput struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) Reset() {
	*m = DoingSomethingWithOutputWithoutActorUsecaseInput{}
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithOutputWithoutActorUsecaseInput) ProtoMessage() {}
func (*DoingSomethingWithOutputWithoutActorUsecaseInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{3}
}

func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseInput.Unmarshal(m, b)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseInput.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseInput.Merge(m, src)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseInput.Size(m)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseInput.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseInput proto.InternalMessageInfo

func (m *DoingSomethingWithOutputWithoutActorUsecaseInput) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type DoingSomethingWithOutputWithoutActorUsecaseOutput struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) Reset() {
	*m = DoingSomethingWithOutputWithoutActorUsecaseOutput{}
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithOutputWithoutActorUsecaseOutput) ProtoMessage() {}
func (*DoingSomethingWithOutputWithoutActorUsecaseOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{4}
}

func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseOutput.Unmarshal(m, b)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseOutput.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseOutput.Merge(m, src)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseOutput.Size(m)
}
func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseOutput.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithOutputWithoutActorUsecaseOutput proto.InternalMessageInfo

func (m *DoingSomethingWithOutputWithoutActorUsecaseOutput) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type DoingSomethingWithoutOutputAndActorUsecaseInput struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) Reset() {
	*m = DoingSomethingWithoutOutputAndActorUsecaseInput{}
}
func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithoutOutputAndActorUsecaseInput) ProtoMessage() {}
func (*DoingSomethingWithoutOutputAndActorUsecaseInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{5}
}

func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithoutOutputAndActorUsecaseInput.Unmarshal(m, b)
}
func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithoutOutputAndActorUsecaseInput.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithoutOutputAndActorUsecaseInput.Merge(m, src)
}
func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithoutOutputAndActorUsecaseInput.Size(m)
}
func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithoutOutputAndActorUsecaseInput.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithoutOutputAndActorUsecaseInput proto.InternalMessageInfo

func (m *DoingSomethingWithoutOutputAndActorUsecaseInput) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

type DoingSomethingWithoutOutputWithActorUsecaseInput struct {
	StringParam          string   `protobuf:"bytes,1,opt,name=string_param,json=stringParam,proto3" json:"string_param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) Reset() {
	*m = DoingSomethingWithoutOutputWithActorUsecaseInput{}
}
func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) String() string {
	return proto.CompactTextString(m)
}
func (*DoingSomethingWithoutOutputWithActorUsecaseInput) ProtoMessage() {}
func (*DoingSomethingWithoutOutputWithActorUsecaseInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3461500917a7cb84, []int{6}
}

func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoingSomethingWithoutOutputWithActorUsecaseInput.Unmarshal(m, b)
}
func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoingSomethingWithoutOutputWithActorUsecaseInput.Marshal(b, m, deterministic)
}
func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoingSomethingWithoutOutputWithActorUsecaseInput.Merge(m, src)
}
func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) XXX_Size() int {
	return xxx_messageInfo_DoingSomethingWithoutOutputWithActorUsecaseInput.Size(m)
}
func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DoingSomethingWithoutOutputWithActorUsecaseInput.DiscardUnknown(m)
}

var xxx_messageInfo_DoingSomethingWithoutOutputWithActorUsecaseInput proto.InternalMessageInfo

func (m *DoingSomethingWithoutOutputWithActorUsecaseInput) GetStringParam() string {
	if m != nil {
		return m.StringParam
	}
	return ""
}

func init() {
	proto.RegisterEnum("codegen.StringEnum", StringEnum_name, StringEnum_value)
	proto.RegisterEnum("codegen.IntEnum", IntEnum_name, IntEnum_value)
	proto.RegisterType((*DoingSomethingWithOutputAndActorUsecaseInput)(nil), "codegen.DoingSomethingWithOutputAndActorUsecaseInput")
	proto.RegisterType((*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam)(nil), "codegen.DoingSomethingWithOutputAndActorUsecaseInput.AnonymousNestedStructParam")
	proto.RegisterType((*DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam)(nil), "codegen.DoingSomethingWithOutputAndActorUsecaseInput.AnonymousNestedStructSliceParam")
	proto.RegisterType((*NamedSomeType)(nil), "codegen.NamedSomeType")
	proto.RegisterType((*DoingSomethingWithOutputAndActorUsecaseOutput)(nil), "codegen.DoingSomethingWithOutputAndActorUsecaseOutput")
	proto.RegisterType((*DoingSomethingWithOutputWithoutActorUsecaseInput)(nil), "codegen.DoingSomethingWithOutputWithoutActorUsecaseInput")
	proto.RegisterType((*DoingSomethingWithOutputWithoutActorUsecaseOutput)(nil), "codegen.DoingSomethingWithOutputWithoutActorUsecaseOutput")
	proto.RegisterType((*DoingSomethingWithoutOutputAndActorUsecaseInput)(nil), "codegen.DoingSomethingWithoutOutputAndActorUsecaseInput")
	proto.RegisterType((*DoingSomethingWithoutOutputWithActorUsecaseInput)(nil), "codegen.DoingSomethingWithoutOutputWithActorUsecaseInput")
}

func init() { proto.RegisterFile("codegen.proto", fileDescriptor_3461500917a7cb84) }

var fileDescriptor_3461500917a7cb84 = []byte{
	// 968 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x97, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0xc7, 0x7f, 0x13, 0xf7, 0xb7, 0x6d, 0x4e, 0x9c, 0xc6, 0x71, 0x4b, 0xbb, 0xa4, 0xaa, 0xe2,
	0x76, 0x25, 0xb0, 0x42, 0xeb, 0x40, 0x16, 0x2a, 0xe0, 0x66, 0x95, 0x6e, 0x23, 0xb6, 0x20, 0x92,
	0xe2, 0x24, 0x14, 0x21, 0xa4, 0xc8, 0x4d, 0x66, 0x53, 0x4b, 0xf1, 0x8c, 0x15, 0xcf, 0xa0, 0xcd,
	0x4b, 0x70, 0xc3, 0x1d, 0x17, 0x7b, 0x01, 0xe2, 0x25, 0x78, 0x04, 0x9e, 0x0a, 0x79, 0x3c, 0x4e,
	0xa6, 0xcd, 0x9f, 0x4d, 0x0b, 0x57, 0x9d, 0x39, 0xe7, 0x7b, 0x3e, 0x73, 0x66, 0xfc, 0x9d, 0xd8,
	0x85, 0x7c, 0x9f, 0x0e, 0xf0, 0x10, 0x13, 0x27, 0x1c, 0x53, 0x46, 0xcd, 0x4d, 0x39, 0x2d, 0x1d,
	0x0c, 0x29, 0x1d, 0x8e, 0x70, 0x55, 0x84, 0x6f, 0xf8, 0xeb, 0x2a, 0x0e, 0x42, 0x36, 0x49, 0x54,
	0xc7, 0x6f, 0x75, 0x38, 0xb9, 0xa0, 0x3e, 0x19, 0xb6, 0x69, 0x80, 0xd9, 0xad, 0x4f, 0x86, 0xd7,
	0x3e, 0xbb, 0x6d, 0x71, 0x16, 0x72, 0x56, 0x27, 0x83, 0x7a, 0x9f, 0xd1, 0x71, 0x37, 0xc2, 0x7d,
	0x2f, 0xc2, 0x97, 0x24, 0xe4, 0xcc, 0x3c, 0x02, 0x3d, 0x62, 0x63, 0x9f, 0x0c, 0x7b, 0xa1, 0x37,
	0xf6, 0x82, 0xa7, 0xc8, 0x42, 0x76, 0xd6, 0xcd, 0x25, 0xb1, 0xab, 0x38, 0x64, 0x1e, 0x40, 0xd6,
	0x27, 0x4c, 0xe6, 0x33, 0x16, 0xb2, 0x35, 0x77, 0xcb, 0x27, 0x2c, 0x49, 0x96, 0x21, 0xe7, 0x13,
	0x76, 0xf6, 0xa9, 0x4c, 0x6b, 0x22, 0x0d, 0x22, 0x94, 0x08, 0x0e, 0x01, 0xf8, 0xac, 0x7c, 0xc3,
	0x42, 0xf6, 0x86, 0x9b, 0xe5, 0xd3, 0xfa, 0x23, 0xd0, 0xb9, 0x0a, 0xf8, 0xbf, 0x10, 0xe4, 0xb8,
	0x42, 0x78, 0x06, 0xf9, 0xd7, 0x23, 0xea, 0xb1, 0xe7, 0x35, 0xa9, 0x79, 0x62, 0x21, 0x3b, 0xe3,
	0xea, 0x32, 0x78, 0x57, 0x34, 0x05, 0x6d, 0x5a, 0xc8, 0x46, 0x52, 0x94, 0x92, 0xca, 0x90, 0xbb,
	0x99, 0x30, 0x1c, 0x49, 0xc9, 0x96, 0x85, 0x6c, 0xdd, 0x05, 0x11, 0x4a, 0x04, 0xbf, 0x22, 0x38,
	0xf4, 0x08, 0x25, 0x93, 0x80, 0xf2, 0xa8, 0x47, 0x70, 0xc4, 0xf0, 0xa0, 0x17, 0xb1, 0x31, 0xef,
	0xa7, 0x1b, 0xc8, 0x5a, 0xc8, 0xce, 0xd5, 0xae, 0x9c, 0xf4, 0xe1, 0x3c, 0xe4, 0xb0, 0x9d, 0x7a,
	0x8a, 0x6e, 0x0a, 0x72, 0x5b, 0x80, 0xc5, 0xca, 0x6e, 0xc9, 0x5b, 0x9a, 0x33, 0xbf, 0x83, 0xf7,
	0x89, 0x17, 0xe0, 0xc1, 0xc2, 0x86, 0x40, 0x34, 0xb4, 0x37, 0x6d, 0xa8, 0x19, 0x2b, 0xe3, 0x86,
	0x3a, 0x93, 0x10, 0xbb, 0x7b, 0xa2, 0x70, 0x1e, 0xf9, 0x02, 0x8a, 0xf2, 0xb1, 0x63, 0xc2, 0x03,
	0x89, 0xca, 0x59, 0xc8, 0xde, 0xae, 0xed, 0x4c, 0x51, 0x6d, 0xa1, 0x68, 0x10, 0x1e, 0xb8, 0x85,
	0x68, 0x3a, 0x4e, 0x00, 0x67, 0xb0, 0x1d, 0x3f, 0x55, 0xa5, 0x5a, 0x17, 0xd5, 0xc6, 0xb4, 0xfa,
	0x92, 0x30, 0x51, 0xaa, 0xfb, 0xc9, 0x20, 0xa9, 0x3b, 0x01, 0x53, 0x2e, 0x1c, 0x8d, 0xfc, 0x3e,
	0x96, 0xb5, 0x79, 0x4b, 0xb3, 0xb3, 0xae, 0x91, 0x64, 0xda, 0x71, 0x22, 0x51, 0x7f, 0x00, 0x85,
	0x78, 0x15, 0x55, 0xba, 0x6d, 0x69, 0xb6, 0xe6, 0xe6, 0x7d, 0xc2, 0x14, 0x5d, 0x05, 0x8a, 0x89,
	0x89, 0x54, 0x65, 0x41, 0x28, 0x0b, 0x22, 0xa1, 0x68, 0x6d, 0x30, 0xf8, 0x7d, 0xa8, 0x61, 0x69,
	0xf6, 0x86, 0xbb, 0xcd, 0xef, 0x52, 0x4f, 0xc0, 0xe4, 0xf3, 0xd8, 0xa2, 0xd0, 0x1a, 0xfc, 0x3e,
	0xd7, 0x81, 0x9d, 0xd4, 0xa6, 0xaa, 0xdc, 0xb4, 0x34, 0x3b, 0xe3, 0x16, 0x65, 0x6a, 0x81, 0xfe,
	0x1e, 0x7e, 0xc7, 0xd2, 0x6c, 0x24, 0xf5, 0x77, 0xf8, 0x15, 0x28, 0x26, 0xe6, 0x55, 0xd5, 0xbb,
	0x96, 0x66, 0xeb, 0x6e, 0x41, 0x24, 0x14, 0xed, 0x1f, 0x08, 0x9e, 0x2d, 0xf3, 0xb1, 0x5a, 0xfe,
	0x9e, 0xa5, 0xd9, 0xb9, 0x5a, 0xf7, 0x3f, 0x74, 0xf3, 0xac, 0x09, 0xb7, 0xec, 0xad, 0x16, 0x98,
	0x3f, 0x41, 0x79, 0x91, 0xaf, 0xd5, 0x06, 0xf7, 0x44, 0x83, 0xcb, 0xdc, 0x7d, 0x30, 0xe7, 0x6e,
	0x85, 0xfe, 0x35, 0xec, 0xab, 0x16, 0x57, 0xa9, 0xfb, 0x96, 0xb6, 0xcc, 0xe8, 0xbb, 0x33, 0xa3,
	0x2b, 0xac, 0x3a, 0xec, 0x4e, 0xdd, 0xae, 0x82, 0x9e, 0x0a, 0xd0, 0xbc, 0xe7, 0x8b, 0xd2, 0xf3,
	0x33, 0x44, 0xe9, 0x05, 0x94, 0x96, 0x5f, 0xff, 0x35, 0x7e, 0x86, 0x4b, 0x17, 0x50, 0x7e, 0xc7,
	0x89, 0xaf, 0x41, 0x39, 0xae, 0x41, 0xfe, 0xce, 0x19, 0xae, 0x53, 0xe3, 0xc2, 0xe9, 0x9a, 0xc6,
	0x48, 0x82, 0xeb, 0x30, 0xbb, 0xf0, 0xf1, 0x32, 0x66, 0x3c, 0xa2, 0x9c, 0x3d, 0xe6, 0x5d, 0x75,
	0xfc, 0x3d, 0x7c, 0xf2, 0x00, 0xec, 0xfa, 0xed, 0x76, 0xa0, 0x3a, 0xcf, 0xa5, 0x9c, 0xfd, 0xab,
	0x37, 0xeb, 0xe2, 0x43, 0x98, 0x52, 0xe3, 0xc9, 0x63, 0xb0, 0x95, 0x16, 0xc0, 0xcc, 0xd1, 0xe6,
	0x3e, 0xec, 0x74, 0x9b, 0xdf, 0x34, 0x5b, 0xd7, 0xcd, 0x5e, 0xbb, 0xe3, 0x5e, 0x36, 0xbf, 0xea,
	0x35, 0x9a, 0xdd, 0x6f, 0x8d, 0xff, 0x99, 0x3a, 0x6c, 0xc9, 0x40, 0xdd, 0x40, 0xca, 0xec, 0xdc,
	0xc8, 0x28, 0xb3, 0x97, 0x86, 0x56, 0x79, 0x05, 0x9b, 0xd2, 0xd9, 0xe6, 0x2e, 0x18, 0x29, 0xed,
	0xb2, 0xd9, 0x49, 0x51, 0x39, 0xd8, 0x8c, 0x67, 0xad, 0x66, 0xc3, 0x40, 0x66, 0x1e, 0xb2, 0xf1,
	0xa4, 0xf3, 0xca, 0x6d, 0x34, 0x8c, 0x4c, 0x9a, 0xeb, 0x5c, 0xb7, 0x0c, 0xad, 0xf6, 0x17, 0x82,
	0x0f, 0xd7, 0xf4, 0x92, 0xf9, 0x0b, 0x82, 0xc3, 0x0b, 0xba, 0x42, 0x68, 0x7e, 0xf6, 0xa8, 0x1f,
	0xae, 0xd2, 0xd9, 0x43, 0xcb, 0x92, 0x60, 0xed, 0x6f, 0x04, 0x1f, 0x3d, 0xc0, 0x5d, 0xe6, 0x6f,
	0x08, 0x8e, 0x16, 0x6e, 0x40, 0x15, 0x9b, 0x5f, 0xbc, 0xb3, 0x9b, 0x65, 0x17, 0xa2, 0xf4, 0xe5,
	0x63, 0x4a, 0xe5, 0x66, 0x7e, 0x47, 0x50, 0x59, 0xdf, 0xd2, 0x26, 0x03, 0xeb, 0xde, 0x56, 0xe6,
	0xa4, 0xe6, 0xe7, 0x2b, 0xda, 0x59, 0x79, 0x57, 0x4a, 0x7b, 0x4e, 0xf2, 0x51, 0xeb, 0xa4, 0x1f,
	0xb5, 0x4e, 0x23, 0xfe, 0xa8, 0xad, 0xfd, 0xb9, 0xf0, 0xc4, 0x97, 0xde, 0x10, 0xf3, 0xe7, 0xb9,
	0x03, 0x9f, 0xd7, 0xae, 0x3c, 0xf0, 0xd5, 0x97, 0x6f, 0x59, 0x9f, 0xe7, 0x6f, 0x11, 0xd4, 0xfa,
	0x34, 0x70, 0x86, 0x3e, 0xbb, 0xe5, 0x37, 0xce, 0x2d, 0x1d, 0xfb, 0xe3, 0x09, 0x65, 0x9e, 0x33,
	0xa4, 0xe9, 0x5a, 0x5e, 0xe8, 0x3b, 0xf8, 0x8d, 0x17, 0x84, 0x23, 0x9c, 0xae, 0x7f, 0xae, 0xbf,
	0x4c, 0x06, 0x57, 0x31, 0xed, 0x0a, 0xfd, 0xf8, 0x83, 0xac, 0xef, 0xd3, 0xa0, 0x1a, 0x33, 0x4e,
	0x05, 0xa4, 0x3a, 0xa4, 0xa7, 0xb2, 0xa2, 0x9a, 0xfe, 0xf5, 0x42, 0xbf, 0xea, 0x13, 0x86, 0xc7,
	0xc4, 0x1b, 0x55, 0x19, 0x8e, 0xd8, 0xc0, 0x63, 0x5e, 0x15, 0xbf, 0x91, 0x11, 0x6f, 0xe0, 0x85,
	0x0c, 0x8f, 0xa7, 0xff, 0x23, 0xdc, 0x3c, 0x11, 0xa3, 0xe7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x14, 0x68, 0x62, 0xfa, 0x55, 0x0c, 0x00, 0x00,
}
