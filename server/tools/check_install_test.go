package tools

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
func TestGoCheck(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "check install",
			want: false,
		},
		{
			name: "check install2",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoCheck(); got != tt.want {
				t.Errorf("GoCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHomeDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test",
			want: "/Users/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHomeDir(); got != tt.want {
				t.Errorf("getHomeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstallGo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InstallGo()
		})
	}
}

