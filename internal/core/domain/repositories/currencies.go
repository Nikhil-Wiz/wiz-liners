package repositories

import(
	"time"
)

type Currencies struct{
	Code string
	Name string
	Created_At time.Time
	Modified_At time.Time
}