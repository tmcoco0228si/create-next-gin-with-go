package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var option string

// ルートコマンド
var rootCmd = &cobra.Command{
	Use:   "nextjs-gin",
	Short: "A brief description of your application",
	Long:  `next-ginと記載すれば「next.js」と「gin」のインストールが開始されます`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//StringPと異なり、入力されたフラグを第一引数で直接使用することができる。
	rootCmd.PersistentFlags().StringVar(&option, "option", "", "option")
}
