package train

import "github.com/Skoowshot/vecspect/domain"

type TrainingOrchestrator struct {
	WorkerCount  int
	WorkerQueues []chan domain.TrainingMessage
}

func NewOrchestrator(workerCount int, workerQueueSize int) *TrainingOrchestrator {
	queues := make([]chan domain.TrainingMessage, workerCount)
	for i := 0; i < workerCount; i++ {
		queues[i] = make(chan domain.TrainingMessage, workerQueueSize)
	}

	orchestrator := &TrainingOrchestrator{
		WorkerCount:  workerCount,
		WorkerQueues: queues,
	}

	orchestrator.Start()

	return orchestrator
}

func (o *TrainingOrchestrator) Start() {
	worker := func(id int) {
		for msg := range o.WorkerQueues[id] {
			println("worker", id, ">", msg.Original, msg.Reply)
		}
	}

	for i := 0; i < o.WorkerCount; i++ {
		go worker(i)
	}
}

func (o *TrainingOrchestrator) Push(msg domain.TrainingMessage) {
	chatId := uint64(msg.ChatId)
	maxWorkers := uint64(o.WorkerCount)

	workerIdx := chatId % maxWorkers

	o.WorkerQueues[workerIdx] <- msg
}
