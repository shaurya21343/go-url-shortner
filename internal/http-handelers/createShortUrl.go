package httphandelers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shaurya21343/go-url-shortner/internal/database"
	"github.com/shaurya21343/go-url-shortner/internal/utils"
)

type BodyType struct {
	Url string
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	var body BodyType

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "invlid json body", http.StatusBadRequest)
		return
	}
	isValid := utils.IsValidURL(body.Url)
	if isValid == false {
		http.Error(w, "url is not valid", http.StatusBadRequest)
		return
	}

	id, err := database.CreateShortUrl(body.Url)
	if err != nil {
		http.Error(w, "db error", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "url created id:", id)

	defer r.Body.Close()

}
