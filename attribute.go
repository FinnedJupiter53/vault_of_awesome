package main

type attributes struct {
	Str attrib
	Dex attrib
	Con attrib
	Int attrib
	Wis attrib
	Cha attrib
}

type attrib struct {
	Value int
}

func newAttributes() attributes {
	a := attributes{
		Str: randomStat(),
		Dex: randomStat(),
		Con: randomStat(),
		Int: randomStat(),
		Wis: randomStat(),
		Cha: randomStat(),
	}

	return a
}

func randomStat() attrib {
	a := attrib{
		Value: die(6) + die(6) + die(6),
	}

	return a
}

func (a attrib) Modifier() int {
	return a.Value/2 - 5
}
