package schedule

type Strategy interface {
	Init()
	Choose(client string, servers []string) string
}