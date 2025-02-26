package meters

import (
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/exp/constraints"
)

type Point[T constraints.Integer | constraints.Float] struct {
	Labels []attribute.KeyValue
	Amount T
}