package main

import (
    "log"
    "net/http"
    "github.com/linoecarrillo/tobq/tobq"
)

func main() {
    http.HandleFunc("/api/v1/report", report.ReportHandler)

    log.Println("Server listening on port 3000")
    log.Panic(
        http.ListenAndServe(":3000", nil),
    )
}
