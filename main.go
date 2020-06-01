package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	optPort   string
	optTarget string

	OK = []byte("OK")
)

func init() {
	optPort = os.Getenv("PORT")
	if optPort == "" {
		optPort = "80"
	}
	optTarget = os.Getenv("TARGET")
	if optTarget == "" {
		optTarget = "https://example.com"
	}
}

func exit(err *error) {
	if *err != nil {
		log.Printf("exited with error: %s", (*err).Error())
		os.Exit(1)
	} else {
		log.Println("exited")
	}
}

func main() {
	var err error
	defer exit(&err)

	// healthz
	http.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain")
		rw.Header().Set("Content-Length", strconv.Itoa(len(OK)))
		_, _ = rw.Write(OK)
	})

	// redirect
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(optTarget, "/") {
			url := *req.URL
			url.Host = ""
			url.Scheme = ""
			http.Redirect(rw, req, optTarget+strings.TrimPrefix(url.String(), "/"), http.StatusTemporaryRedirect)
		} else {
			http.Redirect(rw, req, optTarget, http.StatusTemporaryRedirect)
		}
	})

	chErr := make(chan error, 1)
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		chErr <- http.ListenAndServe(":"+optPort, nil)
	}()

	select {
	case err = <-chErr:
	case sig := <-chSig:
		log.Printf("signal: %s", sig.String())
		time.Sleep(time.Second * 3)
	}
}
