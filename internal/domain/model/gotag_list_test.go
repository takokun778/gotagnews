package model_test

import (
	"reflect"
	"testing"

	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
)

func TestGotagListTake(t *testing.T) {
	t.Parallel()

	type args struct {
		list model.GotagList
	}

	tests := []struct {
		name string
		gl   model.GotagList
		args args
		want model.GotagList
	}{
		{
			name: "take one",
			gl: model.GotagList{
				model.Gotag{ID: gotag.ID("a")},
				model.Gotag{ID: gotag.ID("b")},
				model.Gotag{ID: gotag.ID("c")},
			},
			args: args{
				list: model.GotagList{
					model.Gotag{ID: gotag.ID("a")},
					model.Gotag{ID: gotag.ID("b")},
				},
			},
			want: model.GotagList{
				model.Gotag{ID: gotag.ID("c")},
			},
		},
		{
			name: "take one",
			gl: model.GotagList{
				model.Gotag{ID: gotag.ID("a")},
				model.Gotag{ID: gotag.ID("b")},
				model.Gotag{ID: gotag.ID("c")},
			},
			args: args{
				list: model.GotagList{
					model.Gotag{ID: gotag.ID("a")},
				},
			},
			want: model.GotagList{
				model.Gotag{ID: gotag.ID("b")},
				model.Gotag{ID: gotag.ID("c")},
			},
		},
		{
			name: "take zero",
			gl: model.GotagList{
				model.Gotag{ID: gotag.ID("a")},
				model.Gotag{ID: gotag.ID("b")},
			},
			args: args{
				list: model.GotagList{
					model.Gotag{ID: gotag.ID("a")},
					model.Gotag{ID: gotag.ID("b")},
				},
			},
			want: model.GotagList{},
		},
		{
			name: "empty",
			gl:   model.GotagList{},
			args: args{
				list: model.GotagList{},
			},
			want: model.GotagList{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.gl.Take(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GotagList.Take() = %v, want %v", got, tt.want)
			}
		})
	}
}
