/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/daanv2/go-factorio-prometheus/pkg/data"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/spf13/cobra"
)

// genDocsCmd represents the genDocs command
var genDocsCmd = &cobra.Command{
	Use:   "gen-docs",
	Short: `Generate markdown documents from the applications data`,
	RunE:  GenerateDocs,
}

func init() {
	rootCmd.AddCommand(genDocsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genDocsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genDocsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GenerateDocs(cmd *cobra.Command, args []string) error {
	manager := meters.NewManager(&meters.FakeExecutor{})
	data.Setup(manager)
	err := manager.LoopOnce(cmd.Context())
	if err != nil {
		return fmt.Errorf("error running once over all metrics: %w", err)
	}

	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return fmt.Errorf("error grabbing metrics: %w", err)
	}
	slices.SortFunc(metrics, func(a, b *io_prometheus_client.MetricFamily) int {
		return strings.Compare(*a.Name, *b.Name)
	})

	tmpl, err := template.New("markdown").Parse(metrics_table)
	if err != nil {
		return fmt.Errorf("error making markdown table template: %w", err)
	}

	// Create or truncate the metrics.md file
	file, err := os.Create(filepath.Clean("./doc/metrics.md"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return tmpl.Execute(file, metrics)
}

const metrics_table = `# Metrics

This document lists all metrics exposed by the application.

| Name | Help | Type |
|------|------|------|
{{range .}}| {{if .Name}}{{.Name}}{{else}}-{{end}} | {{if .Help}}{{.Help}}{{else}}-{{end}} | {{if .Type}}{{.Type.String}}{{else}}-{{end}} |
{{end}}
`
