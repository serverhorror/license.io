package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/serverhorror/license.io/data"
)

type Link struct {
	HRef string `json:"href"`
	Rel  string `json:"rel,omitempty"`
}

func NewLink(href string, rel string) *Link {
	return &Link{
		HRef: href,
		Rel:  rel,
	}
}

type License struct {
	Name  string `json:"license"`
	Links []Link `json:"links"`
}

func NewLicense(name string, links []Link) *License {
	return &License{
		Name:  name,
		Links: links,
	}
}

func HandleLicense(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.URL.Path: %#q", r.URL.Path)

	p := r.URL.Path
	switch r.URL.Path {
	case "/api/":
		var lic []*License
		for _, elem := range data.AssetNames() {
			lnks := []Link{
				Link{
					HRef: fmt.Sprintf("http://localhost:8080/api/%s", elem),
					Rel:  "self",
				},
			}
			li := NewLicense(elem, lnks)
			lic = append(lic, li)
		}
		b, err := json.Marshal(lic)
		if err != nil {
			log.Printf("err: %#q", err)
			msg := fmt.Sprintf("\n\nerr: %#q (%T)\n", err, err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
		w.Header().Set("content-type", "application/vnd.licensio.v1+json")
		fmt.Fprintf(w, "%s", b)
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
