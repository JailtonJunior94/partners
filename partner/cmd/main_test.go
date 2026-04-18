package main

import (
	"testing"

	"github.com/jailtonjunior94/partners/partner/internal/domain/entities"
)

func TestPartnerEntityConstruction(t *testing.T) {
	p := entities.NewPartner("Loja A", "Joao", "12345678000100")
	if p == nil {
		t.Fatal("expected non-nil Partner")
	}
	if p.TradingName != "Loja A" {
		t.Errorf("TradingName = %q, want %q", p.TradingName, "Loja A")
	}
}

func TestAddressEntityConstruction(t *testing.T) {
	a := entities.NewAddress("01001000", "Praca da Se", "Se", "Sao Paulo", "SP")
	if a == nil {
		t.Fatal("expected non-nil Address")
	}
	a.AddLocation(-23.55, -46.63)
	if a.Location == nil {
		t.Fatal("expected non-nil Location after AddLocation")
	}
}
