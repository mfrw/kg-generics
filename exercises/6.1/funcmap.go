package funcmap

// Your implementation of FuncMap goes here!

type FuncMap[I any, O any] map[string]func(i I) O

func (fm FuncMap[I, O]) Apply(fn string, i I) O {
	var def O
	if f, ok := fm[fn]; ok {
		return f(i)
	}
	return def
}
