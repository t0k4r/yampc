package main

func maps[I, O any](in []I, fn func(I) O) []O {
	out := []O{}
	for _, i := range in {
		out = append(out, fn(i))
	}
	return out
}
