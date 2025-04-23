package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-prometheus/pkg/data"
	"github.com/daanv2/go-factorio-prometheus/pkg/factorio"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
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
	pf.String("rcon-port", os.Getenv("RCON_PORT"), "The port to connect to the Factorio RCON server")
	pf.String("rcon-password", os.Getenv("RCON_PASSWORD"), "The password to connect to the Factorio RCON server")
	pf.String("rcon-host", os.Getenv("RCON_HOST"), "The host to connect to the Factorio RCON server")
	pf.String("prometheus-address", os.Getenv("PROMETHEUS_ADDRESS"), "The address for prometheus to reach this instance")
}

func serverWorkload(cmd *cobra.Command, args []string) error {
	var (
		rconPort     = cmd.Flag("rcon-port").Value.String()
		rconPassword = cmd.Flag("rcon-password").Value.String()
		rconHost     = cmd.Flag("rcon-host").Value.String()
		promaddress  = cmd.Flag("prometheus-address").Value.String()
		logger       = log.WithPrefix("server")
	)
	if rconPort == "" || rconPassword == "" || rconHost == "" {
		return errors.New("rcon-port, rcon-password, and rcon-host are required")
	}
	if promaddress == "" {
		return errors.New("prometheus-address cannot be empty")
	}

	address := fmt.Sprintf("%s:%s", rconHost, rconPort)
	logger.Info("Connecting to Factorio RCON server", "address", address)
	conn, err := factorio.NewRCONClient(address, rconPassword)
	if err != nil {
		return fmt.Errorf("failed to connect to Factorio RCON server: %w", err)
	}

	manager := meters.NewManager(conn)
	data.Setup(manager)

	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close connection", "error", err)
		}
	}()

	go manager.Start(cmd.Context())

	// Start the Prometheus server
	go func() {
		log.Info("Starting Prometheus server", "address", promaddress + "/metrics")
		http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("Serving metrics")
			promhttp.Handler().ServeHTTP(w, r)
		}))
		err := http.ListenAndServe(promaddress, nil)
		if err != nil {
			logger.Error("Failed to start Prometheus server", "error", err)
		}
	}()

	<-cmd.Context().Done()
	logger.Info("Shutting down")

	return nil
}
