package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// LISTENS service
const LISTENS string = "0.0.0.0:8080"

// LISTENM manage
const LISTENM string = "0.0.0.0:8081"

// IndexHandler index
type IndexHandler struct {
	name string
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, _ = w.Write([]byte(h.name))
}

// CloseHandler close chan
type CloseHandler struct {
	CloseChan chan error
}

func (h *CloseHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, _ = w.Write([]byte("closing"))

	select {
	default:
		h.CloseChan <- errors.New("api shutdown")
	case <-h.CloseChan:
	}

}

func startServer(stopCh chan<- struct{}) {

	mux1 := http.NewServeMux()
	mux1.Handle("/", &IndexHandler{name: "hello golang, service is running.\n"})
	closeHandler1 := &CloseHandler{}
	mux1.Handle("/close", closeHandler1)

	s1Ch := make(chan error, 1)

	mux2 := http.NewServeMux()
	mux2.Handle("/", &IndexHandler{name: "hello golang, manage is running.\n"})
	closeHandler2 := &CloseHandler{}
	mux2.Handle("/close", closeHandler1)

	s2Ch := make(chan error, 1)

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	ctx := context.Background()
	group, _ := errgroup.WithContext(ctx)

	server1 := &http.Server{Addr: LISTENS, Handler: mux1}
	closeHandler1.CloseChan = s1Ch
	server2 := &http.Server{Addr: LISTENM, Handler: mux2}
	closeHandler2.CloseChan = s2Ch

	group.Go(func() error {

		select {
		case <-ch:
			fmt.Println("receive close signal!")
		case err := <-s1Ch:
			fmt.Printf("receive service close! %+v\n", err)
		case err := <-s2Ch:
			fmt.Printf("receive manage close! %+v\n", err)
		}

		signal.Stop(ch)
		close(s1Ch)
		close(s2Ch)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err1 := server1.Shutdown(ctx)
		fmt.Printf("service close %+v \n", err1)

		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		err2 := server2.Shutdown(ctx)
		fmt.Printf("manage close %+v \n", err2)

		return nil
	})

	group.Go(func() error {
		err := server1.ListenAndServe()
		select {
		default:
			s1Ch <- err
		case <-s1Ch:
		}
		fmt.Printf("service closed %+v \n", err)
		return nil
	})

	group.Go(func() error {
		err := server2.ListenAndServe()
		select {
		default:
			s1Ch <- err
		case <-s1Ch:
		}
		fmt.Printf("manage closed %+v \n", err)
		return nil
	})

	err := group.Wait()
	stopCh <- struct{}{}
	fmt.Printf("group err %+v \n", err)
}

func main() {
	stopCh := make(chan struct{}, 1)
	go startServer(stopCh)
	<-stopCh
	close(stopCh)
}
