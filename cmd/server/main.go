package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	test_service "slash_mochi/cmd/server/test"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"github.com/rs/cors"
)

func main() {
	// get args
	var (
		ipAddr            = flag.String("IpAddr", "127.0.0.1", "Server IP Addr (default=\"127.0.0.1\")")
		connectServerPort = flag.Int("ConnectServerPort", 3081, "Connect server port number (default=3081)")
		logPath           = flag.String("LogPath", "/var/slash-mochi/log.txt", "log file path (default=\"/var/slash-mochi/log.txt\")")
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

	mux := http.NewServeMux()
	testService := test_service.NewTestService()
	testPath, testHandler := testv1connect.NewTestServiceHandler(testService)
	mux.Handle(testPath, testHandler)

	// TODO: make CORS rules.
	c := cors.AllowAll()
	corsHandler := c.Handler(mux)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%v", *ipAddr, *connectServerPort),
			corsHandler,
		),
	)
}
