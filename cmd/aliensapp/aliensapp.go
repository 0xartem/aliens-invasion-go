package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"

    "github.com/artem-braznikov/aliens-invasion-go/pkg/alieninvasion"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type ConsoleObserver struct {
    printAliensLeft bool
}

func (c *ConsoleObserver) OnCityOutput(city *alieninvasion.City) {
    if c.printAliensLeft {
        fmt.Println(city.StringWithAlien())
    } else {
        fmt.Println(city.String())
    }
}

func (c *ConsoleObserver) OnAliensDestroyedCity(fstAl *alieninvasion.Alien, secAl *alieninvasion.Alien, city *alieninvasion.City) {
    fmt.Printf("%s has been destroyed by alien %d and alien %d", city.Name(), fstAl.Id(), secAl.Id())
    fmt.Println()
}

func main() {
    argsCount := 3
    if len(os.Args) != argsCount {
        fmt.Println("You must provide a path to the world's map and amount of aliens. Usage: aliens-invasion map.txt N")
        os.Exit(1)
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        check(err)
    }
    defer file.Close()

    consoleObserver := ConsoleObserver{true}
    mapManager := alieninvasion.NewMapManager()
    aliensManager := alieninvasion.NewAliensManager(mapManager, &consoleObserver)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        mapManager.BuildCity(scanner.Text())
    }
    check(scanner.Err())

    aliensCount, err := strconv.Atoi(os.Args[2])
    check(err)

    for i := 0; i < aliensCount; i++ {

        created, err := aliensManager.AddNewAlien(i)
        check(err)

        // If an alien was not created - the map became empty
        if !created {
            break
        }
    }

    for i := 0; i < 10000; i++ {

        if aliensManager.AreAllDestroyed() {
            break
        }

        err := aliensManager.MoveAll()
        check(err)
    }

    if mapManager.IsMapEmpty() {
        fmt.Println("The map is empty!")
    } else {
        mapManager.OutputAll(&consoleObserver)
    }
}
