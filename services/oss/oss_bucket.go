package oss

import (
	"crypto/rand"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
)

var Pool *BucketPool

// BucketPool 链接的bucket池
type BucketPool struct {
	Buckets []*oss.Bucket
	sync.Locker
}

func NewBucketPool() *BucketPool {
	return &BucketPool{
		Buckets: make([]*oss.Bucket, 0, 100),
	}
}

func (pool *BucketPool) GetBucket() *oss.Bucket {
	pool.Lock()
	defer pool.Unlock()
	if len(pool.Buckets) <= 5 {
		// 此时pool中没有可用的bucket
		newBucket, err := connAliOssBucket(os.Getenv("BUCKET_NAME"))
		if err != nil {
			panic(err)
		}
		// put
		pool.Buckets = append(pool.Buckets, newBucket)
		return newBucket
	} else {
		// 从pool中随机选出一个bucket
		n, _ := rand.Int(rand.Reader, big.NewInt(1000))
		index := n.Int64() / 5
		return pool.Buckets[index]
	}
}

// connAliOssBucket 链接阿里云上oss指定的bucket,没有则创建一个
func connAliOssBucket(bucketName string) (*oss.Bucket, error) {
	client, err := oss.New(os.Getenv("ALI_ENDPOINT"), os.Getenv("ACCESS_KEY"), os.Getenv("ACCESS_SECRET"))
	if err != nil {
		return nil, errors.New("链接阿里云oss失败：" + err.Error())
	}
	// 检测是否存在bucket
	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return nil, errors.New("检测bucket是否存在失败：" + err.Error())
	}
	if !exist {
		// 创建bucket,oss访问权限默认为私有
		if err := client.CreateBucket(bucketName); err != nil {
			return nil, errors.New("创建bucket失败：" + err.Error())
		}
	}
	// 获取存储空间
	return client.Bucket(bucketName)
}

// 生成put签名url进行临时授权,链接有效时间设置为600s
func PutSignedUrl(bucket *oss.Bucket, objectName string) (string, error) {
	return signedUrl(bucket, objectName, http.MethodPut, 600)
}

// 生成下载签名url进行临时授权,链接有效时间设置为600s
func GetSignedUrl(bucket *oss.Bucket, objectName string) (string, error) {
	return signedUrl(bucket, objectName, http.MethodGet, 600)
}

// signedUrl
func signedUrl(bucket *oss.Bucket, objectName, httpMethod string, expiredTime int64) (string, error) {
	var method oss.HTTPMethod
	// 检测httpMethod
	switch strings.ToUpper(httpMethod) {
	case "GET":
		method = oss.HTTPGet
	case "PUT":
		method = oss.HTTPPut
	default:
		return "", errors.New("目前只支持GET PUT 方法")
	}

	signedUrl, err := bucket.SignURL(objectName, method, expiredTime) // 设置此链接的有效时间为60s
	return signedUrl, err
}
