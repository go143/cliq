package cmd

import (
	"fmt"
	"os"

	u "github.com/go143/cliq/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RUN wraps the run for each command
func RUN(f func()) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		runnerCommons()
		f()
	}
}

// Add this to each runner
func runnerCommons() {}

// ERR is called before exiting in case of errors
func ERR(msg interface{}) {
	u.Red("Error:", msg)
	os.Exit(1)
}

// REQUIRED checks if the required config field is present
func REQUIRED(f string) {
	if CONF(f).(string) == "" {
		ERR("Required config parameter missing: " + f + fmt.Sprintf("\n(set it with flag --%s)", f))
	}
}

// CONF Returns corresponding value from viper
func CONF(key string) interface{} {
	return viper.Get(key)
}
