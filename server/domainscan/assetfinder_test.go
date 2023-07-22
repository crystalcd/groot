package domainscan

import (
	"testing"

	"github.com/crystal/groot/bean"
)

func TestAssetfinder_run(t *testing.T) {
	type fields struct {
		DomainScan DomainScan
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				DomainScan: DomainScan{
					Config: bean.Config{
						Path: "/Users/byronchen/go/bin/assetfinder",
					},
				},
			},
			args: args{
				domain: "baidu.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Assetfinder{
				DomainScan: tt.fields.DomainScan,
			}
			a.run(tt.args.domain)
		})
	}
}
