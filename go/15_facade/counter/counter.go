package counter

type counter struct {}

func (c *counter) count(list []string) map[string]int {
	m := map[string]int{}
	for _, word := range list {
		m[word]++
	}
	return m
}