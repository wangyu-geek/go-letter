package oss

import (
	"bytes"
	sdkOSS "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"leeter/pkg/errors"
)

type Client struct {
	sdkClient *sdkOSS.Client
}

func NewClient() (*Client, error) {
	client, err := sdkOSS.New(ossConfig.EndPoint, ossConfig.AccessKey, ossConfig.AccessKeySecret)
	if err != nil {
		err = errors.Wrapf(err, "oss create client fail; param:[%+v]", ossConfig)
		return nil, err
	}

	return &Client{sdkClient: client}, nil
}

// PutBytes 上传字节流，TODO 后续可以把client做成单例
func PutBytes(fileName string, allBytes []byte) error {
	client, err := NewClient()
	if err != nil {
		err = errors.Wrapf(err, "oss create client fail")
		return err
	}

	bucket, err := client.sdkClient.Bucket(ossConfig.Bucket)
	if err != nil {
		err = errors.Wrapf(err, "oss create bucket fail; param:[%v]", ossConfig.Bucket)
		return err
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := sdkOSS.ObjectStorageClass(sdkOSS.StorageStandard)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := sdkOSS.ObjectACL(sdkOSS.ACLPublicRead)

	err = bucket.PutObject(fileName, bytes.NewReader(allBytes), storageType, objectAcl)
	if err != nil {
		err = errors.Wrapf(err, "oss put object fail param:[fileName:%v]", fileName)
		return err
	}

	return nil
}
