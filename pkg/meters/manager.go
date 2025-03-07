package meters

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters/cost"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"golang.org/x/exp/constraints"
)

type Scrape[T constraints.Integer | constraints.Float] func(ctx context.Context, executor Executor) ([]Point[T], error)

type Executor interface {
	Execute(cmd string) (string, error)
}

type CustomMeter interface {
	Name() string
	Scrape(ctx context.Context, executor Executor) error
	Cost() cost.Cost
}

type Manager struct {
	executor Executor
	meter    metric.Meter
	meters   []CustomMeter
}

func NewManager(executor Executor) *Manager {
	meter := otel.Meter("factorio-otel")
	meters := make([]CustomMeter, 0)

	return &Manager{
		executor,
		meter,
		meters,
	}
}

func (m *Manager) AddMeter(s CustomMeter) {
	log.WithPrefix("meters").Debug("adding meter", "name", s.Name(), "cost", s.Cost())
	m.meters = append(m.meters, s)
}

func (m *Manager) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			m.loop(ctx)
		}
	}
}

func (m *Manager) LoopOnce(ctx context.Context) error {
	for _, s := range m.meters {
		err := s.Scrape(ctx, m.executor)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) loop(ctx context.Context) {
	// Create a ticker that ticks every 100 ms (10 commands per second).
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for _, s := range m.meters {
		err := s.Scrape(ctx, m.executor)
		if err != nil {
			log.WithPrefix(s.Name()).Error("failed to scrape", "error", err)
			continue
		}

		if s.Cost() != cost.NONE {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return
			}
		}
	}

	// Sleep 5 or exit if the context is done.
	select {
	case <-time.After(5 * time.Second):
	case <-ctx.Done():
	}
}


func (m *Manager) NewGaugeInt64(name, description string, labels []string, scrape Scrape[int64]) (*Gauge[int64]) {
	g := NewGauge(name, description, labels, scrape)
	m.AddMeter(g)
	return g
}

func (m *Manager) NewGaugeFloat64(name, description string, labels []string, scrape Scrape[float64]) (*Gauge[float64]) {
	g := NewGauge(name, description, labels, scrape)
	m.AddMeter(g)
	return g
}


func (m *Manager) NewCounterInt64(name, description string, labels []string, scrape Scrape[int64]) (*Counter[int64]) {
	g := NewCounter(name, description, labels, scrape)
	m.AddMeter(g)
	return g
}

func (m *Manager) NewCounterFloat64(name, description string, labels []string, scrape Scrape[float64]) (*Counter[float64]) {
	g := NewCounter(name, description, labels, scrape)
	m.AddMeter(g)
	return g
}