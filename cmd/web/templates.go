package main

import (
	"html/template"
	"main/internal/models"
	"path/filepath"
	"time" // New import
)

// Add a CurrentYear field to the templateData struct.
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Find all page templates
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Create a new template and register the custom FuncMap
		ts := template.New(name).Funcs(functions)

		// Parse the base template
		ts, err = ts.ParseFiles("./ui/html/pages/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse partial templates
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse the specific page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the compiled template to the cache
		cache[name] = ts
	}

	return cache, nil
}
