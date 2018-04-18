package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	allowOrigin    = "Access-Control-Allow-Origin"
	allowMethods   = "Access-Control-Allow-Methods"
	allowHeaders   = "Access-Control-Allow-Headers"
	exposeHeaders  = "Access-Control-Expose-Headers"
	requestMethod  = "Access-Control-Request-Method"
	requestHeaders = "Access-Control-Request-Headers"
	maxAge         = "Access-Control-Max-Age"
	headerOrigin   = "Origin"
)

type Cors struct {
	pattern        string
	regex          bool
	allowedHeaders []string
	allowedMethods []string
	exposedHeaders []string
}

func NewCors() *Cors {
	return &Cors{}
}

func (c *Cors) AllowAll() *Cors {
	return c.Domain("*")
}

func (c *Cors) Regex(pattern string) *Cors {
	c.pattern = pattern
	c.regex = true
	return c
}

func (c *Cors) Domain(domain string) *Cors {
	c.pattern = domain
	c.regex = false
	return c
}

func (c *Cors) AllowHeaders(headers ...string) *Cors {
	c.allowedHeaders = append(c.allowedHeaders, headers...)
	return c
}

func (c *Cors) AllowMethods(methods ...string) *Cors {
	c.allowedMethods = append(c.allowedMethods, methods...)
	return c
}

func (c *Cors) ExposeHeaders(headers ...string) *Cors {
	c.exposedHeaders = append(c.exposedHeaders, headers...)
	return c
}

func (c *Cors) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("CORS!")
		if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		origin := r.Header.Get(headerOrigin)

		var (
			matches bool
			err     error
		)

		if c.regex {
			matches, err = regexp.MatchString(c.pattern, origin)
		} else if c.pattern == "*" {
			matches = true
		} else {
			matches = strings.Compare(c.pattern, origin) == 0
		}

		if err != nil {
			log.Fatal(err)
		}

		if matches {
			w.Header().Set(allowOrigin, origin)
			if len(c.allowedHeaders) > 0 {
				w.Header().Set(allowHeaders, join(c.allowedHeaders))
			}
			if len(c.allowedMethods) > 0 {
				w.Header().Set(allowMethods, join(c.allowedMethods))
			}
			if len(c.exposedHeaders) > 0 {
				w.Header().Set(exposeHeaders, join(c.exposedHeaders))
			}
		}

		w.WriteHeader(http.StatusOK)
	})
}

func join(values []string) string {
	var sb strings.Builder
	for i, item := range values {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(item)
	}

	str := sb.String()
	return str
}
