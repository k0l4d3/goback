import (
	"fmt"
	"net/http"
)

// Daily feed
func dailyFeedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the daily feed.")
}