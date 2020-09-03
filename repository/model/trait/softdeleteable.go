package trait

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

// DO NOT USE!
type Softdeleteable struct {
	DeletedAt *time.Time `gorm:"column:created_at;null;" json:"-"`
}

//
func (this *Softdeleteable) BeforeDelete(scope *gorm.Scope) error {
	scope.SetColumn("DeletedAt", time.Now().Unix())
	return errors.New("error.soft_deleteable")
}

// TODO: Not working like this!
func (this *Softdeleteable) AfterFind() (err error) {
	if this.DeletedAt != nil {
		*this = Softdeleteable{}
		this = nil
	}
	return
}
