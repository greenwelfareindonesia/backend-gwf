package endpointcount

import (
	"gorm.io/gorm"
)

type Statistics struct {
	gorm.Model
	Endpoint  string
	Count     int
	UserAgent string 
	UniqueUserAgent int
}

