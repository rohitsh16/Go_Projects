// Used to route traffic to different backends and 
//  return the results to the original client.

package main

import(
  _"fmt"
)

// Load Balancer using simple Round Robin method

type Backend struct {
  URL           *url.URL
  Alive         bool
  mux           sync.RWMutex
  ReverseProxy  *httputil.ReverseProxy
}

type ServerPool struct {
  backends []*Backend
  current uint64
}



func main() {

}
