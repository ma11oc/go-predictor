package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	// "os"
	"time"

	"github.com/ma11oc/go-predictor/internal/core"
	pb "github.com/ma11oc/go-predictor/pkg/api/v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type config struct {
	birthday string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.birthday, "birthday", "", "birthday in format YY-mm-dd")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPredictorClient(conn)

	// Contact the server and print out its response.
	// name := defaultName
	/*
	 * if len(os.Args) > 1 {
	 *     name = os.Args[1]
	 * }
	 */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetBaseMatrix(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, v := range r.M.Values {
		fmt.Printf("%v\n", v.GetNumberValue())
	}

	birthday, err := time.Parse("2006-01-02", cfg.birthday)

	if err != nil {
		log.Fatalf("Unable to parse birthday")
	}

	fmt.Println(birthday)

	newDate := time.Date(2000, birthday.Month(), birthday.Day(), 0, 0, 0, 0, time.UTC)

	pc := core.GetCurrentPeriodicityCircles(newDate)

	core.PrintAllPeriodicityCicles()

	fmt.Println()

	for i, v := range pc {
		fmt.Printf(" %v: %v-%v | %v\n", i, v.Start.Format("02/01"), v.End.Format("02/01"), core.Planets[i])
	}

	// fmt.Printf("your card is: %v\n", core.GetCardByBirthDate(&birthday))
	// fmt.Printf("originMatrix: %v", core.originMatrix)
}
