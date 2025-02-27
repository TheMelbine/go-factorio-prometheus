package meters

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/exp/constraints"
)

type Gauge[T constraints.Integer | constraints.Float] struct {
	counter *prometheus.GaugeVec
	name    string
	scrape  Scrape[T]
}

func (g *Gauge[T]) Name() string {
	return g.name
}

func (g *Gauge[T]) Scrape(ctx context.Context, executor Executor) error {
	logger := log.WithPrefix("gauge-" + g.name)
	logger.Debug("Scraping gauge")
	data, err := g.scrape(ctx, executor)
	logger.Debug("Scraped gauge", "error", err, "data", data)
	if err != nil {
		return err
	}

	for _, p := range data {
		g.counter.With(p.Labels).Set(float64(p.Amount))
	}
	return nil
}

func NewGauge[T constraints.Integer | constraints.Float](name, description string, labels []string, scrape Scrape[T]) *Gauge[T] {
	return &Gauge[T]{
		name:   name,
		scrape: scrape,
		counter: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name:      name,
				Help:      description,
				Namespace: "factorio",
				Subsystem: "game",
			},
			labels,
		),
	}
}

func (m *Manager) NewGaugeInt64(name, description string, labels []string, scrape Scrape[int64]) {
	m.AddMeter(NewGauge[int64](name, description, labels, scrape))
}

func (m *Manager) NewGaugeFloat64(name, description string, labels []string, scrape Scrape[float64]) {
	m.AddMeter(NewGauge[float64](name, description, labels, scrape))
}
