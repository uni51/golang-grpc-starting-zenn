// このファイルはprotoc-gen-go-grpcによって生成されました。編集しないでください。
// バージョン:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// ソース: hello.proto

package grpc

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// このコンパイル時アサーションは、この生成ファイルが
// 対象としてコンパイルされているgrpcパッケージと互換性があることを確認します。
// gRPC-Go v1.32.0以降が必要です。
const _ = grpc.SupportPackageIsVersion7

// GreetingServiceClientはGreetingServiceサービスのクライアントAPIです。
//
// ctxの使用やストリーミングRPCの終了に関するセマンティクスについては、https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream を参照してください。
// クライアントが呼び出せるメソッド一覧をインターフェースで定義
type GreetingServiceClient interface {
	// サービスが持つメソッドの定義
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

// クライアント実装の構造体
type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

// リクエストを送るクライアントを作るコンストラクタ
func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

// Helloメソッドの実装
func (c *greetingServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	// サーバーにリクエストを送信し、応答を受け取る
	err := c.cc.Invoke(ctx, "/myapp.GreetingService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServiceServerはGreetingServiceサービスのサーバーAPIです。
// すべての実装はUnimplementedGreetingServiceServerを埋め込む必要があります。
// 将来の互換性のために
type GreetingServiceServer interface {
	// サービスが持つメソッドの定義
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServerは将来の互換性のために埋め込む必要があります。
type UnimplementedGreetingServiceServer struct {
}

// Helloメソッドが実装されていない場合のエラーメッセージを返す
func (UnimplementedGreetingServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServerは、このサービスの将来の互換性を放棄するために埋め込むことができます。
// このインターフェースの使用は推奨されません。GreetingServiceServerにメソッドが追加されると、
// コンパイルエラーが発生するためです。
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

// GreetingServiceサーバーを登録する関数
func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

// Helloメソッドのハンドラー
func _GreetingService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myapp.GreetingService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingService_ServiceDescはGreetingServiceサービスのgrpc.ServiceDescです。
// これはgrpc.RegisterServiceで直接使用することのみを目的としており、
// イントロスペクションや変更（コピーであっても）を目的としたものではありません。
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myapp.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreetingService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
