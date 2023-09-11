package utils

import (
	"sync"
)

var wg sync.WaitGroup

type Golimit struct {
	limit chan int
}

func NewLimit(max int) *Golimit {
	return &Golimit{limit: make(chan int, max)}
}

func (g *Golimit) Add() {
	g.limit <- 1
	wg.Add(1)
}

func (g *Golimit) Done() {
	<-g.limit
	wg.Done()
}

func (g *Golimit) Wait() {
	wg.Wait()
}
