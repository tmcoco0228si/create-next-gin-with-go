package cmd

import (
	"fmt"

	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

func newCreateCmd(ui *rwi.RWI) *cobra.Command {
	// Next.jsのコマンド実装
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  ``,
		//インスタンスを内部関数を使って動的に生成
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ui.Outputln("createCommand called"); err != nil {
				ui.Outputln("失敗")
				return err
			}
			switch args[0] {
			case "nextjs":
				if typescript, _ := cmd.Flags().GetBool("typescript"); typescript == true {
					// err := exec.Command("npx", "create-next-app", "sample-app", " --typescript").Run()
					fmt.Println("typescript on")
				} else {
					// err := exec.Command("npx", "create-next-app", "sample-app").Run()
					fmt.Println("typescript off")
				}
				if eslint, _ := cmd.Flags().GetBool("eslint"); eslint == true {
					//eslint
					fmt.Println("eslint on")
					// yarn add -D eslint @typescript-eslint/eslint-plugin @typescript-eslint/parser
				}
				if prettier, _ := cmd.Flags().GetBool("prettier"); prettier == true {
					//prettier
					fmt.Println("prittier on")
					// prittier
					// yarn add -D prettier eslint-config-prettier eslint-plugin-prettier			}
				}
				fmt.Println("Next.js install start")
			case "gin":
				fmt.Println("Gin install start")
			// err := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin").Run()
			default:
				fmt.Println("引数に[nextjs]か[gin]と入力してください")
			}
			return nil
		},
	}
	//フラグの値を変数にバインド
	// createCmd.Flags().StringVarP(&nextjs, "nextjs", "n", "", "ネクスト")
	createCmd.Flags().String("nextjs", "n", "nextjs support")
	createCmd.Flags().String("gin", "n", "gin support")
	createCmd.Flags().BoolP("typescript", "t", false, "typescript support")
	createCmd.Flags().BoolP("eslint", "e", false, "eslint support")
	createCmd.Flags().BoolP("prettier", "p", false, "prettier support")
	return createCmd
}

// func init() {
// 	// rootCmd.AddCommand(createCmd)
// 	//フラグの値を変数にバインド
// 	// createCmd.Flags().StringVarP(&nextjs, "nextjs", "n", "", "ネクスト")
// 	createCmd.Flags().String("nextjs", "n", "nextjs support")
// 	createCmd.Flags().String("gin", "n", "gin support")
// 	createCmd.Flags().BoolP("typescript", "t", false, "typescript support")
// 	createCmd.Flags().BoolP("eslint", "e", false, "eslint support")
// 	createCmd.Flags().BoolP("prettier", "p", false, "prettier support")
// }
