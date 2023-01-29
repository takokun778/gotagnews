package gotag_test

import (
	"testing"

	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		want    gotag.ID
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: "id",
			},
			want:    gotag.ID("id"),
			wantErr: false,
		},
		{
			name:    "empty",
			args:    args{id: ""},
			want:    gotag.ID(""),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := gotag.NewID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewID() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("NewID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		id   gotag.ID
		want string
	}{
		{
			name: "success",
			id:   gotag.ID("id"),
			want: "id",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.id.String(); got != tt.want {
				t.Errorf("ID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
