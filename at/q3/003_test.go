package q3

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	client, mockDB, _ = sqlmock.New()

	db, _ = gorm.Open("mysql", client)

	mq = new(MockMessageQueue)
)

func TestTeacherDAO_CreateTeacher(t *testing.T) {
	type fields struct {
		mq MessageQueue
		db *gorm.DB
	}
	type args struct {
		name string
		mock func()
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
			fields:  fields{mq: mq, db: db},
			args:    args{name: "test", mock: func() {
				mockDB.ExpectBegin()
				mockDB.ExpectExec("INSERT INTO `teachers` (`name`,`deleted_at`) VALUES (?,?)").
					WithArgs("test", "").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockDB.ExpectCommit()
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

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
