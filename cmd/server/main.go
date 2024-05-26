package main

import (
	// 必要なパッケージのインポート
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	// 生成されたgRPCコードのパッケージをインポート
	hellopb "mygrpc/pkg/grpc"

	// gRPC関連のパッケージをインポート
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// サーバーの構造体を定義
type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

// サーバーのインスタンスを作成する関数
func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// 1. 8080番ポートのListenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// 3. gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// 4. サーバーリフレクションの設定
	reflection.Register(s)

	// 5. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 6. Ctrl+Cが入力されたらGraceful shutdownされるようにする
	// シグナルを受け取るためのチャネルを作成する。
	// os.Signal型のチャネルquitをバッファサイズ1で作成。
	quit := make(chan os.Signal, 1)

	// os.Interruptシグナルを受け取ったら、quitチャネルに通知するように設定。
	// ここでos.Interruptは、一般的にCtrl+Cによって発生するシグナル。
	signal.Notify(quit, os.Interrupt)

	// quitチャネルからの信号を待機。
	// ここでブロックして、os.Interruptシグナルが来るまで待つ。
	<-quit

	// シグナルを受け取ったことをログに出力。
	log.Println("stopping gRPC server...")

	// gRPCサーバーをGracefulに停止する。
	// 現在処理中のRPCが終了するのを待ってから、サーバーを停止する。
	s.GracefulStop()
}

// Helloメソッドの実装
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
