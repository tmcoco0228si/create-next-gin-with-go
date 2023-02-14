package cmd

import (
	"errors"
	"fmt"

	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

func newCreateCmd(ui *rwi.RWI) *cobra.Command {
	// Next.jsのコマンド実装
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "createコマンドの後に引数やフラグを入力してください",
		Long: `createコマンドの後に引数やフラグを入力してください
					引数には「gin」「nextjs」と打ち込むことでGinのインストール、Next.jsのプロジェクト作成ができます。
					「nextjs」のフラグでは「typescriptのon・off, eslint,prettier」のインストールするかどうか選べます。
					詳しくは --hで詳細を確認してください。
					`,
		//インスタンスを内部関数を使って動的に生成
		RunE: func(cmd *cobra.Command, args []string) error {

			// プログラム名を除いた引数の要素数取得
			if len(args) < 1 {
				return errors.New("引数を入力してください")
			}
			if err := ui.Outputln("createCommand called"); err != nil {
				ui.Outputln("失敗")
				return err
			}
			switch args[0] {
			case "nextjs": //引数が「nextjs」だった場合、以下の条件判断
				if typescript, _ := cmd.Flags().GetBool("typescript"); typescript == true {
					fmt.Println("typescript on")
				} else {
					fmt.Println("typescript off")
				}
				if eslint, _ := cmd.Flags().GetBool("eslint"); eslint == true {
					fmt.Println("eslint on")
				}
				if prettier, _ := cmd.Flags().GetBool("prettier"); prettier == true {
					fmt.Println("prittier on")
				}
				fmt.Println("Next.js install start")
			case "gin":
				fmt.Println("Gin install start")
			default:
				return errors.New("引数には、「gin」「nextjs」のいずれかを入力してください")
			}
			return nil
		},
	}

	// Next.jsに関するフラグの設定
	createCmd.Flags().BoolP("typescript", "t", false, "typescript support") //オプション:typescript
	createCmd.Flags().BoolP("eslint", "e", false, "eslint support")         //オプション: eslint
	createCmd.Flags().BoolP("prettier", "p", false, "prettier support")     //オプション: prettier
	return createCmd
}
