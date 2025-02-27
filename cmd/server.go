package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-otel/pkg/factorio"
	"github.com/daanv2/go-factorio-otel/pkg/meters"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:  "server",
	RunE: serverWorkload,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if cmd.Flag("rcon-port").Value.String() == "" {
			err = errors.Join(err, errors.New("rcon-password is required"))
		}
		if cmd.Flag("rcon-password").Value.String() == "" {
			err = errors.Join(err, errors.New("rcon-password is required"))
		}
		if cmd.Flag("rcon-host").Value.String() == "" {
			err = errors.Join(err, errors.New("rcon-host is required"))
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	pf := serverCmd.PersistentFlags()
	pf.String("rcon-port", "8090", "The port to connect to the Factorio RCON server")
	pf.String("rcon-password", "", "The password to connect to the Factorio RCON server")
	pf.String("rcon-host", "localhost", "The host to connect to the Factorio RCON server")
}

func serverWorkload(cmd *cobra.Command, args []string) error {
	var (
		rconPort     = cmd.Flag("rcon-port").Value.String()
		rconPassword = cmd.Flag("rcon-password").Value.String()
		rconHost     = cmd.Flag("rcon-host").Value.String()
		logger       = log.WithPrefix("server")
	)
	if rconPort == "" || rconPassword == "" || rconHost == "" {
		return errors.New("rcon-port, rcon-password, and rcon-host are required")
	}

	address := fmt.Sprintf("%s:%s", rconHost, rconPort)
	logger.Info("Connecting to Factorio RCON server", "address", address)
	conn, err := factorio.NewRCONClient(address, rconPassword)
	if err != nil {
		return fmt.Errorf("failed to connect to Factorio RCON server: %w", err)
	}

	manager := meters.NewManager(conn)
	meters.PlayerMeters(manager)
	meters.PlanetsMeters(manager)
	meters.ForcesMeters(manager)
	meters.ResearchMeters(manager)
	meters.LogisticsMeters(manager)

	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close connection", "error", err)
		}
	}()

	go manager.Start(cmd.Context())

	// Start the Prometheus server
	go func() {
		log.Info("Starting Prometheus server", "address", ":2112")
		http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("Serving metrics")
			promhttp.Handler().ServeHTTP(w, r)
		}))
		err := http.ListenAndServe(":2112", nil)
		if err != nil {
			logger.Error("Failed to start Prometheus server", "error", err)
		}
	}()

	<-cmd.Context().Done()
	logger.Info("Shutting down")

	return nil
}
