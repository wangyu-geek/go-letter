package oss

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestPutBytes(t *testing.T) {
	fd, err := os.Open("test/3ec12ea1-281f-41a1-b0d8-3b7fac325662.wav")
	if err != nil {
		t.Errorf("open file fail; err:%s", err)
		return
	}
	defer fd.Close()

	byteSlice := make([]byte, 256)
	var allByte []byte
	for {
		readByteNum, err := fd.Read(byteSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return
		}
		//注意append切片的时候后面要加上三个点...
		//注意当最后一次读取当时候，有可能读取当不是256个字节，
		//所以不能把byteSlice整个进行append，否则会出现数据重复，这里是个坑
		allByte = append(allByte, byteSlice[:readByteNum]...)
	}

	if len(allByte) == 0 {
		t.Errorf("read file empty fail;")
		return
	}

	type args struct {
		fileName string
		allBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "测试文件",
			args:    args{fileName: "test/3ec12ea1-281f-41a1-b0d8-3b7fac325662.wav", allBytes: allByte},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PutBytes(tt.args.fileName, tt.args.allBytes); (err != nil) != tt.wantErr {
				t.Errorf("PutBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
