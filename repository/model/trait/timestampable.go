package trait

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Add a created_at and updated_at field to the entity and update with every change
type Timestampable struct {
	CreatedAt time.Time  `gorm:"column:created_at;not null;DEFAULT:current_timestamp"`
	UpdatedAt *time.Time `gorm:"column:updated_at;null;DEFAULT:current_timestamp"`
}

// Automatically update the updated_at field before updating the entry
func (this *Timestampable) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
