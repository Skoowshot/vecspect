package train

import (
	"log"
	"strings"

	"github.com/Skoowshot/vecspect/domain"
)

type Worker struct {
	orchestrator *TrainingOrchestrator
}

func NewWorker(orchestrator *TrainingOrchestrator) *Worker {
	return &Worker{
		orchestrator: orchestrator,
	}
}

func (w *Worker) Start(id int) {
	for msg := range w.orchestrator.WorkerQueues[id] {
		if msg.Reply != nil {
			w.HandleReply(id, msg)
		} else {
			w.HandleDefault(id, msg)
		}
	}
}

func (w *Worker) HandleDefault(id int, msg domain.TrainingMessage) {
	tokenized := w.orchestrator.Tokenizer.Tokenize(msg.Original)

	log.Printf("[worker %v] tokenizer - %v", id, strings.Join(tokenized, "|"))
}

func (w *Worker) HandleReply(id int, msg domain.TrainingMessage) {

}
