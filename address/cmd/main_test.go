package main

import (
	"testing"

	"github.com/jailtonjunior94/partners/address/internal/services"
)

func TestCepGrpcServerConstruction(t *testing.T) {
	srv := services.NewCepGrpcServer(nil, nil)
	if srv == nil {
		t.Fatal("expected non-nil CepGrpcServer")
	}
}
