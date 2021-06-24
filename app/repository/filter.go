package repository

type ActivityProgramOrderBy int

const (
	ActivityProgramOrderByDatetimeAsc = iota + 1
)

type ActivityProgramFilter struct {
	OrderBy ActivityProgramOrderBy
}
