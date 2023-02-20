package funfucts

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	funFacts := new(FunFacts)

	funFacts.Sun = []string{"test", "test"}
	funFacts.Luna = []string{"test1", "test11"}
	funFacts.Terra = []string{"test2", "test22"}

	if about == "sun" {
		return funFacts.Sun
	} else if about == "luna" {
		return funFacts.Luna
	} else {
		return funFacts.Terra
	}

}
