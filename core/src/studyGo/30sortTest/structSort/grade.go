package main

type grade struct {
	chinese uint32
	math uint32
	english uint32
}

type grades []*grade

func (s grades) Len() int { return len(s) }

func (s grades) Swap(i, j int){ s[i], s[j] = s[j], s[i] }

func (s grades) Less(i, j int) bool {
	if s[i].chinese == s[j].chinese{
		return s[i].english > s[j].english
	}
	return s[i].chinese > s[j].chinese
}

