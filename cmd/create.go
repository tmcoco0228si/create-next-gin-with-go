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

			f, _ := cmd.Flags().GetString("flameWork")

			switch f {
			case "nextjs": //引数が「nextjs」だった場合、以下の条件判断

				name, _ := cmd.Flags().GetString("name")
				tsFlg, _ := cmd.Flags().GetBool("typescript")

				//実際のNext.jsプロジェクト作成 実行
				err := createTP(name, tsFlg)

				if err != nil {
					fmt.Println("プロジェクト作成に失敗しました")
					return err
				}
				err = option(cmd)

				if err != nil {
					fmt.Println("オプションのインストールに失敗しました。")
					return err
				}

				//Next.jsプロジェクトを上位パスに移動
				err = moveFolder(name)

				if err != nil {
					fmt.Println("Next.jsプロジェクトの移動に失敗しました")
					return err
				}
			case "gin":
				//実際の処理実行
				err := gin()
				if err != nil {
					fmt.Println("ginのインストールに失敗しました。")
					return nil
				}
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
	if len(args) < 4 {
		fmt.Println("引数を入力してください")
	}
	return nil
}

// typescript・nameの有無を判断し、プロジェクト作成実行
func createTP(n string, t bool) error {
	fmt.Println(n)
	// UTF-8エンコードされたバイト列に変換
	nBytes := []byte(n)
	// UTF-8エンコードされたバイト列をstring型に変換
	ns := string(nBytes)
	if t == true {
		err := exec.Command("npx", "create-next-app", ns, " --typescript").Run()
		if err != nil {
			return err
		}
	} else {
		err := exec.Command("npx", "create-next-app", ns).Run()
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
	createCmd.Flags().StringP("flameWork", "f", "", "Framework name")
	createCmd.Flags().StringP("name", "n", "sample-app", "give something name support") //オプション: nameオプション
	createCmd.Flags().BoolP("typescript", "t", false, "typescript support")             //オプション:typescript
	createCmd.Flags().BoolP("eslint", "e", false, "eslint support")                     //オプション: eslint
	createCmd.Flags().BoolP("prettier", "p", false, "prettier support")                 //オプション: prettier
}

// オプションインストール
func option(cmd *cobra.Command) error {
	e, _ := cmd.Flags().GetBool("eslint")
	p, _ := cmd.Flags().GetBool("prettier")
	//eslintのフラグが入力されている場合実行
	if e != false {
		err := exec.Command("yarn", "add", "-D", "eslint", "@typescript-eslint/eslint-plugin", "@typescript-eslint/parser").Run()
		if err != nil {
			return err
		}
		return nil
	}

	if p != false {
		err := exec.Command("yarn", "add", "-D", "prettier", "eslint-config-prettier", "eslint-plugin-prettier").Run()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func gin() error {
	err := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin").Run()
	if err != nil {
		return err
	}
	return nil
}
