package logic

type LearnPair struct {
	ChatId   int64
	Original string
	Reply    string
}

type LearnManager struct {
	workerQueues []chan LearnPair
	workerCount  int
}

func NewPair(chatId int64, original, reply string) LearnPair {
	return LearnPair{
		ChatId:   chatId,
		Original: original,
		Reply:    reply,
	}
}

func NewLearnManager(workerCount int, queueLength int) *LearnManager {
	queues := make([]chan LearnPair, workerCount)
	for i := 0; i < workerCount; i++ {
		queues[i] = make(chan LearnPair, queueLength)
	}

	lm := &LearnManager{
		workerQueues: queues,
		workerCount:  workerCount,
	}
	lm.startWorkers()
	return lm
}

func (lm *LearnManager) startWorkers() {
	for i := 0; i < lm.workerCount; i++ {
		go func(id int) {
			for pair := range lm.workerQueues[id] {
				println("worker", id, ">", pair.Original, pair.Reply)
			}
		}(i)
	}
}

func (lm *LearnManager) Push(pair LearnPair) {
	workerIdx := uint64(pair.ChatId) % uint64(lm.workerCount)
	lm.workerQueues[workerIdx] <- pair
}
