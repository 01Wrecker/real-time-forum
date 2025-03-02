package main

import "net/http"

func main() {
	server := &http.Server{Addr: ":8080"}
	

	go func() {
		log.Info("server started: http://localhost:8080")
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error("error starting server:", err)
			sigChan <- syscall.SIGTERM
		}
	}()
}
