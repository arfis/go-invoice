package util

type TestingFunc func(string, int)

func (f TestingFunc) Test(s string, i int) {
	f(s, i)
}
