package main

import "log"

type ConnLimiter struct {
	cocurrentConn int
	bucket        chan int
}

// no size channel is syncs
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		cocurrentConn: cc,
		bucket:        make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.cocurrentConn {
		log.Println("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New Connection coming:%d", c)
}
