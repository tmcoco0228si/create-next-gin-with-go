package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

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
			//バリデーションチェック
			valiCmd(args)

			switch args[0] {
			case "nextjs": //引数が「nextjs」だった場合、以下の条件判断

				name, _ := cmd.Flags().GetString("name")
				tsFlg, _ := cmd.Flags().GetBool("typescript")

				//実際のNext.jsプロジェクト作成 実行
				err := createTP(name, tsFlg)

				if err != nil {
					fmt.Println("プロジェクト作成に失敗しました")
					return err
				}

				//Next.jsプロジェクトを上位パスに移動
				err = moveFolder(name)

				if err != nil {
					fmt.Println("Next.jsプロジェクトの移動に失敗しました")
					return err
				}

				// eslint, _ := cmd.Flags().GetBool("eslint");
				//  prettier, _ := cmd.Flags().GetBool("prettier");
				fmt.Println("Next.js install start")

			case "gin":
				fmt.Println("Gin install start")
			default:
				return errors.New("引数には、「gin」「nextjs」のいずれかを入力してください")
			}
			return nil
		},
	}
	//フラグの作成
	createFlg(createCmd)
	return createCmd
}

// コマンドのバリデーションチェック
func valiCmd(args []string) error {
	// プログラム名を除いた引数の要素数取得
	if len(args) < 1 {
		return errors.New("引数を入力してください")
	}
	return nil
}

// typescript・nameの有無を判断し、プロジェクト作成実行
func createTP(n string, f bool) error {
	if f == true {
		err := exec.Command("npx", "create-next-app", n, " --typescript").Run()
		if err != nil {
			return err
		}
	} else {
		err := exec.Command("npx", "create-next-app", n).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// フォルダ移動
func moveFolder(n string) error {
	if err := os.Rename("./"+n, "../"+n); err != nil {
		return err
	}

	return nil
}

// フラグ作成
func createFlg(createCmd *cobra.Command) {
	// Next.jsに関するフラグの設定
	createCmd.Flags().BoolP("typescript", "t", false, "typescript support")             //オプション:typescript
	createCmd.Flags().BoolP("eslint", "e", false, "eslint support")                     //オプション: eslint
	createCmd.Flags().BoolP("prettier", "p", false, "prettier support")                 //オプション: prettier
	createCmd.Flags().StringP("name", "n", "sample-app", "give something name support") //オプション: nameオプション
}
