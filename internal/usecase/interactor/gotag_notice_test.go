package interactor_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/takokun778/gotagnews/internal/domain/external"
	"github.com/takokun778/gotagnews/internal/domain/model"
	"github.com/takokun778/gotagnews/internal/domain/model/gotag"
	"github.com/takokun778/gotagnews/internal/domain/repository"
	"github.com/takokun778/gotagnews/internal/usecase/interactor"
	"github.com/takokun778/gotagnews/internal/usecase/mock"
	"github.com/takokun778/gotagnews/internal/usecase/port"
)

func TestGotagNoticeExecute(t *testing.T) {
	t.Parallel()

	type fields struct {
		gotagRepository  repository.Gotag
		githubRepository repository.GitHub
		external         external.Gotag
	}

	type args struct {
		input port.GotagNoticeInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    port.GotagNoticeOutput
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				gotagRepository: &mock.Gotag{
					T:    t,
					List: model.GotagList{model.Gotag{ID: gotag.ID("a")}},
				},
				githubRepository: &mock.GitHub{
					T:    t,
					List: model.GotagList{model.Gotag{ID: gotag.ID("a")}, model.Gotag{ID: gotag.ID("b")}},
				},
				external: &mock.Gotag{
					T:    t,
					List: model.GotagList{},
					Want: model.GotagList{model.Gotag{ID: gotag.ID("b")}},
				},
			},
			args: args{
				input: port.GotagNoticeInput{},
			},
			want:    port.GotagNoticeOutput{},
			wantErr: false,
		},
		{
			name: "gotag repository find all error",
			fields: fields{
				gotagRepository: &mock.Gotag{
					T:          t,
					List:       model.GotagList{},
					ErrFindAll: fmt.Errorf("find all error"),
				},
				githubRepository: &mock.GitHub{
					T:    t,
					List: model.GotagList{},
				},
				external: &mock.Gotag{
					T:    t,
					List: model.GotagList{},
				},
			},
			args: args{
				input: port.GotagNoticeInput{},
			},
			want:    port.GotagNoticeOutput{},
			wantErr: true,
		},
		{
			name: "github repository find all error",
			fields: fields{
				gotagRepository: &mock.Gotag{
					T:    t,
					List: model.GotagList{},
				},
				githubRepository: &mock.GitHub{
					T:          t,
					List:       model.GotagList{},
					ErrFindAll: fmt.Errorf("find all error"),
				},
				external: &mock.Gotag{
					T:    t,
					List: model.GotagList{},
				},
			},
			args: args{
				input: port.GotagNoticeInput{},
			},
			want:    port.GotagNoticeOutput{},
			wantErr: true,
		},
		{
			name: "gotag repository save all error",
			fields: fields{
				gotagRepository: &mock.Gotag{
					T:          t,
					List:       model.GotagList{},
					ErrSaveAll: fmt.Errorf("save all error"),
				},
				githubRepository: &mock.GitHub{
					T:    t,
					List: model.GotagList{},
				},
				external: &mock.Gotag{
					T:    t,
					List: model.GotagList{},
				},
			},
			args: args{
				input: port.GotagNoticeInput{},
			},
			want:    port.GotagNoticeOutput{},
			wantErr: true,
		},
		{
			name: "gotag external notice error",
			fields: fields{
				gotagRepository: &mock.Gotag{
					T:    t,
					List: model.GotagList{},
				},
				githubRepository: &mock.GitHub{
					T:    t,
					List: model.GotagList{},
				},
				external: &mock.Gotag{
					T:         t,
					List:      model.GotagList{},
					ErrNotice: fmt.Errorf("notice error"),
				},
			},
			args: args{
				input: port.GotagNoticeInput{},
			},
			want:    port.GotagNoticeOutput{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gni := interactor.NewGotagNotice(
				tt.fields.gotagRepository,
				tt.fields.githubRepository,
				tt.fields.external,
			)
			got, err := gni.Execute(context.Background(), tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GotagNotice.Execute() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GotagNotice.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
