package di

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/Coldwws/chat_practice/internal/closer"
	"github.com/Coldwws/chat_practice/internal/config"
	"github.com/Coldwws/chat_practice/internal/interceptor"
	desc "github.com/Coldwws/chat_practice/pkg/chat_v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config          *config.Config
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.InitDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
	}
	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Run() error {
	defer func ()  {
		closer.CloseAll()
		closer.Wait()
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
		if err != nil{
			log.Printf("GRPC server error: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := a.RunHTTPServer()
		if err != nil{
			log.Printf("HTTP server error: %v", err)
		}

	}()
	
	wg.Wait()
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	//для локального запуска
	if err := godotenv.Load("local.env"); err != nil {
		log.Println("Warning: local.env not found, using system env")
	}

	cfg := config.LoadConfig()

	a.config = &cfg

	return nil

}

func (a *App) initServiceProvider(_ context.Context) error {

	a.serviceProvider = NewServiceProvider(a.config)

	return nil

}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.ValidateInterceptor,
			interceptor.AuthInterceptor(a.serviceProvider.AuthClient()),
		),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatAPI())

	return nil

}

func (a *App) runGRPCServer() error {

	log.Println("GRPC server is running on:", a.serviceProvider.config.GRPC.Addr())

	list, err := net.Listen("tcp", a.serviceProvider.config.GRPC.Addr())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil

}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	a.httpServer = &http.Server{
		Addr:    a.config.Http.Address(),
		Handler: mux,
	}
	go func() {

		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		grpcAddr := a.serviceProvider.config.GRPC.Addr()
		for {
			log.Println("Trying to register HTTP Gateway...")

			err := desc.RegisterChatV1HandlerFromEndpoint(ctx, mux, grpcAddr, opts)

			if err != nil {
				log.Println("Failed to connect to gRPC, retrying in 1s:", err)
				time.Sleep(time.Second)
				continue
			}

			log.Println("HTTP Gateway registered successfully")
			break
		}
	}()

	return nil
}


func (a *App) RunHTTPServer() error {
	log.Println("HTTP server is running on:", a.config.Http.Address())
	err := a.httpServer.ListenAndServe()
	if err != nil{
		return err
	}
	return nil
}

