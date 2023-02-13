package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"regexp"
	"uploadService/services/api-gateway/config"
	upload_service "uploadService/services/api-gateway/proto/upload-service"
)

func main() {
	conf := config.GatewayConfig()

	uploadServiceHost := conf.HostUpload
	uploadServicePort := conf.PortUpload

	gatewayHost := conf.HostGateway
	gatewayPort := conf.PortGateway

	uploadServiceAddress := uploadServiceHost + ":" + uploadServicePort
	proxyAddr := gatewayHost + ":" + gatewayPort

	GatewayStart(proxyAddr, uploadServiceAddress)
}

func GatewayStart(proxyAddr, uploadServiceAddress string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := upload_service.RegisterUploadServiceHandlerFromEndpoint(ctx, mux, uploadServiceAddress, opts)
	if err != nil {
		log.Fatalln("Failed to connect to Upload service", err)
	}

	gwServer := &http.Server{
		Addr:    proxyAddr,
		Handler: cors(mux),
	}

	fmt.Println("starting gateway server at " + proxyAddr)
	log.Fatalln(gwServer.ListenAndServe())
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func allowedOrigin(origin string) bool {
	conf := config.GatewayConfig()
	if conf.Cors == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(conf.Cors, origin); matched {
		return true
	}
	return false
}
