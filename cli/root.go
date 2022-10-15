package cli

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "pinger",
		Short: "pinger",
		Long:  `pinger service.`,
	}
)

func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "sepehrsoh", "sepehrxsohrabpour")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
}
