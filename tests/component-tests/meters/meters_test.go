package component_meters_test

import (
	"context"
	"testing"
	"time"

	"github.com/daanv2/go-factorio-prometheus/pkg/data"
	"github.com/daanv2/go-factorio-prometheus/pkg/meters"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func Test_All_Meters(t *testing.T) {
	manager := meters.NewManager(&meters.FakeExecutor{})
	data.Setup(manager)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	manager.Start(ctx)

	// Await manager to do its thing
	<-ctx.Done()

	meters, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	for _, m := range meters {
		require.NotNil(t, m)
		require.NotNil(t, m.Help)
		require.NotNil(t, m.Name)
		require.NotNil(t, m.Type)
		require.Len(t, m.Metric, 1)
	}
}
