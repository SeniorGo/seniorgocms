package api

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/SeniorGo/seniorgocms/persistence"
)

// type SitemapURL struct {
// 	Loc        string    // URL de la página
// 	LastMod    time.Time // Fecha de última modificación
// 	ChangeFreq string    // Frecuencia de cambio (daily, weekly, etc.)
// 	Priority   float64   // Prioridad (de 0.0 a 1.0)
// }

type sitemap struct {
	postRepo persistence.Persistencer[Post]
}

func newSitemap(postRepo persistence.Persistencer[Post]) *sitemap {
	return &sitemap{postRepo: postRepo}
}

type urlset struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []urlElement `xml:"url"`
}

type urlElement struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float64 `xml:"priority,omitempty"`
}

func (h *sitemap) Handle(w http.ResponseWriter, ctx context.Context) error {
	baseUrl := "https://seniorgocms.holacloud.app" // TODO: hardcoded!

	u := urlset{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	u.URLs = append(u.URLs, urlElement{
		Loc:        baseUrl + "/",
		ChangeFreq: "weekly",
	})

	posts, err := h.postRepo.List(ctx)
	if err != nil {
		return err
	}

	for _, post := range posts {
		u.URLs = append(u.URLs, urlElement{
			Loc:     baseUrl + "/posts/" + post.Id,
			LastMod: post.Item.ModificationTime.Format("2006-01-02"),
			// ChangeFreq: "weekly",
			// Priority:   0.8,
		})
	}

	w.Header().Set("Content-Type", "application/xml")

	w.Write([]byte(xml.Header))

	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	e.Encode(u)
	return nil
}
