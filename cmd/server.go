package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-otel/internal/setup"
	"github.com/daanv2/go-factorio-otel/pkg/factorio"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
		if cmd.Flag("otel-collector").Value.String() == "" {
			err = errors.Join(err, errors.New("otel-collector is required"))
		}
		if cmd.Flag("otel-service-name").Value.String() == "" {
			err = errors.Join(err, errors.New("otel-service-name is required"))
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	pf := serverCmd.PersistentFlags()
	pf.String("rcon-port", "25575", "The port to connect to the Factorio RCON server")
	pf.String("rcon-password", "", "The password to connect to the Factorio RCON server")
	pf.String("rcon-host", "localhost", "The host to connect to the Factorio RCON server")
	pf.String("otel-collector", "localhost:4317", "The host and port of the OpenTelemetry collector")
	// pf.String("otel-service-name", "factorio-otel", "The service name to use for OpenTelemetry")
}

func serverWorkload(cmd *cobra.Command, args []string) error {
	var (
		rconPort     = cmd.Flag("rcon-port").Value.String()
		rconPassword = cmd.Flag("rcon-password").Value.String()
		rconHost     = cmd.Flag("rcon-host").Value.String()
		otelCollector = cmd.Flag("otel-collector").Value.String()
		otelServiceName = cmd.Flag("otel-service-name").Value.String()
		logger       = log.WithPrefix("server")
	)
	if rconPort == "" || rconPassword == "" || rconHost == "" {
		return errors.New("rcon-port, rcon-password, and rcon-host are required")
	}
	if otelCollector == "" || otelServiceName == "" {
		return errors.New("otel-collector and otel-service-name are required")
	}


	address := fmt.Sprintf("%s:%s", rconHost, rconPort)
	logger.Info("Connecting to Factorio RCON server", "address", address)
	conn, err := factorio.NewRCONClient(address, rconPassword)
	if err != nil {
		return fmt.Errorf("failed to connect to Factorio RCON server: %w", err)
	}

	// Set up OpenTelemetry.
	otelShutdown, err := setup.OTelSDK(cmd.Context(), otelCollector)
	if err != nil {
		return err
	}
	var errs error
	// Handle shutdown properly so nothing leaks.
	defer func() {
		errs = errors.Join(errs, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return cmd.Context() },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errs = errors.Join(errs, err)
		}
	}()

	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close connection", "error", err)
			errs = errors.Join(errs, err)
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

	return errs
}


func newHTTPHandler() http.Handler {
	// Add HTTP instrumentation for the whole server.
	mux := http.NewServeMux()
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}