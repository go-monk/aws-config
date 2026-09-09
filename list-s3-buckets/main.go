package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	log.SetPrefix(filepath.Base(os.Args[0]) + ": ")
	log.SetFlags(0)

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	out, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range out.Buckets {
		fmt.Println(*bucket.Name)
	}
}
