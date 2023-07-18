package domainscan

import (
	"testing"

	"github.com/crystal/groot/bean"
)

func TestAssetfinder_Do(t *testing.T) {
	type fields struct {
		S bean.DomainScan
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assetfinder{
				S: tt.fields.S,
			}
			a.Do()
		})
	}
}
