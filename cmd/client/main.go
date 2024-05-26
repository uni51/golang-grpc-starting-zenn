package main

import (
	// 必要なパッケージのインポート
	"bufio"
	"context"
	"fmt"
	"log"
	hellopb "mygrpc/pkg/grpc"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// 標準入力から文字列を受け取るスキャナの変数
	scanner *bufio.Scanner
	// gRPCクライアントの変数
	client hellopb.GreetingServiceClient
)

func main() {
	// gRPCクライアントの開始メッセージを表示
	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		// セキュリティを無効にして接続を確立
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// ブロッキングダイヤル（接続が確立するまでブロックする）
		grpc.WithBlock(),
	)
	if err != nil {
		// 接続に失敗した場合はエラーメッセージを表示して終了
		log.Fatal("Connection failed.")
		return
	}
	// プログラム終了時にコネクションを閉じる
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = hellopb.NewGreetingServiceClient(conn)

	// メインループ
	for {
		// ユーザーに選択肢を提示
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		// ユーザーの入力をスキャン
		scanner.Scan()
		in := scanner.Text()

		// 入力に応じた処理を実行
		switch in {
		case "1":
			// リクエストを送信する関数を呼び出し
			Hello()

		case "2":
			// 終了メッセージを表示してループを抜ける
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

// リクエストを送信する関数
func Hello() {
	// ユーザーに名前の入力を促す
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	// リクエストに使うHelloRequest型の生成
	req := &hellopb.HelloRequest{
		Name: name,
	}
	// Helloメソッドの実行 -> HelloResponse型のレスポンスresを入手
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		// エラーが発生した場合はエラーメッセージを表示
		fmt.Println(err)
	} else {
		// 正常にレスポンスを受け取った場合はメッセージを表示
		fmt.Println(res.GetMessage())
	}
}
