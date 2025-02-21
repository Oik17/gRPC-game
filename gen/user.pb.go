// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: proto/user.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Points        int32                  `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

type Quiz struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Question      string                 `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Answer        string                 `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Quiz) Reset() {
	*x = Quiz{}
	mi := &file_proto_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Quiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quiz) ProtoMessage() {}

func (x *Quiz) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quiz.ProtoReflect.Descriptor instead.
func (*Quiz) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *Quiz) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Quiz) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Quiz) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Answer        string                 `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question      string                 `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	IsCorrect     bool                   `protobuf:"varint,4,opt,name=is_correct,json=isCorrect,proto3" json:"is_correct,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_proto_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Response) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Response) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Response) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

type Connect struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active        bool                   `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Connect) Reset() {
	*x = Connect{}
	mi := &file_proto_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *Connect) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Connect) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type Leaderboard struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Leaderboard) Reset() {
	*x = Leaderboard{}
	mi := &file_proto_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Leaderboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaderboard) ProtoMessage() {}

func (x *Leaderboard) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaderboard.ProtoReflect.Descriptor instead.
func (*Leaderboard) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *Leaderboard) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Close struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Close) Reset() {
	*x = Close{}
	mi := &file_proto_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{5}
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x4a, 0x0a,
	0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x7d, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x22, 0x41, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x2d, 0x0a, 0x0b, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x32, 0xe7, 0x01, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x1a, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x30,
	0x01, 0x12, 0x25, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0b,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x1a, 0x0b, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x1a, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData []byte
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_proto_rawDesc), len(file_proto_user_proto_rawDesc)))
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_user_proto_goTypes = []any{
	(*User)(nil),        // 0: game.User
	(*Quiz)(nil),        // 1: game.Quiz
	(*Response)(nil),    // 2: game.Response
	(*Connect)(nil),     // 3: game.Connect
	(*Leaderboard)(nil), // 4: game.Leaderboard
	(*Close)(nil),       // 5: game.Close
}
var file_proto_user_proto_depIdxs = []int32{
	0, // 0: game.Response.user:type_name -> game.User
	0, // 1: game.Connect.user:type_name -> game.User
	0, // 2: game.Leaderboard.user:type_name -> game.User
	3, // 3: game.gameService.CreateStream:input_type -> game.Connect
	2, // 4: game.gameService.SubmitAnswer:input_type -> game.Response
	5, // 5: game.gameService.StreamLeaderboard:input_type -> game.Close
	5, // 6: game.gameService.StartGame:input_type -> game.Close
	5, // 7: game.gameService.EndGame:input_type -> game.Close
	2, // 8: game.gameService.CreateStream:output_type -> game.Response
	5, // 9: game.gameService.SubmitAnswer:output_type -> game.Close
	0, // 10: game.gameService.StreamLeaderboard:output_type -> game.User
	5, // 11: game.gameService.StartGame:output_type -> game.Close
	5, // 12: game.gameService.EndGame:output_type -> game.Close
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_proto_rawDesc), len(file_proto_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}
