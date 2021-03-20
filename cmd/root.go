package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const envPrefix = "ottosocial"

var consumerKey, consumerSecret, token, tokenSecret, logpath string

func rootCmd(v *viper.Viper) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ottosocial",
		Short: "Twitter in the shell",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bindFlagToConfig(cmd, v)
		},
	}

	rootCmd.AddCommand(csvCmd(InitLoggerFile(logpath)))
	rootCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Your Twitter Consumer Key (required)")
	rootCmd.PersistentFlags().StringVarP(&consumerSecret, "secret", "s", "", "Your Twitter Consumer Secret (required)")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Your Twitter Access Token (required)")
	rootCmd.PersistentFlags().StringVarP(&tokenSecret, "token-secret", "j", "", "Your Twitter Access Token Secret (required)")
	rootCmd.Flags().StringVarP(&logpath, "logpath", "l", "", "path for logs")

	return rootCmd
}

func Execute() {
	if err := rootCmd(initConfig()).Execute(); err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
}

func initConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(filepath.Join(xdg.ConfigHome, "ottosocial"))
	v.AddConfigPath(".")
	v.SetConfigName("config")

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	v.ReadInConfig()

	return v
}

func bindFlagToConfig(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.Replace(f.Name, "-", "_", -1)
			v.BindEnv(f.Name, strings.ToUpper(fmt.Sprintf("%s_%s", envPrefix, envVarSuffix)))
		}
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
