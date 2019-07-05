package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const DefaultPort int = 8080
const DefaultMapping string = "redirects.conf"

func handler(w http.ResponseWriter, r *http.Request) {
	target := getRedirectTarget(r.URL.Path)
	if target != "" {
		http.Redirect(w, r, target, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

func getRedirectTarget(abbr string) string {
	if strings.HasPrefix(abbr, "/") {
		abbr = abbr[1:]
	}

	fileHandle, err := os.Open(getMappingFile())
	defer fileHandle.Close()

	if err != nil {
		log.Print("Unable to find mappings at: ", getMappingFile())
		return ""
	}

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) == 2 {
			if parts[0] == abbr {
				return parts[1]
			}
		}
	}
	return ""
}

func getMappingFile() string {
	return getEnv("SL_MAPPING", DefaultMapping)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func main() {
	port := fmt.Sprintf(":%d", getEnvInt("SL_PORT", DefaultPort))
	log.Print("Starting shortlink server on ", port)
	log.Print("Using mappings: ", getMappingFile())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
