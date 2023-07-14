package http

import (
	"fmt"
	"net/http"
)

func startHttp() {
	http.HandleFunc("/receive-event", handleEvent)
	http.ListenAndServe("10.56.180.181:8090", nil)
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	bytes := make([]byte, 1024)
	r.Body.Read(bytes)
	fmt.Printf("接受数据：%v\n", string(bytes))
	w.Write([]byte("success"))
}
