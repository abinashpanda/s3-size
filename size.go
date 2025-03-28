package main

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type s3Config struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

const (
	root = "ROOT"
)

func showSize(config s3Config) error {
	ctx := context.Background()

	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: config.UseSSL,
	})

	if err != nil {
		return err
	}

	sizeMap := make(map[string]int64)
	sizeMap[root] = 0

	childrenMap := make(map[string][]string)

	objectCh := minioClient.ListObjects(ctx, config.Bucket, minio.ListObjectsOptions{Recursive: true})

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return object.Err
		}

		path := strings.Split(object.Key, "/")
		for i := 0; i < len(path); i++ {
			sizeMap[root] += object.Size
			name := strings.Join(path[:i+1], "/")
			if _, ok := sizeMap[name]; !ok {
				sizeMap[name] = object.Size
			} else {
				sizeMap[name] += object.Size
			}

			parent := strings.Join(path[:i], "/")
			if parent == "" {
				parent = root
			}
			if _, ok := childrenMap[parent]; !ok {
				childrenMap[parent] = []string{name}
			} else {
				if !contains(childrenMap[parent], name) {
					childrenMap[parent] = append(childrenMap[parent], name)
				}
			}
		}
	}

	format(root, childrenMap, sizeMap, 0)

	return nil
}

func format(rootNode string, childrenMap map[string][]string, sizeMap map[string]int64, depth int) {
	fmt.Printf("%-20s%s%s\n", formatSize(sizeMap[rootNode]), strings.Repeat(" ", depth*2), rootNode)
	for _, child := range childrenMap[rootNode] {
		format(child, childrenMap, sizeMap, depth+1)
	}
}

func formatSize(size int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	for i, s := range sizes {
		sizeUnit := math.Pow(1024, float64(i))
		nextSizeUnit := math.Pow(1024, float64(i+1))
		if size < int64(nextSizeUnit) {
			return fmt.Sprintf("%.1f %s", float64(size)/sizeUnit, s)
		}
	}
	return ""
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
