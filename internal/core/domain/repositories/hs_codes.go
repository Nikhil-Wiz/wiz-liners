package repositories

import(
	"time"
)

type Hs_codes struct{
	Code string
	Name string
	Description string
	Parent_code string
	CreatedAt time.Time
	ModifiedAt time.Time
}