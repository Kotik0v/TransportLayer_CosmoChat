package main

import (
        "encoding/json"
        "fmt"
        "io"
        "log"
        "net/http"
)

type httpSegmentMessage struct {
        Payload       string `json:"payload"`
        Username      string `json:"username"`
        SendTime      string `json:"sendTime"`
        Number        uint32 `json:"segmentNumber"`
        TotalSegments uint32 `json:"totalSegments"`
}

func main() {
        http.HandleFunc("/segment", func(w http.ResponseWriter, r *http.Request) {
                if r.Method != http.MethodPost {
                        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
                        return
                }

                body, err := io.ReadAll(r.Body)
                if err != nil {
                        http.Error(w, "Failed to read request body", http.StatusInternalServerError)
                        return
                }
                defer r.Body.Close()

                var segment httpSegmentMessage
                err = json.Unmarshal(body, &segment)
                if err != nil {
                        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
                        return
                }

                log.Printf("Received segment: %+v\n", segment)

                // Имитация успешного ответа
                w.WriteHeader(http.StatusOK)
                _, _ = w.Write([]byte("OK"))
        })

        port := 8088
        log.Printf("Mock Data Link service is running on port %d", port)
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}