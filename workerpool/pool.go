package workerpool

import (
	"sync"
)

type Pool struct {
	Tasks   []*Task
	Workers []*Worker

	concurrency   int
	collector     chan *Task
	runBackground chan bool
	wg            sync.WaitGroup
}

func (p *Pool) AddTask(task *Task) {
	p.collector <- task
}

func (p *Pool) RunBackground() {
	for i := 1; i <= p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.StartBackground()
	}

	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}
	<-p.runBackground
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:         tasks,
		concurrency:   concurrency,
		collector:     make(chan *Task, 1000),
		runBackground: make(chan bool),
	}
}

func (p *Pool) Stop() {
	for i := range p.Workers {
		p.Workers[i].Stop()
	}

	p.runBackground <- true
}
