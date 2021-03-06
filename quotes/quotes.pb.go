// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: quotes/quotes.proto

package quotes

import (
	market "github.com/bradenrayhorn/ledger-protos/market"
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

type GetQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol            string              `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	MarketRequestData *market.RequestData `protobuf:"bytes,2,opt,name=MarketRequestData,proto3" json:"MarketRequestData,omitempty"`
}

func (x *GetQuoteRequest) Reset() {
	*x = GetQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quotes_quotes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteRequest) ProtoMessage() {}

func (x *GetQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quotes_quotes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteRequest.ProtoReflect.Descriptor instead.
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return file_quotes_quotes_proto_rawDescGZIP(), []int{0}
}

func (x *GetQuoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetQuoteRequest) GetMarketRequestData() *market.RequestData {
	if x != nil {
		return x.MarketRequestData
	}
	return nil
}

type GetQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *Quote `protobuf:"bytes,1,opt,name=Quote,proto3" json:"Quote,omitempty"`
}

func (x *GetQuoteResponse) Reset() {
	*x = GetQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quotes_quotes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteResponse) ProtoMessage() {}

func (x *GetQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quotes_quotes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteResponse.ProtoReflect.Descriptor instead.
func (*GetQuoteResponse) Descriptor() ([]byte, []int) {
	return file_quotes_quotes_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuoteResponse) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string  `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	AskPrice  float64 `protobuf:"fixed64,2,opt,name=AskPrice,proto3" json:"AskPrice,omitempty"`
	BidPrice  float64 `protobuf:"fixed64,3,opt,name=BidPrice,proto3" json:"BidPrice,omitempty"`
	LastPrice float64 `protobuf:"fixed64,4,opt,name=LastPrice,proto3" json:"LastPrice,omitempty"`
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quotes_quotes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_quotes_quotes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_quotes_quotes_proto_rawDescGZIP(), []int{2}
}

func (x *Quote) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Quote) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *Quote) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *Quote) GetLastPrice() float64 {
	if x != nil {
		return x.LastPrice
	}
	return 0
}

var File_quotes_quotes_proto protoreflect.FileDescriptor

var file_quotes_quotes_proto_rawDesc = []byte{
	0x0a, 0x13, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x1a, 0x13, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x41, 0x0a,
	0x11, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x11, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x75, 0x0a, 0x05, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x73,
	0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x41, 0x73,
	0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x42, 0x69, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x32, 0x50, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x72, 0x61, 0x64, 0x65, 0x6e, 0x72, 0x61, 0x79, 0x68, 0x6f, 0x72, 0x6e, 0x2f, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quotes_quotes_proto_rawDescOnce sync.Once
	file_quotes_quotes_proto_rawDescData = file_quotes_quotes_proto_rawDesc
)

func file_quotes_quotes_proto_rawDescGZIP() []byte {
	file_quotes_quotes_proto_rawDescOnce.Do(func() {
		file_quotes_quotes_proto_rawDescData = protoimpl.X.CompressGZIP(file_quotes_quotes_proto_rawDescData)
	})
	return file_quotes_quotes_proto_rawDescData
}

var file_quotes_quotes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_quotes_quotes_proto_goTypes = []interface{}{
	(*GetQuoteRequest)(nil),    // 0: quotes.GetQuoteRequest
	(*GetQuoteResponse)(nil),   // 1: quotes.GetQuoteResponse
	(*Quote)(nil),              // 2: quotes.Quote
	(*market.RequestData)(nil), // 3: market.RequestData
}
var file_quotes_quotes_proto_depIdxs = []int32{
	3, // 0: quotes.GetQuoteRequest.MarketRequestData:type_name -> market.RequestData
	2, // 1: quotes.GetQuoteResponse.Quote:type_name -> quotes.Quote
	0, // 2: quotes.QuotesService.GetQuote:input_type -> quotes.GetQuoteRequest
	1, // 3: quotes.QuotesService.GetQuote:output_type -> quotes.GetQuoteResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_quotes_quotes_proto_init() }
func file_quotes_quotes_proto_init() {
	if File_quotes_quotes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quotes_quotes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteRequest); i {
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
		file_quotes_quotes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteResponse); i {
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
		file_quotes_quotes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
			RawDescriptor: file_quotes_quotes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quotes_quotes_proto_goTypes,
		DependencyIndexes: file_quotes_quotes_proto_depIdxs,
		MessageInfos:      file_quotes_quotes_proto_msgTypes,
	}.Build()
	File_quotes_quotes_proto = out.File
	file_quotes_quotes_proto_rawDesc = nil
	file_quotes_quotes_proto_goTypes = nil
	file_quotes_quotes_proto_depIdxs = nil
}
