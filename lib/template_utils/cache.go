package template_utils

import (
	"html/template"
	"sync"
)

type TemplateCache struct {
	cache map[string]*template.Template
	mu    sync.RWMutex
}

var globalCache *TemplateCache

func init() {
	globalCache = &TemplateCache{
		cache: make(map[string]*template.Template),
	}
}

func GetTemplate(name string, files ...string) (*template.Template, error) {
	globalCache.mu.RLock()
	if tmpl, exists := globalCache.cache[name]; exists {
		globalCache.mu.RUnlock()
		return tmpl, nil
	}
	globalCache.mu.RUnlock()

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	globalCache.mu.Lock()
	globalCache.cache[name] = tmpl
	globalCache.mu.Unlock()

	return tmpl, nil
}

func ClearCache() {
	globalCache.mu.Lock()
	globalCache.cache = make(map[string]*template.Template)
	globalCache.mu.Unlock()
}
