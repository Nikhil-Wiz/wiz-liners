package repositories

import(
	"time"
)

type Hs_codes struct{
	Code string
	Name string
	Description string
	Parent_code string
	Created_At time.Time
	Modified_At time.Time
}