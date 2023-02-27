// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/player.proto

package go_playlist_grpc

import (
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

type NewPlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NewPlaylistResponse) Reset() {
	*x = NewPlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlaylistResponse) ProtoMessage() {}

func (x *NewPlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlaylistResponse.ProtoReflect.Descriptor instead.
func (*NewPlaylistResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{0}
}

func (x *NewPlaylistResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Duration int64  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{1}
}

func (x *Song) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Song) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{2}
}

type SongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SongRequest) Reset() {
	*x = SongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongRequest) ProtoMessage() {}

func (x *SongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongRequest.ProtoReflect.Descriptor instead.
func (*SongRequest) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{3}
}

func (x *SongRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Duration int64  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *SongResponse) Reset() {
	*x = SongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongResponse) ProtoMessage() {}

func (x *SongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongResponse.ProtoReflect.Descriptor instead.
func (*SongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{4}
}

func (x *SongResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SongResponse) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type AddSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddSongResponse) Reset() {
	*x = AddSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongResponse) ProtoMessage() {}

func (x *AddSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongResponse.ProtoReflect.Descriptor instead.
func (*AddSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{5}
}

func (x *AddSongResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetSongsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*Song `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *GetSongsResponse) Reset() {
	*x = GetSongsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsResponse) ProtoMessage() {}

func (x *GetSongsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsResponse.ProtoReflect.Descriptor instead.
func (*GetSongsResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{6}
}

func (x *GetSongsResponse) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type UpdateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateSongResponse) Reset() {
	*x = UpdateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongResponse) ProtoMessage() {}

func (x *UpdateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongResponse.ProtoReflect.Descriptor instead.
func (*UpdateSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSongResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteSongResponse) Reset() {
	*x = DeleteSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongResponse) ProtoMessage() {}

func (x *DeleteSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongResponse.ProtoReflect.Descriptor instead.
func (*DeleteSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSongResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{9}
}

func (x *PlayResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PauseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PauseResponse) Reset() {
	*x = PauseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PauseResponse) ProtoMessage() {}

func (x *PauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PauseResponse.ProtoReflect.Descriptor instead.
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{10}
}

func (x *PauseResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type NextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NextResponse) Reset() {
	*x = NextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextResponse) ProtoMessage() {}

func (x *NextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextResponse.ProtoReflect.Descriptor instead.
func (*NextResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{11}
}

func (x *NextResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PrevResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PrevResponse) Reset() {
	*x = PrevResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrevResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrevResponse) ProtoMessage() {}

func (x *PrevResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrevResponse.ProtoReflect.Descriptor instead.
func (*PrevResponse) Descriptor() ([]byte, []int) {
	return file_proto_player_proto_rawDescGZIP(), []int{12}
}

func (x *PrevResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_player_proto protoreflect.FileDescriptor

var file_proto_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x2f,
	0x0a, 0x13, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x36, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x21, 0x0a, 0x0b, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x6c,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x28, 0x0a, 0x0c, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x72, 0x65,
	0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xd2, 0x04, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53,
	0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53,
	0x6f, 0x6e, 0x67, 0x1a, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x0f, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x6f, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x6c,
	0x61, 0x79, 0x12, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x05, 0x50, 0x61, 0x75, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x72, 0x65, 0x76, 0x12, 0x0f, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x6f, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_player_proto_rawDescOnce sync.Once
	file_proto_player_proto_rawDescData = file_proto_player_proto_rawDesc
)

func file_proto_player_proto_rawDescGZIP() []byte {
	file_proto_player_proto_rawDescOnce.Do(func() {
		file_proto_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_proto_rawDescData)
	})
	return file_proto_player_proto_rawDescData
}

var file_proto_player_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_player_proto_goTypes = []interface{}{
	(*NewPlaylistResponse)(nil), // 0: playlist.NewPlaylistResponse
	(*Song)(nil),                // 1: playlist.Song
	(*Empty)(nil),               // 2: playlist.Empty
	(*SongRequest)(nil),         // 3: playlist.SongRequest
	(*SongResponse)(nil),        // 4: playlist.SongResponse
	(*AddSongResponse)(nil),     // 5: playlist.AddSongResponse
	(*GetSongsResponse)(nil),    // 6: playlist.GetSongsResponse
	(*UpdateSongResponse)(nil),  // 7: playlist.UpdateSongResponse
	(*DeleteSongResponse)(nil),  // 8: playlist.DeleteSongResponse
	(*PlayResponse)(nil),        // 9: playlist.PlayResponse
	(*PauseResponse)(nil),       // 10: playlist.PauseResponse
	(*NextResponse)(nil),        // 11: playlist.NextResponse
	(*PrevResponse)(nil),        // 12: playlist.PrevResponse
}
var file_proto_player_proto_depIdxs = []int32{
	1,  // 0: playlist.GetSongsResponse.songs:type_name -> playlist.Song
	2,  // 1: playlist.PlaylistService.NewPlaylist:input_type -> playlist.Empty
	1,  // 2: playlist.PlaylistService.AddSong:input_type -> playlist.Song
	2,  // 3: playlist.PlaylistService.GetSongs:input_type -> playlist.Empty
	3,  // 4: playlist.PlaylistService.GetSong:input_type -> playlist.SongRequest
	1,  // 5: playlist.PlaylistService.UpdateSong:input_type -> playlist.Song
	3,  // 6: playlist.PlaylistService.DeleteSong:input_type -> playlist.SongRequest
	2,  // 7: playlist.PlaylistService.Play:input_type -> playlist.Empty
	2,  // 8: playlist.PlaylistService.Pause:input_type -> playlist.Empty
	2,  // 9: playlist.PlaylistService.Next:input_type -> playlist.Empty
	2,  // 10: playlist.PlaylistService.Prev:input_type -> playlist.Empty
	0,  // 11: playlist.PlaylistService.NewPlaylist:output_type -> playlist.NewPlaylistResponse
	5,  // 12: playlist.PlaylistService.AddSong:output_type -> playlist.AddSongResponse
	6,  // 13: playlist.PlaylistService.GetSongs:output_type -> playlist.GetSongsResponse
	4,  // 14: playlist.PlaylistService.GetSong:output_type -> playlist.SongResponse
	7,  // 15: playlist.PlaylistService.UpdateSong:output_type -> playlist.UpdateSongResponse
	8,  // 16: playlist.PlaylistService.DeleteSong:output_type -> playlist.DeleteSongResponse
	9,  // 17: playlist.PlaylistService.Play:output_type -> playlist.PlayResponse
	10, // 18: playlist.PlaylistService.Pause:output_type -> playlist.PauseResponse
	11, // 19: playlist.PlaylistService.Next:output_type -> playlist.NextResponse
	12, // 20: playlist.PlaylistService.Prev:output_type -> playlist.PrevResponse
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_player_proto_init() }
func file_proto_player_proto_init() {
	if File_proto_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlaylistResponse); i {
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
		file_proto_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
		file_proto_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongRequest); i {
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
		file_proto_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongResponse); i {
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
		file_proto_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSongResponse); i {
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
		file_proto_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSongsResponse); i {
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
		file_proto_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongResponse); i {
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
		file_proto_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongResponse); i {
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
		file_proto_player_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
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
		file_proto_player_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PauseResponse); i {
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
		file_proto_player_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextResponse); i {
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
		file_proto_player_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrevResponse); i {
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
			RawDescriptor: file_proto_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_player_proto_goTypes,
		DependencyIndexes: file_proto_player_proto_depIdxs,
		MessageInfos:      file_proto_player_proto_msgTypes,
	}.Build()
	File_proto_player_proto = out.File
	file_proto_player_proto_rawDesc = nil
	file_proto_player_proto_goTypes = nil
	file_proto_player_proto_depIdxs = nil
}