package meters

import (
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/exp/constraints"
)

type Point[T constraints.Integer | constraints.Float] struct {
	Labels prometheus.Labels
	Amount T
}