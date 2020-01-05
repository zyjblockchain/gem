package oss

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
	"time"
)

func init() {
	if err := godotenv.Load(".env_test"); err != nil {
		panic(err)
	}
}

func TestPutSignedUrl(t *testing.T) {
	if err := godotenv.Load(".env_test"); err != nil {
		panic(err)
	}
	bucket1, err := connAliOssBucket(os.Getenv("BUCKET_NAME"))
	bucket2, err := connAliOssBucket(os.Getenv("BUCKET_NAME"))
	t.Error(err)
	for i := 0; i < 5; i++ {
		go func() {
			signedUrl, err := PutSignedUrl(bucket1, "image/小刚几"+strconv.Itoa(1)+".jpg")
			t.Error(err)
			t.Log("go func: ", signedUrl)
		}()
	}

	signedUrl, err := PutSignedUrl(bucket2, "image/小刚几1.jpg")
	t.Error(err)
	t.Log("222: ", signedUrl)
	time.Sleep(2 * time.Second)
}

func TestConnAliOssBucket(t *testing.T) {
	if err := godotenv.Load(".env_test"); err != nil {
		panic(err)
	}
	// bucket,err := connAliOssBucket(os.Getenv("BUCKET_NAME"))
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// err = bucket.PutObjectFromFile("testFile.txt","C:/Users/18382/Documents/test.txt")
	// t.Error(err)
	signedUrl, err := PutSignedUrl(Pool.GetBucket(), "image/小刚几.jpg")
	if err != nil {
		t.Error(err)
	}
	t.Log(signedUrl)
	err = Pool.GetBucket().PutObjectFromFileWithURL(signedUrl, "C:/Users/18382/Pictures/微信图片_20191016231912.jpg")
	t.Log(err)
	//
	// signedUrl, err = bucket.SignURL("image/小刚几.jpg",oss.HTTPGet,60,oss.Process("image/format,png")) // 设置此链接的有效时间为60s
	// t.Log(err)
	// t.Log(signedUrl)
}

func TestGetSignedUrl(t *testing.T) {
	// url,err := GetSignedUrl(os.Getenv("BUCKET_NAME"),"image/小刚几.jpg")
	// t.Log(err)
	// t.Log(url)
	signedUrl, err := PutSignedUrl(Pool.GetBucket(), "image/小刚几1.jpg")
	if err != nil {
		t.Error(err)
	}
	t.Log(signedUrl)
}
