package meters

import (
	"context"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/exp/constraints"
)

type Gauge[T constraints.Integer | constraints.Float] struct {
	counter *prometheus.GaugeVec
	name    string
	scrape  Scrape[T]
	cost    cost.Cost
}

func (g *Gauge[T]) Name() string {
	return g.name
}

func (g *Gauge[T]) Cost() cost.Cost {
	return g.cost
}

func (g *Gauge[T]) SetCost(c cost.Cost) {
	g.cost = c
}

func (g *Gauge[T]) Scrape(ctx context.Context, executor Executor) error {
	data, err := g.scrape(ctx, executor)
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
