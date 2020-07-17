package main

import (
  "fmt"
  "net/http"
  "strings"
)

var props = map[string]string{
  //"90min": "www.90min.com/",
  "90min_tr": "www.90min.com/tr",
  "90min": "www.90min.com",
  "90min_es": "www.90min.com/es",
  "90min_es1": "www.90min.com/es/1",
  "90min_es1_asd": "www.90min.com/es/1/asdasd",
  "90min_es2": "www.90min.com/es/2",
  "90min_es3": "www.90min.com/es/3",
  "90min_es4": "www.90min.com/es/4",
  "90min_es5": "www.90min.com/es/5",
}

func main() {
  fmt.Println("hello world")
  
  mux := http.NewServeMux()
  for prop, endpoint := range props {
    router := createHandler(prop)
    prefix := Prefix(endpoint)
    mux.Handle(prefix + "/", router)
    mux.Handle(prefix, router)
  }
  
  // -> route = "www.90min.com/sdes/1/asdasd/asdasd
  
  
  // https://www.90min.com/es/sdes/1/asdasd/asdasd -> pilot
  // pilot -> routing-srv: get route for "www.90min.com/sdes/1/asdasd/asdasd"
  // routing-srv: get property: /es/sdes/1/asdasd/asdasd (by the mux)
  //              get route for prop and path /sdes/1/asdasd/asdasd
  r, err := http.NewRequest("GET", "/es/6/sdes/1/asdasd/asdasd", nil)
  if err != nil {
    panic(err)
  }
  
  w := myWriter{}
  mux.ServeHTTP(&w, r)
  
  fmt.Println("----------------->>>", w.prop)
}


func Prefix(endpoint string) string {
  endpointParts := strings.SplitN(endpoint, "/", 2)
  
  prefix := "/"
  if len(endpointParts) > 1 {
    prefix = "/" + endpointParts[1]
  }
  
  return prefix
}



func createHandler(prop string) http.Handler {
  fmt.Println("----------->> ", prop)
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("-=========--ds LALALALA", prop)
    fmt.Fprint(w, prop)
  })
}

type myWriter struct {
  prop string
}

func (w *myWriter) Write(b []byte) (int, error) {
  w.prop = string(b)
  return 0, nil
}

func (w *myWriter) WriteHeader(status int) {
}

func (w *myWriter) Header() http.Header {
  return http.Header{}
}