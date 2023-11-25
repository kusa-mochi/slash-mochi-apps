package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	omikuji_service "slash_mochi/cmd/server/omikuji"
	test_service "slash_mochi/cmd/server/test"
	"slash_mochi/gen/go/slash_mochi/v1/omikuji/omikujiv1connect"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
)

func webServerRoutine(ipAddr string, port int, dirPath string) {

	// debug -->
	log.Println(dirPath)
	log.Println(http.Dir(dirPath))
	log.Printf("%#v",
		http.FileServer(
			http.Dir(dirPath),
		),
	)
	log.Printf("%#v",
		http.StripPrefix(
			"/",
			http.FileServer(
				http.Dir(dirPath),
			),
		),
	)
	// <-- debug

	http.Handle(
		"/",
		http.StripPrefix(
			"/",
			http.FileServer(
				http.Dir(dirPath),
			),
		),
	)
	log.Println("listening requests to the web server...")
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", ipAddr, port),
			nil,
		),
	)
}

func newServeMuxWithReflection() *http.ServeMux {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"slash_mochi.v1.TestService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return mux
}

func newInterCeptors() connect.Option {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(
			func(
				ctx context.Context,
				req connect.AnyRequest,
			) (connect.AnyResponse, error) {
				procedureName := req.Spec().Procedure
				log.Printf("start %s", procedureName)
				procedureResult, err := next(ctx, req) // Unary RPC call
				log.Printf("end %s", procedureName)

				return procedureResult, err
			},
		)
	}
	return connect.WithInterceptors(
		connect.UnaryInterceptorFunc(interceptor),
	)
}

func main() {
	// get args
	var (
		ipAddr            = flag.String("IpAddr", "127.0.0.1", "Server IP Addr (default=\"127.0.0.1\")")
		connectServerPort = flag.Int("ConnectServerPort", 3081, "Connect server port number (default=3081)")
		webServerPort     = flag.Int("WebServerPort", 80, "Web Server port number (default=80)")
		logPath           = flag.String("LogPath", "./log.txt", "log file path (default=\"./log.txt\")")
	)
	flag.Parse()

	// initialize logger
	logFile, err := os.OpenFile(
		*logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	mux := newServeMuxWithReflection()
	interceptor := newInterCeptors()

	// test service
	testService := test_service.NewTestService()
	testPath, testHandler := testv1connect.NewTestServiceHandler(testService, interceptor)
	mux.Handle(testPath, testHandler)

	// omikuji service
	omikujiService := omikuji_service.NewOmikujiService()
	omikujiPath, omikujiHandler := omikujiv1connect.NewOmikujiServiceHandler(omikujiService, interceptor)
	mux.Handle(omikujiPath, omikujiHandler)

	// TODO: make CORS rules.
	c := cors.AllowAll()
	corsHandler := c.Handler(mux)

	// start the web server
	go webServerRoutine(
		*ipAddr,
		*webServerPort,
		"/project/dist/linux/amd64/clients/test",
	)

	log.Println("listening requests to the connect server...")

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", *ipAddr, *connectServerPort),
			corsHandler,
		),
	)
}
