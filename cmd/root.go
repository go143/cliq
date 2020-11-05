package cmd

import (
	u "github.com/go143/cliq/utils"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cliq",
		Short: "CLI interface for your next big thing",
		Long:  `cliq is a quickstart for building awesome CLIs`,
		Run:   RUN(rootRun),
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// Sample keys in config
const (
	PIPELINE = "pipeline"
	ENV      = "env"
	NOCOLOR  = "no-color"
	CONFIG   = "config"
)

// DEFAULT values for keys in config
var DEFAULT = map[string]interface{}{
	ENV:     "dev",
	NOCOLOR: false,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, CONFIG, "", "config file (default is $HOME/.cliqConfig.yaml)")

	rootCmd.PersistentFlags().StringP(ENV, "e", "", "Environment: dev/prod")
	rootCmd.PersistentFlags().StringP(PIPELINE, "p", "", "Pipeline name, should be a valid name")
	rootCmd.PersistentFlags().Bool(NOCOLOR, false, "Set this to disable colored output")

	/**This step is necessary for config to work with flags**/
	viper.BindPFlag(ENV, rootCmd.PersistentFlags().Lookup(ENV))
	viper.BindPFlag(PIPELINE, rootCmd.PersistentFlags().Lookup(PIPELINE))
	viper.BindPFlag(NOCOLOR, rootCmd.PersistentFlags().Lookup(NOCOLOR))
	viper.SetDefault(ENV, DEFAULT[ENV])
	viper.SetDefault(NOCOLOR, DEFAULT[NOCOLOR])
}

func rootRun() {
	u.Print(`Usage: 
	> cliq --help
	
     CCCC  L    III   QQ
    C      L     I   Q  Q
    C	   L     I   Q  Q
    C      L     I   Q  Q
     CCCC  LLL  III   QQ Q`)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			ERR(err)
		}
		// Search config in home directory with name ".c360config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cliqConfig")
	}
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if CONF(NOCOLOR) == true {
		u.DisableColor()
	}
	if err == nil {
		u.Green("Using config file:", viper.ConfigFileUsed())
	}
}
