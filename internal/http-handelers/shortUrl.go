package httphandelers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shaurya21343/go-url-shortner/internal/database"
)

func ShortUrl(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	numId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	url, err := database.GetRealUrl(int64(numId))
	if err != nil {
		fmt.Fprint(w, "ERROR IN ID")
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
