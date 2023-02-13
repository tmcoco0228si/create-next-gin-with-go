package cmd

import (
	"github.com/spf13/cobra"
)

// nextjsCmd represents the nextjs command
var nextjsCmd = &cobra.Command{
	Use:   "nextjs",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(nextjsCmd)
	// nextjsCmd.PersistentFlags().String("foo", "", "A help for foo")
	// nextjsCmd.Flags().StringP("", "t", false, "Help message for toggle")
}
