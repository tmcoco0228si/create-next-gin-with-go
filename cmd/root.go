package cmd

import (
	"fmt"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

var option string

func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {

	// ルートコマンド
	rootCmd := &cobra.Command{
		Use:   "nextjs-gin",
		Short: "A brief description of your application",
		Long:  `next-ginと記載すれば「next.js」と「gin」のインストールが開始されます`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create")
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)            //コマンドの引数を設定
	rootCmd.SetIn(ui.Reader())       //入力データのソースを設定
	rootCmd.SetOut(ui.ErrorWriter()) //使用法メッセージの送信先を設定
	rootCmd.SetErr(ui.ErrorWriter()) //エラーメッセージの送信先を設定
	rootCmd.AddCommand(newCreateCmd(ui))//この親コマンドに 1 つまたは複数のコマンドを追加します。

	return rootCmd
}

func Execute(ui *rwi.RWI, args []string) exitcode.ExitCode {
	if err := newRootCmd(ui, args).Execute(); err != nil {
		return exitcode.Abnormal
	}
	return exitcode.Normal
}

// func init() {
// 	//StringPと異なり、入力されたフラグを第一引数で直接使用することができる。
// 	rootCmd.PersistentFlags().StringVar(&option, "option", "", "option")
// }
