package utils

type Cache struct {
	TimeTable *TimeTable
}

func InitCache() (c Cache) {
	return Cache{TimeTable: new(TimeTable)}
}
