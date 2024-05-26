// protoのバージョンの宣言

// このファイルはprotoc-gen-goによって生成されました。編集しないでください。
// バージョン:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// ソース: hello.proto

// packageの宣言

package grpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// この生成されたコードが十分に最新であることを確認する。
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// runtime/protoimplが十分に最新であることを確認する。
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 型の定義
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// リクエストに含まれる名前
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

// HelloRequestのリセットメソッド
func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

// HelloRequestの文字列表現を返すメソッド
func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// HelloRequestがプロトコルメッセージであることを示すメソッド
func (*HelloRequest) ProtoMessage() {}

// HelloRequestのプロトコル反射メソッド
func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// 非推奨: HelloRequestのディスクリプタを返すメソッド
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

// HelloRequestのNameフィールドを取得するメソッド
func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 型の定義
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// レスポンスに含まれるメッセージ
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

// HelloResponseのリセットメソッド
func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

// HelloResponseの文字列表現を返すメソッド
func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// HelloResponseがプロトコルメッセージであることを示すメソッド
func (*HelloResponse) ProtoMessage() {}

// HelloResponseのプロトコル反射メソッド
func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// 非推奨: HelloResponseのディスクリプタを返すメソッド
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

// HelloResponseのMessageフィールドを取得するメソッド
func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ファイルディスクリプタ
var File_hello_proto protoreflect.FileDescriptor

// protoファイルのバイナリデータ
var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x79, 0x61, 0x70, 0x70, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x45, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x13, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

// protoファイルのバイナリデータを圧縮して取得する関数
var file_hello_proto_rawDescOnce sync.Once
var file_hello_proto_rawDescData = file_hello_proto_rawDesc

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

// メッセージタイプの情報を保持する変数
var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)

// Goタイプの情報を保持する変数
var file_hello_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: myapp.HelloRequest
	(*HelloResponse)(nil), // 1: myapp.HelloResponse
}

// ディペンデンシーインデックスの情報を保持する変数
var file_hello_proto_depIdxs = []int32{
	0, // 0: myapp.GreetingService.Hello:input_type -> myapp.HelloRequest
	1, // 1: myapp.GreetingService.Hello:output_type -> myapp.HelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

// 初期化関数
func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
	// protoファイルの型ビルダー
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	// ファイルディスクリプタを設定
	File_hello_proto = out.File
	// データをクリア
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}
