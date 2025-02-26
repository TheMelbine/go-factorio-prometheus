package meters

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/metric"
	"golang.org/x/exp/constraints"
)

type Scrape[T constraints.Integer | constraints.Float] func(ctx context.Context, executor Executor) ([]Point[T], error)

type Gauge struct {
	name   string
	meter  metric.Int64ObservableGauge
	data   []Point[int64]
	lock   sync.Mutex
	scrape Scrape[int64]
}

func (g *Gauge) Name() string {
	return g.name
}

func (g *Gauge) Scrape(ctx context.Context, executor Executor) error {
	data, err := g.scrape(ctx, executor)
	if err != nil {
		return err
	}

	g.lock.Lock()
	defer g.lock.Unlock()

	g.data = data
	return nil
}

func (g *Gauge) Observe(ctx context.Context, observer metric.Observer) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	for _, point := range g.data {
		observer.ObserveInt64(g.meter, point.Amount, metric.WithAttributes(point.Labels...))
	}

	return nil
}

func (g *Gauge) Instrument() metric.Observable {
	return g.meter
}

func (m *Manager) NewGaugeInt64(name, description, unit string, scrape func(ctx context.Context, executor Executor) ([]Point[int64], error)) error {
	gauge, err := m.meter.Int64ObservableGauge(name,
		metric.WithDescription(description),
		metric.WithUnit(unit),
	)
	if err != nil {
		return err
	}

	return m.AddMeter(&Gauge{meter: gauge})
}
