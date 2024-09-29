package main

import (
	"context"
	"et_test_cl/et_test"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	async := flag.Bool("async", false, "Request multiple thumbnails asynchronously")
	flag.Parse()
	args := flag.Args()

	if !*async && len(args) == 0 {
		fmt.Println("You must provide at least one thumbnail ID or use the --async flag with multiple IDs.")
		return
	}

	conn, err := grpc.Dial("localhost:7002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %s", err.Error())
	}
	defer conn.Close()

	client := et_test.NewThumbnailServiceClient(conn)

	if *async {
		requestAllThumbnails(client, args)
	} else {
		requestThumbnail(client, args[0])
	}
}

func requestThumbnail(client et_test.ThumbnailServiceClient, id string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	res, err := client.GetThumbnail(ctx, &et_test.GetThumbnailRequest{
		Id: id,
	})
	if err != nil {
		log.Printf("Error when calling GetThumbnail for ID %s: %s", id, err.Error())
		return
	}

	fmt.Printf("Received thumbnail for ID %s: %x", id, res.Image)
}

func requestAllThumbnails(client et_test.ThumbnailServiceClient, ids []string) {
	var wg sync.WaitGroup
	wg.Add(len(ids))

	for _, id := range ids {
		go func(id string) {
			defer wg.Done()
			requestThumbnail(client, id)
		}(id)
	}

	wg.Wait()
	fmt.Println("All asynchronous thumbnail requests completed.")
}
