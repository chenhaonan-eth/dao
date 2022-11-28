package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/chenhaonan-eth/dao/biz"
	"github.com/chenhaonan-eth/dao/pkg/ui/data/swagger"
	"github.com/chenhaonan-eth/dao/pkg/utils"
	pb "github.com/chenhaonan-eth/dao/proto/server"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ServerPort   string
	HttpPort     string
	SwaggerDir   string
	gRPCEndPoint string
	HTTPEndPoint string
	Network      string
)

func Run() {
	gRPCEndPoint = ":" + ServerPort
	HTTPEndPoint = ":" + HttpPort

	errCh := make(chan error, 2)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		if err := rungRPC(ctx); err != nil {
			errCh <- fmt.Errorf("cannot run grpc service: %v", err)
		}
	}()

	go func() {
		if err := rungateway(ctx); err != nil {
			errCh <- fmt.Errorf("cannot run gateway service: %v", err)
		}
	}()

	err := <-errCh
	fmt.Fprintln(os.Stderr, err)
}
func rungateway(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	conn, err := dial(ctx, Network, gRPCEndPoint)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close a client connection to the gRPC server: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/openapiv2/", openAPIServer(SwaggerDir))
	mux.HandleFunc("/healthz", healthzServer(conn))
	serveSwaggerUI(mux)

	gw, err := newGateway(ctx, conn)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    HTTPEndPoint,
		Handler: utils.AllowCORS(mux),
	}
	go func() {
		<-ctx.Done()
		log.Println("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown http server: %v", err)
		}
	}()

	log.Printf("Starting gateway listening at %s", HTTPEndPoint)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}

// Run starts the example gRPC service.
// network and address are passed to net.Listen.
func rungRPC(ctx context.Context) error {
	l, err := net.Listen(Network, gRPCEndPoint)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			log.Printf("Failed to close %s %s: %v", Network, gRPCEndPoint, err)
		}
	}()

	// Create a gRPC server object
	grpcServer := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(grpcServer, biz.NewServer())
	// Serve gRPC server
	log.Printf("Serving gRPC on %s", gRPCEndPoint)
	go func() {
		defer grpcServer.GracefulStop()
		<-ctx.Done()
	}()
	return grpcServer.Serve(l)
}

// openAPIServer returns OpenAPI specification files located under "/openapiv2/"
func openAPIServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			log.Printf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/openapiv2/")
		p = path.Join(dir, p)
		log.Printf("Serving swagger-file: %s", p)
		http.ServeFile(w, r, p)
	}
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func dial(ctx context.Context, network, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr)
	case "unix":
		return dialUnix(ctx, addr)
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// dialUnix creates a client connection via a unix domain socket.
// "addr" must be a valid path to the socket.
func dialUnix(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	d := func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", addr)
	}
	return grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(d))
}

// healthzServer returns a simple health handler which returns ok.
func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {
	mux := runtime.NewServeMux()
	for _, f := range []func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error{
		pb.RegisterGreeterHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}
