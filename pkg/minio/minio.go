package minio

import (
	"context"
	"db-go-gin/internal/global"
	"fmt"
	"github.com/piupuer/go-helper/pkg/oss"
	"time"
)

func InitMinio() (*oss.MinioOss, error) {
	init := false
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		if !init {
			panic(fmt.Sprintf("initialize object storage minio failed: connect timeout(%ds)", 10))
		}
		return
	}()

	ops := []func(*oss.MinioOptions){
		oss.WithMinioEndpoint(global.CONFIG.Minio.Endpoint),
		oss.WithMinioAccessId(global.CONFIG.Minio.AccessId),
		oss.WithMinioSecret(global.CONFIG.Minio.Secret),
		oss.WithMinioHttps(global.CONFIG.Minio.UseHttps),
	}

	minio, err := oss.NewMinio(ops...)
	if err != nil {
		return nil, err
	}

	minio.MakeBucket(ctx, global.CONFIG.Minio.Bucket)

	return minio, err
}
