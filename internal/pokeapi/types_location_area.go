package pokeapi

type Map struct {
	Name string
	URL  string
}

type LocationArea struct {
	Count    int
	Next     string
	Previous string
	Results  []Map
}
