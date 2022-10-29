package counter

type WordCounter struct {
	counter
	spliter
	lower
}

func NewWordCounter(sep string) *WordCounter {
	return &WordCounter{
		counter: counter{},
		spliter: spliter{sep: sep},
		lower: lower{},
	}
}

// 複雑な処理を行うAPIを提供する。
// このクラスを利用することで、単語のカウントが容易に行える。
// その代わりに小文字化や文字列分割の機能は利用できない。
// →機能性とシンプルさのトレードオフ
func (c *WordCounter) Count(sentence string) map[string]int {
	lower_sentence := c.lower.lower(sentence)
	word_list := c.spliter.split(lower_sentence)
	counts := c.counter.count(word_list)

	return counts
}