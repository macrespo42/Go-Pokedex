# Pokedexcli

A Pokedex-like REPL in Go that use the [ pokeapi ](https://pokeapi.co/) to fetch data about pokemon.

## Quickstart

First you need to compile the binary: 
```bash
go build
```
 and then execute it 
```bash
./pokedexcli
```

Then you can get a list of available command by typing help, get a list of map from the pokemon world or even  
try to catch your first pokemon!

## Usage

There are a lot of available command like map, mapb, explore, inspect catch. You can find all command documentation in this section.

### help 

```pokedexcli> help```: Display a list of all available commands with they description

### map

```pokedexcli> map``` Displays the names of 20 location areas, each consecutive call displays 20 next maps.


### mapb

```pokedexcli> mapb``` Displays the names of 20 previous location areas, each consecutive call displays 20 previous maps.
If you call this but don't have previous map set it just gonna send an empty response


### explore

```pokedexcli> explore canalave-city-area``` Displays pokemon who lives in the area specified as argument. Send an empty response in case of invalid specified area

### catch

```pokedexcli> catch tentacool``` Try to catch the pokemon given as argument, the more the pokemon as a high base experience the harder its gonna be to catch him!  
Send an empty response in case of invalid pokemon name

### inspect

```pokedexcli> inspect tentacool``` Display information about a catched pokemon. If you haven't catch the pokemon yet you can't  
inspect it!

### pokedex

```pokedexcli> pokedex``` List all catched pokemon. If you didn't catch any yet the list will be empty
inspect it!

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/macrespo42/pokedexcli
cd pokedexcli
```

### Build the project

```bash
go build
```

### Run the project

```bash
./pokedexcli
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
