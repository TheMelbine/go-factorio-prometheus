package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-factorio-otel",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	defer cancel()
	rootCmd.SetContext(ctx)

	go func() {
		<- ctx.Done()
		log.Info("Shutdown received")
	}()

	defer func() {
		if e := recover(); e != nil {
			log.Fatal("uncaught error", "error", e)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal("error during executing command", "error", err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-factorio-otel.yaml)")

	pf := rootCmd.PersistentFlags()
	pf.String("rcon-port", "25575", "The port to connect to the Factorio RCON server")
	pf.String("rcon-password", "", "The password to connect to the Factorio RCON server")
	pf.String("rcon-host", "localhost", "The host to connect to the Factorio RCON server")
	pf.String("otel-collector", "localhost:4317", "The host and port of the OpenTelemetry collector")
	pf.String("otel-service-name", "factorio-otel", "The service name to use for OpenTelemetry")
}