package q3

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestTeacherDAO_CreateTeacher(t *testing.T) {
	type fields struct {
		mq MessageQueue
		db *gorm.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Teacher
		wantErr bool
	}{
		{
			name:    "",
			fields:  fields{},
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := TeacherDAO{
				mq: tt.fields.mq,
				db: tt.fields.db,
			}
			got, err := dao.CreateTeacher(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTeacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTeacher() got = %v, want %v", got, tt.want)
			}
		})
	}
}
