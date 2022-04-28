package oss

import (
	"log"
	"youcaibi/conf"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var EP string
var AK string
var SK string

func init() {
	AK = ""                //access key
	EP = conf.GetOssAddr() // endpoint
	SK = ""                //
}

func UploadToOss(filename, path, buckname string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service error : %s", err)
		return false
	}

	bucket, err := client.Bucket(buckname)
	if err != nil {
		log.Printf("Init oss bucket error : %s", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error : %s", err)
		return false
	}

	return true
}

func DeleteObject(filename, bucketname string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss service erroe:%s", err)
		return false
	}
	bucket, err := client.Bucket(bucketname)
	if err != nil {
		log.Printf("Bucket error:%s", err)
		return false
	}
	err = bucket.DeleteObject(filename)
	if err != nil {
		log.Printf("Delete from oss error:%s", err)
		return false
	}
	return true
}
