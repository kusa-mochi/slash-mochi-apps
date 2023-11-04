package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	test_service "slash_mochi/cmd/server/test"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
)

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
	testService := test_service.NewTestService()
	testPath, testHandler := testv1connect.NewTestServiceHandler(testService, interceptor)
	mux.Handle(testPath, testHandler)

	// TODO: make CORS rules.
	c := cors.AllowAll()
	corsHandler := c.Handler(mux)

	log.Println("listening...")

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", *ipAddr, *connectServerPort),
			corsHandler,
		),
	)
}
