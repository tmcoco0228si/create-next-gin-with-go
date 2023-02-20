package cmd

import (
	"errors"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {

	// ルートコマンド
	rootCmd := &cobra.Command{
		Use:   "nextjs-gin",
		Short: "A brief description of your application",
		Long:  `next-ginと記載すれば「next.js」と「gin」のインストールが開始されます`,
		RunE: func(cmd *cobra.Command, args []string) error {
			//与えられたテキストをフォーマットし、エラーを返す。
			return errors.New("サブコマンド「create」を入力してください")
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)                //コマンドの引数を設定
	rootCmd.SetIn(ui.Reader())           //入力データのソースを設定
	rootCmd.SetOut(ui.ErrorWriter())     //使用法メッセージの送信先を設定
	rootCmd.SetErr(ui.ErrorWriter())     //エラーメッセージの送信先を設定
	rootCmd.AddCommand(newCreateCmd(ui)) //サブコマンド追加

	return rootCmd
}

// ルートコマンド作成・実行
func Execute(ui *rwi.RWI, args []string) exitcode.ExitCode {
	if err := newRootCmd(ui, args).Execute(); err != nil {
		return exitcode.Abnormal //OS終了コード "異常"
	}
	return exitcode.Normal //OSの終了コード "正常"
}
