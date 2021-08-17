package errors

import (
	"log"
	"testing"
)

func TestErrorType_New(t *testing.T) {
	err := New("hello")
	err = Wrap(err, "world")
	err = Wrap(err, "test")
	log.Println(err)
	ret := Cause(err)
	log.Println(ret)

	//type args struct {
	//	msg string
	//}
	//tests := []struct {
	//	name      string
	//	errorType ErrorType
	//	args      args
	//	wantErr   bool
	//}{
	//	// TODO: Add test cases.
	//	{
	//		name:"hello",
	//		errorType:AliTTSToken,
	//		args:args{msg:"hello"},
	//		wantErr:true,
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if err := tt.errorType.New(tt.args.msg); (err != nil) != tt.wantErr {
	//			t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
	//		}
	//	})
	//}
}
