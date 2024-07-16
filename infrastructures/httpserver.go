package infrastructures

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type HttpServer struct {
	Port     string
	ServeMux *http.ServeMux
}

func (hs *HttpServer) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Handler: hs.ServeMux,
	}

	//Create http listener
	listener, err := net.Listen("tcp", ":"+hs.Port)
	if err != nil {
		log.Printf("failed to create http listener with port %s", hs.Port)
		return err
	}

	//Print request url
	log.Printf("URL: %v", fmt.Sprintf("http://%s", listener.Addr().String()))

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		//ErrServerClosed is returned by Server.Serve methods after a call to Server.Shutdown
		if err := server.Serve(listener); err != http.ErrServerClosed {
			log.Printf("failed to close server: %v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred while executing server shutdown: %v", err)
	}

	return eg.Wait()
}
