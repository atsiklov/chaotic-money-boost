package template

type ChallengeTemplate struct {
	ID          int64  `db:"id"`
	CategoryID  int64  `db:"challenge_category_id"`
	Description string `db:"description"`
}
