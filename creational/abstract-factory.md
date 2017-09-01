# Abstract Factory Pattern

Abstract Factory creational design pattern provides an interface for 
creating families of related or dependent objects without specifying
their concrete classes.

## Implementation

```go
package afactory

import "fmt"

// Coin hierarchy
type Coin interface {
    name() string
}

type Etherium struct {
}

type Bitcoin struct {
}

func (c Etherium) name() string {
    return "Etherium"
}

func (c Bitcoin) name() string {
    return "Bitcoin"
}

// Abstract Factory implementation
type AbstractCoinFactory interface {
    createCoin() Coin
}

type EtheriumFactory struct {
}

type BitcoinFactory struct {
}

func (c EtheriumFactory) createCoin() Coin {
    return Etherium{}
}

func (c BitcoinFactory) createCoin() Coin {
    return Bitcoin{}
}

```

## Usage

```go

func printNewCoinName (factory AbstractCoinFactory) {
    coin := factory.createCoin()
    fmt.Printf("%s\n", coin.name())
}

printNewCoinName(EtheriumFactory{})
printNewCoinName(BitcoinFactory{})

```