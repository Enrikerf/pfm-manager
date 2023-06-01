package Result

import (
	"github.com/Enrikerf/pfm/commandManager/app/Domain/Result/Content"
	"time"
)

type Result interface {
	GetId() Id
	GetBatchId() BatchId
	GetContent() Content.Content
	GetCreateAt() time.Time
}

func New(batchId BatchId, content Content.Content) Result {
	result := &result{}
	result.id = NewId()
	result.batchId = batchId
	result.content = content
	result.createdAt = time.Now()
	return result
}

func Load(
	id Id,
	batchId BatchId,
	content Content.Content,
	createdAt time.Time,
) Result {
	return &result{
		id,
		batchId,
		content,
		createdAt,
	}
}

type result struct {
	id        Id
	batchId   BatchId
	content   Content.Content
	createdAt time.Time
}

func (r *result) GetId() Id {
	return r.id
}

func (r *result) GetBatchId() BatchId {
	return r.batchId
}

func (r *result) GetContent() Content.Content {
	return r.content
}

func (r *result) GetCreateAt() time.Time {
	return r.createdAt
}
