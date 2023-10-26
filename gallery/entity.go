package gallery

import (
	"time"
)

type Gallery struct {
	ID    int
	Image string
	Alt   string
	//Likes int
	CreatedAt time.Time
	UpdatedAt time.Time
}
