package schedule

type Strategy interface {
	Choose(urls []string) string
}