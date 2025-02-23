package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-otel/pkg/factorio"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:  "server",
	RunE: serverWorkload,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serverWorkload(cmd *cobra.Command, args []string) error {
	var (
		rconPort     = cmd.Flag("rcon-port").Value.String()
		rconPassword = cmd.Flag("rcon-password").Value.String()
		rconHost     = cmd.Flag("rcon-host").Value.String()
		logger       = log.WithPrefix("server")
	)
	if rconPassword == "" {
		return fmt.Errorf("rcon-password is required")
	}
	if rconPort == "" {
		return fmt.Errorf("rcon-port is required")
	}
	if rconHost == "" {
		return fmt.Errorf("rcon-host is required")
	}

	address := fmt.Sprintf("%s:%s", rconHost, rconPort)
	logger.Info("Connecting to Factorio RCON server", "address", address)
	conn, err := factorio.NewRCONClient(address, rconPassword)
	if err != nil {
		return fmt.Errorf("failed to connect to Factorio RCON server: %w", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close connection", "error", err)
		}
	}()

	// test
	go func() {
		rep, err := conn.Send("/c game.forces[0]")
		if err != nil {
			logger.Error("Failed to send command", "error", err)
		}
		logger.Info("Received response", "response", rep)
	}()

	<-cmd.Context().Done()
	logger.Info("Shutting down")

	return nil
}
