package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	helloWorldSvr := getHelloWorldServer()
	helloNameSvr := getHelloNameServer()
	echoSvr := getEchoServer()

	go helloWorldSvr.ListenAndServe()
	go helloNameSvr.ListenAndServe()
	go echoSvr.ListenAndServe()

	fmt.Println("all servers are started")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}

func getHelloWorldServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello, world!`))
	})

	return &http.Server{Addr: ":7000", Handler: mux}
}

func getHelloNameServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		name := params.Get("name")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	})

	return &http.Server{Addr: ":8000", Handler: mux}
}

func getEchoServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.Copy(w, r.Body)
	})

	return &http.Server{Addr: ":9000", Handler: mux}
}
