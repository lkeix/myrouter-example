package benchmark

import (
	"net/http"
	"testing"

	"github.com/lkeix/myrouter"
)

// write my router benchmark
func BenchmarkMyRoute(b *testing.B) {
	myrouter.NewRouter()

	r := myrouter.NewRouter()

	testcases := []struct {
		endpoint string
		handler  http.Handler
		request  *http.Request
	}{
		{
			endpoint: "/",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/health",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/health", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/hoge",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/hoge", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/fuga",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/fuga", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/piyo",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/piyo", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/piyo/hoge",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/piyo/hoge", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/piyo/fuga",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/piyo/fuga", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/piyo/fuga/hoge",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/piyo/fuga/hoge", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/piyo/fugahogefuga/piyo/fugahogefuga/piyo/fugahogefuga/piyo/fugahogefuga",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/piyo/fugahogefuga/piyo/fugahogefuga/piyo/fugahogefuga/piyo/fugahogefuga", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
		{
			endpoint: "/piyofugahogefugapiyo/piyofugahogefugapiyo/piyofugahogefugapiyo/piyofugahogefugapiyo",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			request: func() *http.Request {
				r, err := http.NewRequest("GET", "/piyofugahogefugapiyo/piyofugahogefugapiyo/piyofugahogefugapiyo/piyofugahogefugapiyo", nil)
				if err != nil {
					b.Fatal(err)
				}
				return r
			}(),
		},
	}

	for _, tc := range testcases {
		r.GET(tc.endpoint, tc.handler)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			r.ServeHTTP(nil, tc.request)
		}
	}
}
