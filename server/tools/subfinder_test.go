package tools

import (
	"testing"
)




func TestSubfinder_Do(t *testing.T) {
	type fields struct {
		Config Config
		Param  Param
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "first",
			fields: fields{
				Config: Config{
					Path : "/Users/byronchen/go/bin/subfinder",
				},
				Param: Param{
					Target : "slack.com,hackerone.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subfinder{
				Config: tt.fields.Config,
				Param:  tt.fields.Param,
			}
			s.Do()
		})
	}
}
