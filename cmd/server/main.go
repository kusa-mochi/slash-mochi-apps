package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slash_mochi/cmd/server/config"
	flexible_reversi_servce "slash_mochi/cmd/server/connect-apis/flexible-reversi"
	omikuji_service "slash_mochi/cmd/server/connect-apis/omikuji"
	test_service "slash_mochi/cmd/server/connect-apis/test"
	flexible_reversi_store "slash_mochi/cmd/server/store/flexible-reversi"
	"slash_mochi/gen/go/slash_mochi/v1/flexible_reversi/flexible_reversiv1connect"
	"slash_mochi/gen/go/slash_mochi/v1/omikuji/omikujiv1connect"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
)

// func webServerRoutine(ip string, port int, dirPath string) {

// 	// // debug -->
// 	// log.Println(dirPath)
// 	// log.Printf("%#v",
// 	// 	http.FileServer(
// 	// 		http.Dir(dirPath),
// 	// 	),
// 	// )
// 	// log.Printf("%#v",
// 	// 	http.StripPrefix(
// 	// 		"/",
// 	// 		http.FileServer(
// 	// 			http.Dir(dirPath),
// 	// 		),
// 	// 	),
// 	// )
// 	// // <-- debug

// 	http.Handle(
// 		"/",
// 		http.StripPrefix(
// 			"/",
// 			http.FileServer(
// 				http.Dir(dirPath),
// 			),
// 		),
// 	)
// 	log.Println("listening requests to the web server...")
// 	log.Fatal(
// 		http.ListenAndServe(
// 			fmt.Sprintf("%s:%v", ip, port),
// 			nil,
// 		),
// 	)
// }

func connectServerRoutine(
	ip string,
	port int,
	flexibleReversiStoreInterface *flexible_reversi_store.FlexibleReversiStoreInterfaces,
) {
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

	// flexible-reversi servce
	flexibleReversiService := flexible_reversi_servce.NewFlexibleReversiService(flexibleReversiStoreInterface)
	flexibleReversiPath, flexibleReversiHandler := flexible_reversiv1connect.NewFlexibleReversiServiceHandler(flexibleReversiService, interceptor)
	mux.Handle(flexibleReversiPath, flexibleReversiHandler)

	// TODO: make CORS rules.
	c := cors.AllowAll()
	corsHandler := c.Handler(mux)

	log.Println("listening requests to the connect server...")
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", ip, port),
			corsHandler,
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
	const (
		CONFIG_FILE_PATH string = "./config.json"
	)

	rawConfig, err := os.ReadFile(CONFIG_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	var config config.Config
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		log.Fatal(err)
	}

	// initialize logger
	logFile, err := os.OpenFile(
		fmt.Sprintf("%s%s", config.Log.Dir, config.Log.FileName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	log.Println("initializing...")

	// // start the web server
	// go webServerRoutine(
	// 	config.WebServer.Ip,
	// 	config.WebServer.Port,
	// 	config.WebServer.Dir,
	// )

	flexibleReversiStore := flexible_reversi_store.NewFlexibleReversiStore()

	// start the connect server using main Go routine
	go connectServerRoutine(
		config.ConnectServer.Ip,
		config.ConnectServer.Port,
		flexibleReversiStore.Interfaces,
	)

	errorToBreak := true
	for errorToBreak {
		select {
		// Flexible Reversi
		case addGlobalChatStreamRequest := <-flexibleReversiStore.Interfaces.AddGlobalChatStreamRequest:
			flexibleReversiStore.Controllers.AddGlobalChatStream(addGlobalChatStreamRequest)
		case broadcastGlobalChat := <-flexibleReversiStore.Interfaces.BroadcastGlobalChatRequest:
			flexibleReversiStore.Controllers.BroadcastGlobalChat(broadcastGlobalChat)
		case flexibleReversiValidateUserId := <-flexibleReversiStore.Interfaces.ValidateUserIdRequest:
			flexibleReversiStore.Controllers.ValidateUserId(flexibleReversiValidateUserId)
		case flexibleReversiGlobalChat := <-flexibleReversiStore.Interfaces.GlobalChatRequest:
			flexibleReversiStore.Controllers.GlobalChat(flexibleReversiGlobalChat)
		}
	}

	log.Println("fin server program")
}
