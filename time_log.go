package apis

import (
	"time"
)

// TimeLog a default time logging code for mongodb
// which takes care of time based basic details of CRUD operations
type TimeLog struct {
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}

// SetCreatedAt sets created at time
func (t *TimeLog) SetCreatedAt(ti ...time.Time) {
	if len(ti) == 0 {
		t.CreatedAtNow()
		return
	}

	t.CreatedAt = ti[0]
}

// UpdatedNow updates the updated at time
func (t *TimeLog) UpdatedNow() {
	t.UpdatedAt = time.Now()
}

// CreatedAtNow sets the created at as current time
func (t *TimeLog) CreatedAtNow() {
	t.CreatedAt = time.Now()
}

// Deleted sets the deleted at as current time
func (t *TimeLog) Deleted() {
	t.DeletedAt = time.Now()
}
