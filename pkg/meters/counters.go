package meters

import (
	"context"

	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/exp/constraints"
)

type Counter[T constraints.Integer | constraints.Float] struct {
	counter *prometheus.CounterVec
	name    string
	scrape  Scrape[T]
	cost    cost.Cost
}

func (g *Counter[T]) Name() string {
	return g.name
}

func (g *Counter[T]) Cost() cost.Cost {
	return g.cost
}

func (g *Counter[T]) SetCost(c cost.Cost) {
	g.cost = c
}

func (g *Counter[T]) Scrape(ctx context.Context, executor Executor) error {
	data, err := g.scrape(ctx, executor)
	if err != nil {
		return err
	}

	for _, p := range data {
		g.counter.With(p.Labels).Add(float64(p.Amount))
	}
	return nil
}

func NewCounter[T constraints.Integer | constraints.Float](name, description string, labels []string, scrape Scrape[T]) *Counter[T] {
	return &Counter[T]{
		name:   name,
		scrape: scrape,
		counter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name:      name,
				Help:      description,
				Namespace: "factorio",
				Subsystem: "game",
			},
			labels,
		),
	}
}
