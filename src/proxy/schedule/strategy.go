package schedule

type Strategy interface {
	Choose(client string, servers []string) string
}