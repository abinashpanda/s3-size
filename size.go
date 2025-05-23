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
	endPoint  string
	bucket    string
	accessKey string
	secretKey string
	useSSL    bool
	maxDepth  int
}

func showSize(config s3Config) error {
	ctx := context.Background()

	minioClient, err := minio.New(config.endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.accessKey, config.secretKey, ""),
		Secure: config.useSSL,
	})

	if err != nil {
		return err
	}

	root := config.bucket
	sizeMap := make(map[string]int64)
	sizeMap[root] = 0

	childrenMap := make(map[string][]string)

	objectCh := minioClient.ListObjects(ctx, root, minio.ListObjectsOptions{Recursive: true})

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return object.Err
		}

		sizeMap[root] += object.Size
		path := strings.Split(object.Key, "/")
		for i := 0; i < len(path); i++ {
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

	printSize(root, childrenMap, sizeMap, 0, config.maxDepth)

	return nil
}

func printSize(rootNode string, childrenMap map[string][]string, sizeMap map[string]int64, depth int, maxDepth int) {
	if depth >= maxDepth {
		return
	}

	fmt.Printf("%-20s%s%s\n", formatSize(sizeMap[rootNode]), strings.Repeat(" ", depth*2), rootNode)
	for _, child := range childrenMap[rootNode] {
		printSize(child, childrenMap, sizeMap, depth+1, maxDepth)
	}
}

func formatSize(size int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	for i, s := range sizes {
		sizeUnit := math.Pow(1024, float64(i))
		nextSizeUnit := math.Pow(1024, float64(i+1))
		if size < int64(nextSizeUnit) {
			return fmt.Sprintf("%.2f %s", float64(size)/sizeUnit, s)
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
