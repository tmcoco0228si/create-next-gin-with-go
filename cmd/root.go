package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ルートコマンド
var rootCmd = &cobra.Command{
	Use:   "create-next-gin-with-go_cli",
	Short: "A brief description of your application",
	Long:  `next-ginと記載すれば「next.js」と「gin」のインストールが開始されます`,
	Run: func(cmd *cobra.Command, args []string) {
		// err := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin").Run()
		// err := exec.Command("npx", "create-next-app", "sample-app", " --typescript").Run()
		// if err != nil {
		// 	fmt.Println(err)
		// }
		fmt.Println("rootとしてきてるよ")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.create-next-gin-with-go_cli.yaml)")
	rootCmd.Flags().StringP("next", "-n", "", "next.js作成時のコマンド")
}
