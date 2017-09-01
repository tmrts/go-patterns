# Facade Pattern

Facade design pattern provides a unified interface to a set
of interfaces in a subsystem. Facade defines a higher-level
interface that makes the subsystem easier to use

## Implementation

```go
package facade

// Models
type Product struct {
    name string
    cost float64
    discount float64
}

type User struct {
    name string
    discount float64
}

// Storages
type ProductStorage struct {
    storageMap map[string]Product
}

func (storage ProductStorage) GetProductByName(name string) Product{
    product, ok := storage.storageMap[name]
    if (!ok){
        panic("Product is absent")
    }

    return product
}

type UserStorage struct {
    storageMap map[string]User
}

func (storage UserStorage) GetUserByName(name string) User{
    user, ok := storage.storageMap[name]
    if (!ok){
        panic("User is absent")
    }

    return user
}

// Facade
type ProductFacade struct {
    productStorage ProductStorage
    userStorage UserStorage
}

func (productFacade ProductFacade) GetDiscountedCost(userName string, productName string) float64 {
    product := productFacade.productStorage.GetProductByName(productName)
    user := productFacade.userStorage.GetUserByName(userName)

    return product.cost * (1 - product.discount) * (1 - user.discount)
}

```

## Usage

```go
    user := User{name: "Alex", discount: 0.2}
    product := Product{name: "Car", cost:10.50, discount: 0.1 }

    userStorage := UserStorage{storageMap: make(map[string]User)}
    //userStorage.storageMap = make(map[string]User)
    userStorage.storageMap[user.name] = user

    productStorage := ProductStorage{storageMap: make(map[string]Product)}
    //productStorage.storageMap = make(map[string]Product)
    productStorage.storageMap[product.name] = product

    facade := ProductFacade{productStorage: productStorage, userStorage: userStorage}
    discount := facade.GetDiscountedCost("Alex", "Car")
    fmt.Printf("Cost with discount is: %f", discount)
```