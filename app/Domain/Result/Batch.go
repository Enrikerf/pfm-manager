package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Task"
	"time"
)

type Batch interface {
	GetId() BatchId
	GetTaskId() Task.Id
	GetResults() []Result
	GetCreatedAt() time.Time
	AddResult(result Result)
	SetResultsFromContent(results []Content.Content)
}

func NewBatch(taskId Task.Id) Batch {
	batch := &batch{}
	batch.id = NewId()
	batch.taskId = taskId
	batch.results = []Result{}
	batch.createdAt = time.Now()
	return batch
}
func LoadBatch(
	id BatchId,
	taskId Task.Id,
	results []Result,
	createdAt time.Time,
) Batch {
	return &batch{
		id,
		taskId,
		results,
		createdAt,
	}
}

type batch struct {
	id        BatchId
	taskId    Task.Id
	results   []Result
	createdAt time.Time
}

func (b *batch) GetId() BatchId {
	return b.id
}

func (b *batch) GetTaskId() Task.Id {
	return b.taskId
}

func (b *batch) GetResults() []Result {
	return b.results
}

func (b *batch) GetCreatedAt() time.Time {
	return b.createdAt
}

func (b *batch) AddResult(result Result) {
	b.results = append(b.results, result)
}

func (b *batch) SetResultsFromContent(contents []Content.Content) {
	for _, content := range contents {
		b.results = append(b.results, New(b.GetId(), content))
	}
}
