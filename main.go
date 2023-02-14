package main

/*
	gocli/rwi 標準入出力をコンテキスト情報として格納する構造体を提供
	gocli/exitcode パッケージは CLI 終了時の終了コードを定義する。
*/

import (
	"os"

	"github.com/goark/gocli/rwi"
	"github.com/tmcoco0228si/create-next-gin-with-go_cli/cmd"
)

func main() {
	//コマンド実行
	// cmd.Execute()
	cmd.Execute(rwi.New(
		rwi.WithReader(os.Stdin),       //標準入力
		rwi.WithWriter(os.Stdout),      //標準出力
		rwi.WithErrorWriter(os.Stderr), //標準エラー
	),//プログラム名から始まるコマンドライン引数格納
		os.Args[1:],
	).Exit()
}
