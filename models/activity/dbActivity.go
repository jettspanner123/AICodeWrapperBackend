package activity

import (
	"time"

	enums "github.com/jettspanner123/AICodeWrapperBackend/models/activity/enums"
)

type APILog struct {
	ID        uint                 `gorm:"primaryKey"`
	Name      string               `gorm:"type:varchar(100);not null;index"`
	Category  enums.APILogCategory `gorm:"type:varchar(100);not null;index"`
	Method    enums.APIMethodTypes `gorm:"type:varchar(100);not null;index"`
	CreatedAt time.Time
}
