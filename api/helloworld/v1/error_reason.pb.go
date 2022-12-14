// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.6
// source: helloworld/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	// 框架级（通用错误码）
	ErrorReason_UnSpecified             ErrorReason = 0
	ErrorReason_UNKNOWN                 ErrorReason = -1
	ErrorReason_PERMISSION_ERROR        ErrorReason = 4
	ErrorReason_SERVICE_ERROR           ErrorReason = 5
	ErrorReason_RETCODE_NOT_REGIST      ErrorReason = 91
	ErrorReason_VERIFY_FAIL             ErrorReason = 100
	ErrorReason_TIME_OUT                ErrorReason = 110
	ErrorReason_DATA_FAIL               ErrorReason = 120
	ErrorReason_USER_PERMISSIONS        ErrorReason = 140
	ErrorReason_SERVICE_UNAVAILABLE     ErrorReason = 150
	ErrorReason_RATELIMIT               ErrorReason = 153
	ErrorReason_CIRCUITBREAKER          ErrorReason = 154
	ErrorReason_MISSING_ACTION          ErrorReason = 160
	ErrorReason_MISSING_SIGNATURE       ErrorReason = 170
	ErrorReason_ERROR_SIGNATURE         ErrorReason = 171
	ErrorReason_USER_NOT_EXISTS         ErrorReason = 172
	ErrorReason_MISSING_API_VERSION     ErrorReason = 180
	ErrorReason_API_VERSION_ERROR       ErrorReason = 190
	ErrorReason_GET_NAMES_ERROR         ErrorReason = 200
	ErrorReason_PARAMS_RANGE_ERROR      ErrorReason = 210
	ErrorReason_MISSING_PARAMS          ErrorReason = 220
	ErrorReason_VALIDATOR               ErrorReason = 221
	ErrorReason_PARAMS_ERROR            ErrorReason = 230
	ErrorReason_CODEC                   ErrorReason = 231
	ErrorReason_LACK_OF_BALANCE         ErrorReason = 250
	ErrorReason_DEDUCTIONS_ERROR        ErrorReason = 260
	ErrorReason_SETTLEMENT_ERROR        ErrorReason = 270
	ErrorReason_NO_ACTIVE_BACKEND_ERROR ErrorReason = 280
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:   "UnSpecified",
		-1:  "UNKNOWN",
		4:   "PERMISSION_ERROR",
		5:   "SERVICE_ERROR",
		91:  "RETCODE_NOT_REGIST",
		100: "VERIFY_FAIL",
		110: "TIME_OUT",
		120: "DATA_FAIL",
		140: "USER_PERMISSIONS",
		150: "SERVICE_UNAVAILABLE",
		153: "RATELIMIT",
		154: "CIRCUITBREAKER",
		160: "MISSING_ACTION",
		170: "MISSING_SIGNATURE",
		171: "ERROR_SIGNATURE",
		172: "USER_NOT_EXISTS",
		180: "MISSING_API_VERSION",
		190: "API_VERSION_ERROR",
		200: "GET_NAMES_ERROR",
		210: "PARAMS_RANGE_ERROR",
		220: "MISSING_PARAMS",
		221: "VALIDATOR",
		230: "PARAMS_ERROR",
		231: "CODEC",
		250: "LACK_OF_BALANCE",
		260: "DEDUCTIONS_ERROR",
		270: "SETTLEMENT_ERROR",
		280: "NO_ACTIVE_BACKEND_ERROR",
	}
	ErrorReason_value = map[string]int32{
		"UnSpecified":             0,
		"UNKNOWN":                 -1,
		"PERMISSION_ERROR":        4,
		"SERVICE_ERROR":           5,
		"RETCODE_NOT_REGIST":      91,
		"VERIFY_FAIL":             100,
		"TIME_OUT":                110,
		"DATA_FAIL":               120,
		"USER_PERMISSIONS":        140,
		"SERVICE_UNAVAILABLE":     150,
		"RATELIMIT":               153,
		"CIRCUITBREAKER":          154,
		"MISSING_ACTION":          160,
		"MISSING_SIGNATURE":       170,
		"ERROR_SIGNATURE":         171,
		"USER_NOT_EXISTS":         172,
		"MISSING_API_VERSION":     180,
		"API_VERSION_ERROR":       190,
		"GET_NAMES_ERROR":         200,
		"PARAMS_RANGE_ERROR":      210,
		"MISSING_PARAMS":          220,
		"VALIDATOR":               221,
		"PARAMS_ERROR":            230,
		"CODEC":                   231,
		"LACK_OF_BALANCE":         250,
		"DEDUCTIONS_ERROR":        260,
		"SETTLEMENT_ERROR":        270,
		"NO_ACTIVE_BACKEND_ERROR": 280,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_helloworld_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var file_helloworld_v1_error_reason_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50005,
		Name:          "helloworld.v1.uerrreason",
		Tag:           "bytes,50005,opt,name=uerrreason",
		Filename:      "helloworld/v1/error_reason.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string uerrreason = 50005;
	E_Uerrreason = &file_helloworld_v1_error_reason_proto_extTypes[0]
)

var File_helloworld_v1_error_reason_proto protoreflect.FileDescriptor

var file_helloworld_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x20, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf3, 0x0a, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x1a,
	0x1e, 0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5, 0x18, 0x16, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x2e, 0x12,
	0x1a, 0x0a, 0x10, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x35, 0x0a, 0x12, 0x52, 0x45, 0x54, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x10, 0x5b, 0x1a, 0x1d, 0xa8, 0x45,
	0xf4, 0x03, 0xaa, 0xb5, 0x18, 0x15, 0x52, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x20, 0x69, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x64, 0x1a, 0x13, 0xa8, 0x45,
	0xf4, 0x03, 0xaa, 0xb5, 0x18, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x20, 0x66, 0x61, 0x69,
	0x6c, 0x12, 0x1e, 0x0a, 0x08, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x6e, 0x1a,
	0x10, 0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5, 0x18, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x20, 0x6f, 0x75,
	0x74, 0x12, 0x20, 0x0a, 0x09, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x78,
	0x1a, 0x11, 0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5, 0x18, 0x09, 0x44, 0x61, 0x74, 0x61, 0x20, 0x66,
	0x61, 0x69, 0x6c, 0x12, 0x2f, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x8c, 0x01, 0x1a, 0x18, 0xa8, 0x45, 0xf4, 0x03,
	0xaa, 0xb5, 0x18, 0x10, 0x55, 0x73, 0x65, 0x72, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x96, 0x01, 0x1a, 0x1b,
	0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5, 0x18, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x52,
	0x41, 0x54, 0x45, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x99, 0x01, 0x1a, 0x26, 0xa8, 0x45, 0xf4,
	0x03, 0xaa, 0xb5, 0x18, 0x1e, 0x41, 0x50, 0x49, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x20, 0x72, 0x61, 0x74, 0x65, 0x20, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x73, 0x20, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x77, 0x0a, 0x0e, 0x43, 0x49, 0x52, 0x43, 0x55, 0x49, 0x54, 0x42, 0x52,
	0x45, 0x41, 0x4b, 0x45, 0x52, 0x10, 0x9a, 0x01, 0x1a, 0x62, 0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5,
	0x18, 0x5a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6c, 0x79, 0x20, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x2c, 0x20, 0x77, 0x65, 0x20, 0x61, 0x72, 0x65, 0x20, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x20, 0x68, 0x61, 0x72, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x20, 0x69, 0x74, 0x2c, 0x20, 0x70, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x20, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x12, 0x2b, 0x0a, 0x0e,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xa0,
	0x01, 0x1a, 0x16, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x11, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0xaa,
	0x01, 0x1a, 0x19, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x11, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x0f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10,
	0xab, 0x01, 0x1a, 0x1f, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x17, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x56, 0x65, 0x72, 0x66, 0x79, 0x41, 0x43, 0x20, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0xac, 0x01, 0x1a, 0x17, 0xa8, 0x45, 0x90, 0x03, 0xaa,
	0xb5, 0x18, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x20, 0x4e, 0x6f, 0x74, 0x20, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x35, 0x0a, 0x13, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x50,
	0x49, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0xb4, 0x01, 0x1a, 0x1b, 0xa8, 0x45,
	0x90, 0x03, 0xaa, 0xb5, 0x18, 0x13, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x70,
	0x69, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x11, 0x41, 0x50, 0x49,
	0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xbe,
	0x01, 0x1a, 0x19, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x11, 0x41, 0x70, 0x69, 0x20, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x0f,
	0x47, 0x45, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0xc8, 0x01, 0x1a, 0x17, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x0f, 0x47, 0x65, 0x74, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x12, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x53, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0xd2, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x2b, 0x0a, 0x0e, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53, 0x10, 0xdc, 0x01, 0x1a,
	0x16, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x41, 0x0a, 0x09, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x41, 0x54, 0x4f, 0x52, 0x10, 0xdd, 0x01, 0x1a, 0x31, 0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5, 0x18,
	0x29, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x20, 0x65, 0x78, 0x63, 0x65, 0x65,
	0x64, 0x73, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2c, 0x20, 0x70, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x20, 0x72, 0x65, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x12, 0x2f, 0x0a, 0x0c, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe6, 0x01, 0x1a, 0x1c, 0xa8,
	0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x14, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x43,
	0x4f, 0x44, 0x45, 0x43, 0x10, 0xe7, 0x01, 0x1a, 0x2f, 0xa8, 0x45, 0xf4, 0x03, 0xaa, 0xb5, 0x18,
	0x27, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2c, 0x20, 0x70, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x20,
	0x72, 0x65, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x12, 0x2d, 0x0a, 0x0f, 0x4c, 0x41, 0x43, 0x4b,
	0x5f, 0x4f, 0x46, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x10, 0xfa, 0x01, 0x1a, 0x17,
	0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x0f, 0x4c, 0x61, 0x63, 0x6b, 0x20, 0x6f, 0x66, 0x20,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x10, 0x44, 0x45, 0x44, 0x55, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x84, 0x02, 0x1a, 0x18,
	0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x10, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x10, 0x53, 0x45, 0x54, 0x54,
	0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x8e, 0x02, 0x1a,
	0x18, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x17, 0x4e, 0x4f, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x98, 0x02, 0x1a, 0x1f, 0xa8, 0x45, 0x90, 0x03, 0xaa, 0xb5, 0x18,
	0x17, 0x6e, 0x6f, 0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x3a, 0x46,
	0x0a, 0x0a, 0x75, 0x65, 0x72, 0x72, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd5, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x65, 0x72, 0x72, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x66, 0x0a, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x62, 0x75, 0x67, 0x74, 0x6f, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2f, 0x75, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x0f, 0x41,
	0x50, 0x49, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_v1_error_reason_proto_rawDescOnce sync.Once
	file_helloworld_v1_error_reason_proto_rawDescData = file_helloworld_v1_error_reason_proto_rawDesc
)

func file_helloworld_v1_error_reason_proto_rawDescGZIP() []byte {
	file_helloworld_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_error_reason_proto_rawDescData)
	})
	return file_helloworld_v1_error_reason_proto_rawDescData
}

var file_helloworld_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_helloworld_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0),                      // 0: helloworld.v1.ErrorReason
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_helloworld_v1_error_reason_proto_depIdxs = []int32{
	1, // 0: helloworld.v1.uerrreason:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_v1_error_reason_proto_init() }
func file_helloworld_v1_error_reason_proto_init() {
	if File_helloworld_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_helloworld_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_helloworld_v1_error_reason_proto_enumTypes,
		ExtensionInfos:    file_helloworld_v1_error_reason_proto_extTypes,
	}.Build()
	File_helloworld_v1_error_reason_proto = out.File
	file_helloworld_v1_error_reason_proto_rawDesc = nil
	file_helloworld_v1_error_reason_proto_goTypes = nil
	file_helloworld_v1_error_reason_proto_depIdxs = nil
}
