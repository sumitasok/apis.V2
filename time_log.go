package apis

import (
	"time"
)

type TimeLog struct {
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

func (t *TimeLog) SetCreatedAt(ti ...time.Time) {
	if len(ti) == 0 {
		t.CreatedAtNow()
		return
	}

	t.CreatedAt = ti[0]
}

func (t *TimeLog) UpdatedNow() {
	t.UpdatedAt = time.Now()
}

func (t *TimeLog) CreatedAtNow() {
	t.CreatedAt = time.Now()
}

func (t *TimeLog) Deleted() {
	t.DeletedAt = time.Now()
}
