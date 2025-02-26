package meters

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type Executor interface {
	Execute(cmd string) (string, error)
}

type CustomMeter interface {
	Name() string
	Scrape(ctx context.Context, executor Executor) error
	Observe(ctx context.Context, observer metric.Observer) error
	Instrument() metric.Observable
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

func (m *Manager) AddMeter(s CustomMeter) error {
	m.meters = append(m.meters, s)
	_, err := m.meter.RegisterCallback(s.Observe, s.Instrument())
	return err
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

func (m *Manager) loop(ctx context.Context) {
	// Create a ticker that ticks every 100 ms (10 commands per second).
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for _, s := range m.meters {
		err := s.Scrape(ctx, m.executor)
		if err != nil {
			log.Error("failed to scrape", "error", err, "meter", s.Name())
			continue
		}

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return
		}
	}
}
