package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

var ipAddr = "192.168.1.50"
var confTemplate = `
events {}

http {
  server {
    listen 80;
    return 302 http://%s/search;
  }
}
`

func getIpAddr(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", ipAddr)
}

func updateIpAddr(w http.ResponseWriter, r *http.Request) {
	ipAddr = mux.Vars(r)["ipAddr"]

	fmt.Printf("Client address = %s\n", ipAddr)
	err := writeToFile("nginx.conf", ipAddr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = restartNginx()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Fprintf(w, "%s\n", ipAddr)
}

func writeToFile(filename string, ipAddr string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	conf := fmt.Sprintf(confTemplate, ipAddr)
	_, err = io.WriteString(file, conf)
	if err != nil {
		return err
	}
	return file.Sync()
}

func restartNginx() error {
	cmd := exec.Command("docker", "restart", "nginx")
	_, err := cmd.Output()
	return err
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getIpAddr).Methods("Get")
	r.HandleFunc("/{ipAddr}", updateIpAddr).Methods("POST")

	fmt.Printf("Listening on :8080\n")
	fmt.Print(http.ListenAndServe(":8080", r))
}
