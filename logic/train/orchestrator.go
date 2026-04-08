package train

import "github.com/Skoowshot/vecspect/domain"

type TrainingOrchestrator struct {
	WorkerCount  int
	WorkerQueues []chan domain.TrainingMessage

	Tokenizer *Tokenizer
}

func NewOrchestrator(workerCount int, workerQueueSize int) *TrainingOrchestrator {
	queues := make([]chan domain.TrainingMessage, workerCount)
	for i := 0; i < workerCount; i++ {
		queues[i] = make(chan domain.TrainingMessage, workerQueueSize)
	}

	orchestrator := &TrainingOrchestrator{
		WorkerCount:  workerCount,
		WorkerQueues: queues,

		Tokenizer: NewTokenizer(),
	}

	orchestrator.Start()

	return orchestrator
}

func (o *TrainingOrchestrator) Start() {
	worker := NewWorker(o)
	
	for i := 0; i < o.WorkerCount; i++ {
		go worker.Start(i)
	}
}

func (o *TrainingOrchestrator) Push(msg domain.TrainingMessage) {
	chatId := uint64(msg.ChatId)
	maxWorkers := uint64(o.WorkerCount)

	workerIdx := chatId % maxWorkers

	o.WorkerQueues[workerIdx] <- msg
}
