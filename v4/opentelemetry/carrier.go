package main

import (
	"go-micro.dev/v4/metadata"
	"go.opentelemetry.io/otel/propagation"
)

type metadataSupplier struct {
	metadata *metadata.Metadata
}

// assert that metadataSupplier implements the TextMapCarrier interface.
var _ propagation.TextMapCarrier = &metadataSupplier{}

func (s *metadataSupplier) Get(key string) string {
	value, ok := s.metadata.Get(key)
	if ok {
		return value
	}
	return ""
}

func (s *metadataSupplier) Set(key string, value string) {
	s.metadata.Set(key, value)
}

func (s *metadataSupplier) Keys() []string {
	out := make([]string, 0, len(*s.metadata))
	for key := range *s.metadata {
		out = append(out, key)
	}
	return out
}
