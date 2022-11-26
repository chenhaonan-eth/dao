package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	"github.com/chenhaonan-eth/dao/pkg/ui/data/swagger"
	"github.com/chenhaonan-eth/dao/pkg/utils"
	pb "github.com/chenhaonan-eth/dao/proto/server"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ServerPort   string
	HttpPort     string
	SwaggerDir   string
	gRPCEndPoint string
	HTTPEndPoint string
)

func Run() {
	gRPCEndPoint = ":" + ServerPort
	HTTPEndPoint = ":" + HttpPort
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", gRPCEndPoint)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	grpcServer := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(grpcServer, &server{})
	// Serve gRPC server
	log.Printf("Serving gRPC on %s", gRPCEndPoint)
	go func() {
		log.Fatalln(grpcServer.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0%s", gRPCEndPoint),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	gwServer := &http.Server{
		Addr:    HTTPEndPoint,
		Handler: utils.GrpcHandlerFunc(grpcServer, mux),
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0%s", HTTPEndPoint)
	log.Fatalln(gwServer.ListenAndServe())
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(SwaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
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
