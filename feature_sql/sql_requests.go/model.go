package sqlrequestsgo

import "time"

type BookModel struct {
	ID             int
	Name           string
	Author         string
	Resolution     string
	Reales_year    int
	Completed      bool
	Time_on_shelf  time.Time
	Time_completed *time.Time
}
