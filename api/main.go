package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/serverhorror/license.io/data"
)

func HandleLicense(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.URL.Path: %#q", r.URL.Path)
	fmt.Fprintf(w, "r.URL.Path: %#q\n", r.URL.Path)

	p := r.URL.Path
	switch r.URL.Path {
	case "/api/":
		fmt.Fprintf(w, "%#q", data.AssetNames())
	default:

		license := p[len("/api/"):]
		licenseData, err := data.Asset(license)
		if err != nil {
			log.Printf("err: %#q", err)
			msg := fmt.Sprintf("\n\nerr: %#q (%T)\n", err, err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Requested license: %#q\n", license)
		fmt.Fprintf(w, "Requested license text:\n%s\n", licenseData)
	}
}
