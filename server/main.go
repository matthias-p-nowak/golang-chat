package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	// "time"

	"gopkg.in/ini.v1"
)

var (
	// the root directory of the web content
	webroot string
	// starting time of the server
	// startTime=time.Now().Unix()
)

// the default serving of files for the webchat
func fileServe(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if len(path) > 0 && (path[0] == '/' || path[0] == '.' || strings.Contains(path, "..")) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	if path == "" {
		path = "index.html"
	}
	// the file to load and transfer
	file := webroot + "/" + path
	http.ServeFile(w, r, file)
	fmt.Printf("served file %s\n", file)
}



// the usual main function
func main() {
	fmt.Println("Starting up")
	cfg, err := ini.Load("server/config.ini")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}
	str := cfg.Section(ini.DefaultSection).Key("hello").String()
	fmt.Println(str)
	server := cfg.Section("server")
	local := server.Key("local").String()
	webroot = server.Key("webroot").String()
	secure := server.Key("secure").MustBool()
	http.HandleFunc("GET /", fileServe)
	fmt.Printf("listening on port %s\n", local)
	if secure {
		fmt.Println("secure")
		certFile := server.Key("certFile").String()
		keyFile := server.Key("keyFile").String()
		err = http.ListenAndServeTLS(local, certFile, keyFile, nil)
		if err != nil {
			fmt.Printf("https server failed %s\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("not secure")
		err = http.ListenAndServe(local, nil)
		if err != nil {
			fmt.Printf("http server failed %s\n", err)
			os.Exit(1)
		}
	}

}
