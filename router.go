package main

import(
    "net/http"
)

type router struct{
  h map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method string, url string, h http.HandlerFunc){
  //todo : map[string]map[string]func 에 핸들러 함수 집어넣기.
  if _, ok := r.h[method]; ok{ //find existing method
    r.h[method][url] = h
  } else{
    newMap := make(map[string]http.HandlerFunc) // make new map
    r.h[method] = newMap
    r.h[method][url] = h
  }
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request){
  if _, ok := r.h[req.Method]; ok { //find existing method
    h := r.h[req.Method][req.URL.Path]
    h(w, req);
  } else{
    w.WriteHeader(http.StatusNotFound)
  }
}
