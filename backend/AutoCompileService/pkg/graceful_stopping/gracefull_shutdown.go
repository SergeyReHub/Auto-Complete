package graceful_stopping

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

func GracefulShutDown(servers ...*grpc.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнал завершения
	<-signalChan
	log.Println("Received shutdown signal, shutting down gracefully...")

	var wg sync.WaitGroup
	for _, srv := range servers {
		if srv != nil {
			wg.Add(1)
			go func(s *grpc.Server) {
				defer wg.Done()
				s.GracefulStop()
				log.Println("gRPC server stopped")
			}(srv)
		}
	}

	wg.Wait()
	log.Println("All servers stopped")
}
