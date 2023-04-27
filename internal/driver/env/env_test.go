package env_test

import (
	"reflect"
	"testing"

	"github.com/takokun778/gotagnews/internal/driver/env"
)

//nolint:paralleltest
func TestInit(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want env.Env
	}{
		{
			name: "prod",
			env:  "prod",
			want: "prod",
		},
		{
			name: "prev",
			env:  "prev",
			want: "prev",
		},
		{
			name: "dev",
			env:  "dev",
			want: "dev",
		},
		{
			name: "local",
			env:  "local",
			want: "local",
		},
		{
			name: "empty",
			env:  "",
			want: "",
		},
		{
			name: "other",
			env:  "",
			want: "",
		},
	}

	// 環境変数を操作するため直列でテスト
	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("ENV", "")
			env.Init()
			t.Setenv("ENV", tt.env)
			env.Init()
			if !reflect.DeepEqual(env.Get(), tt.want) {
				t.Errorf("Env = %v, want %v", env.Get(), tt.want)
			}
		})
	}
}
