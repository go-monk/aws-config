package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	log.SetPrefix(filepath.Base(os.Args[0]) + ": ")
	log.SetFlags(0)

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("loading default config: %s", err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(writer, "SOURCE\tTYPE\tREGION")
	for i, source := range cfg.ConfigSources {
		var region string
		switch source := source.(type) {
		case config.LoadOptions:
			region = source.Region
		case config.EnvConfig:
			region = source.Region
		case config.SharedConfig:
			region = source.Region
		}
		fmt.Fprintf(writer, "%d\t%T\t%s\n", i+1, source, region)
	}
	writer.Flush()
}
