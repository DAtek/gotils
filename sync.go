package gotils

import "sync"

type goroGroup struct {
	wg         *sync.WaitGroup
	goroutines []func()
}

func NewGoroGroup() *goroGroup {
	return &goroGroup{
		wg:         &sync.WaitGroup{},
		goroutines: []func(){},
	}
}

func (g *goroGroup) Add(f func()) {
	g.wg.Add(1)
	g.goroutines = append(g.goroutines, f)
}

func (g *goroGroup) Run() {
	for _, goro := range g.goroutines {
		f := goro

		go func() {
			defer g.wg.Done()
			f()
		}()
	}

	g.wg.Wait()
	g.goroutines = []func(){}
}
