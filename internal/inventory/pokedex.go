package inventory

type Pokemon struct {
	Name      string
	Height    int
	Weight    int
	Stats     map[string]int
	Abilities []string
	Types     []string
}

var pokemonCollection map[string]Pokemon

func AddPokemon(data Pokemon) {
	if pokemonCollection == nil {
		pokemonCollection = make(map[string]Pokemon)
	}
	pokemonCollection[data.Name] = data
}

func ListPokemon() []string {
	var names []string
	for k := range pokemonCollection {
		names = append(names, k)
	}

	return names
}

func GetPokemon(name string) (Pokemon, bool) {
	p, ok := pokemonCollection[name]

	return p, ok
}
