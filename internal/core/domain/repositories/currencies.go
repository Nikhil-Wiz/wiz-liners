package repositories

import(
	"time"
)

type Currencies struct{
	Code string
	Name string
	CreatedAt time.Time
	ModifiedAt time.Time
}