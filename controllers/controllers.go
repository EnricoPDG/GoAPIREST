package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/EnricoPDG/GoAPIREST/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
