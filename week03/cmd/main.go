package main

import (
	"context"
	"fmt"
	"homework/week03/httpserver"
	"homework/week03/signal"
	"homework/week03/workgroup"
	"net/http"
	"time"
)

func main()  {
	// Create workgroup
	var wg workgroup.Group
	// Add function to cancel execution using os signal
	wg.Add(signal.Signal())
	// Create http server
	srv := http.Server{Addr: "127.0.0.1:8080"}
	// Add function to start and stop http server
	wg.Add(httpserver.Server(
		func() error {
			fmt.Printf("Server listen at %v\n", srv.Addr)
			err := srv.ListenAndServe()
			fmt.Printf("Server stopped listening with error: %v\n", err)
			if err != http.ErrServerClosed {
				return err
			}
			return nil
		},
		func() error {
			fmt.Println("Server is about to shutdown")
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
			defer cancel()

			err := srv.Shutdown(ctx)
			fmt.Printf("Server shutdown with error: %v\n", err)
			return err
		},
	))
	// Create context to cancel execution after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	// Add function to cancel execution using context
	wg.Add(workgroup.Context(ctx))
	// Execute each function
	err := wg.Run()
	fmt.Printf("Workgroup run stopped with error: %v\n", err)
	// Output in case execution has been canceled with context:
	// Server listen at 127.0.0.1:8080
	// Context canceled
	// Server is about to shutdown
	// Server stopped listening with error: http: Server closed
	// Server shutdown with error: <nil>
	// Workgroup run stopped with error: context canceled

	// Output in case execution has been interrupted with os signal:
	// Server listen at 127.0.0.1:8080
	// Server is about to shutdown
	// Server stopped listening with error: http: Server closed
	// Server shutdown with error: <nil>
	// Workgroup run stopped with error: exist by os signal
}
