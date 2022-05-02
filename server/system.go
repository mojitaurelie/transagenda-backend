package server

import (
	"github.com/transagenda-back/config"
	"net/http"
)

func AllowRegister(w http.ResponseWriter, r *http.Request) {
	result := map[string]bool{
		"allow_register": config.Features().AllowRegister,
	}
	ok(result, w, r)
}
