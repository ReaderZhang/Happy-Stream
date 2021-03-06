// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/order/v1/order.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type OrderAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *OrderAddReq) Reset() {
	*x = OrderAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddReq) ProtoMessage() {}

func (x *OrderAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddReq.ProtoReflect.Descriptor instead.
func (*OrderAddReq) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *OrderAddReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderAddReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type OrderAddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderAddReply) Reset() {
	*x = OrderAddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAddReply) ProtoMessage() {}

func (x *OrderAddReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAddReply.ProtoReflect.Descriptor instead.
func (*OrderAddReply) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{1}
}

type OrderRemoveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderRemoveReq) Reset() {
	*x = OrderRemoveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRemoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRemoveReq) ProtoMessage() {}

func (x *OrderRemoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRemoveReq.ProtoReflect.Descriptor instead.
func (*OrderRemoveReq) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderRemoveReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderRemoveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderRemoveReply) Reset() {
	*x = OrderRemoveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRemoveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRemoveReply) ProtoMessage() {}

func (x *OrderRemoveReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRemoveReply.ProtoReflect.Descriptor instead.
func (*OrderRemoveReply) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{3}
}

type OrderChangeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Text   string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Status int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderChangeReq) Reset() {
	*x = OrderChangeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderChangeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderChangeReq) ProtoMessage() {}

func (x *OrderChangeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderChangeReq.ProtoReflect.Descriptor instead.
func (*OrderChangeReq) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderChangeReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderChangeReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderChangeReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *OrderChangeReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type OrderChangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderChangeReply) Reset() {
	*x = OrderChangeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderChangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderChangeReply) ProtoMessage() {}

func (x *OrderChangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderChangeReply.ProtoReflect.Descriptor instead.
func (*OrderChangeReply) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{5}
}

type ListOrderQueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       *PageInfoReq `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	QueryType  string       `protobuf:"bytes,2,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	QueryParam string       `protobuf:"bytes,3,opt,name=query_param,json=queryParam,proto3" json:"query_param,omitempty"`
}

func (x *ListOrderQueryReq) Reset() {
	*x = ListOrderQueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderQueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderQueryReq) ProtoMessage() {}

func (x *ListOrderQueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderQueryReq.ProtoReflect.Descriptor instead.
func (*ListOrderQueryReq) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrderQueryReq) GetPage() *PageInfoReq {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListOrderQueryReq) GetQueryType() string {
	if x != nil {
		return x.QueryType
	}
	return ""
}

func (x *ListOrderQueryReq) GetQueryParam() string {
	if x != nil {
		return x.QueryParam
	}
	return ""
}

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text        string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Status      int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderInfo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *OrderInfo) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *OrderInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ListOrderQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderInfo   `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Page   *PageInfoReply `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListOrderQueryReply) Reset() {
	*x = ListOrderQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderQueryReply) ProtoMessage() {}

func (x *ListOrderQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderQueryReply.ProtoReflect.Descriptor instead.
func (*ListOrderQueryReply) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{8}
}

func (x *ListOrderQueryReply) GetOrders() []*OrderInfo {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *ListOrderQueryReply) GetPage() *PageInfoReply {
	if x != nil {
		return x.Page
	}
	return nil
}

type OrderQueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderQueryReq) Reset() {
	*x = OrderQueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderQueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderQueryReq) ProtoMessage() {}

func (x *OrderQueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderQueryReq.ProtoReflect.Descriptor instead.
func (*OrderQueryReq) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{9}
}

func (x *OrderQueryReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *OrderInfo `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderQueryReply) Reset() {
	*x = OrderQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderQueryReply) ProtoMessage() {}

func (x *OrderQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderQueryReply.ProtoReflect.Descriptor instead.
func (*OrderQueryReply) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{10}
}

func (x *OrderQueryReply) GetOrder() *OrderInfo {
	if x != nil {
		return x.Order
	}
	return nil
}

type PageInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PageInfoReq) Reset() {
	*x = PageInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoReq) ProtoMessage() {}

func (x *PageInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoReq.ProtoReflect.Descriptor instead.
func (*PageInfoReq) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{11}
}

func (x *PageInfoReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageInfoReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PageInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total    int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PageInfoReply) Reset() {
	*x = PageInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_v1_order_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoReply) ProtoMessage() {}

func (x *PageInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_v1_order_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoReply.ProtoReflect.Descriptor instead.
func (*PageInfoReply) Descriptor() ([]byte, []int) {
	return file_api_order_v1_order_proto_rawDescGZIP(), []int{12}
}

func (x *PageInfoReply) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageInfoReply) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageInfoReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_api_order_v1_order_proto protoreflect.FileDescriptor

var file_api_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x0f, 0x0a,
	0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20,
	0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x12, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x64, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x7e,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x6a,
	0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6f, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2b, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2b,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0f,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x29, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x0b, 0x50, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x5d, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x32, 0xc9, 0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x32, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x5c,
	0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x71, 0x7a, 0x2f, 0x48,
	0x61, 0x70, 0x70, 0x79, 0x2d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_order_v1_order_proto_rawDescOnce sync.Once
	file_api_order_v1_order_proto_rawDescData = file_api_order_v1_order_proto_rawDesc
)

func file_api_order_v1_order_proto_rawDescGZIP() []byte {
	file_api_order_v1_order_proto_rawDescOnce.Do(func() {
		file_api_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_order_v1_order_proto_rawDescData)
	})
	return file_api_order_v1_order_proto_rawDescData
}

var file_api_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_order_v1_order_proto_goTypes = []interface{}{
	(*OrderAddReq)(nil),         // 0: order.v1.OrderAddReq
	(*OrderAddReply)(nil),       // 1: order.v1.OrderAddReply
	(*OrderRemoveReq)(nil),      // 2: order.v1.OrderRemoveReq
	(*OrderRemoveReply)(nil),    // 3: order.v1.OrderRemoveReply
	(*OrderChangeReq)(nil),      // 4: order.v1.OrderChangeReq
	(*OrderChangeReply)(nil),    // 5: order.v1.OrderChangeReply
	(*ListOrderQueryReq)(nil),   // 6: order.v1.ListOrderQueryReq
	(*OrderInfo)(nil),           // 7: order.v1.OrderInfo
	(*ListOrderQueryReply)(nil), // 8: order.v1.ListOrderQueryReply
	(*OrderQueryReq)(nil),       // 9: order.v1.OrderQueryReq
	(*OrderQueryReply)(nil),     // 10: order.v1.OrderQueryReply
	(*PageInfoReq)(nil),         // 11: order.v1.PageInfoReq
	(*PageInfoReply)(nil),       // 12: order.v1.PageInfoReply
}
var file_api_order_v1_order_proto_depIdxs = []int32{
	11, // 0: order.v1.ListOrderQueryReq.page:type_name -> order.v1.PageInfoReq
	7,  // 1: order.v1.ListOrderQueryReply.orders:type_name -> order.v1.OrderInfo
	12, // 2: order.v1.ListOrderQueryReply.page:type_name -> order.v1.PageInfoReply
	7,  // 3: order.v1.OrderQueryReply.order:type_name -> order.v1.OrderInfo
	0,  // 4: order.v1.Order.AddOrder:input_type -> order.v1.OrderAddReq
	2,  // 5: order.v1.Order.RemoveOrder:input_type -> order.v1.OrderRemoveReq
	4,  // 6: order.v1.Order.ChangeOrder:input_type -> order.v1.OrderChangeReq
	6,  // 7: order.v1.Order.ListOrder:input_type -> order.v1.ListOrderQueryReq
	9,  // 8: order.v1.Order.QueryOrder:input_type -> order.v1.OrderQueryReq
	1,  // 9: order.v1.Order.AddOrder:output_type -> order.v1.OrderAddReply
	3,  // 10: order.v1.Order.RemoveOrder:output_type -> order.v1.OrderRemoveReply
	5,  // 11: order.v1.Order.ChangeOrder:output_type -> order.v1.OrderChangeReply
	8,  // 12: order.v1.Order.ListOrder:output_type -> order.v1.ListOrderQueryReply
	10, // 13: order.v1.Order.QueryOrder:output_type -> order.v1.OrderQueryReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_order_v1_order_proto_init() }
func file_api_order_v1_order_proto_init() {
	if File_api_order_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_order_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddReq); i {
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
		file_api_order_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAddReply); i {
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
		file_api_order_v1_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRemoveReq); i {
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
		file_api_order_v1_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRemoveReply); i {
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
		file_api_order_v1_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderChangeReq); i {
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
		file_api_order_v1_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderChangeReply); i {
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
		file_api_order_v1_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderQueryReq); i {
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
		file_api_order_v1_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfo); i {
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
		file_api_order_v1_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderQueryReply); i {
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
		file_api_order_v1_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderQueryReq); i {
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
		file_api_order_v1_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderQueryReply); i {
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
		file_api_order_v1_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoReq); i {
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
		file_api_order_v1_order_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoReply); i {
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
			RawDescriptor: file_api_order_v1_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_order_v1_order_proto_goTypes,
		DependencyIndexes: file_api_order_v1_order_proto_depIdxs,
		MessageInfos:      file_api_order_v1_order_proto_msgTypes,
	}.Build()
	File_api_order_v1_order_proto = out.File
	file_api_order_v1_order_proto_rawDesc = nil
	file_api_order_v1_order_proto_goTypes = nil
	file_api_order_v1_order_proto_depIdxs = nil
}
