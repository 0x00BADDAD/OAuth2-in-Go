package main

import (
    "fmt"
    "mymodule/hello-world/src/generic"
    )




func main() {
    // fmt.Println("Hello, I am in the main! in basketballLeague")
    // var sampleData basketball.League = basketball.League{
    //    Teams :[]string{"CSK", "MI", "RCB"},
    //    Wins: map[string]int{
    //        "CSK": 0,
    //        "MI": 0,
    //        "RCB": 0,
    //    },
    //}
    //sampleData.MatchResult("CSK", 2, "MI", 4)
    //sampleData.MatchResult("MI", 5, "RCB", 3)
    //sampleData.MatchResult("RCB", 4, "CSK", 1)
    //var winOrder []string = sampleData.Ranking()
    //basketball.RankPrinter(sampleData, os.Stdout)
    //fmt.Printf("\nThe winOrder is %v\n", winOrder)

    //var initialVal int = 10
    //var initialVal_ float64 = 5.56
    //val1, val2 := generic.Double(initialVal), generic.Double(initialVal_)
    //fmt.Printf("The values have been doubled! %v, %v\n", val1, val2)
    //var printableInt generic.PrintableInt = 35
    //var printablefloat64 generic.Printablefloat64 = 18.36

    //generic.PrintToScreen(printableInt)
    //generic.PrintToScreen(printablefloat64)
    list := generic.Node[int]{
        Val: 1,
        Next: nil,
    }
    list.Add(3)
    list.Add(5)
    is3 := list.Contains(3)
    is5 := list.Contains(5)
    is2 := list.Contains(2)
    fmt.Printf("is3: %v, is5: %v, is2: %v\n", is3, is5, is2)

    listStrings := generic.Node[string]{
        Val: "Starts here!",
        Next: nil,
    }
    listStrings.Add("Idk lol!")
    listStrings.Add("Yup still here baby!")
    listStrings.PrettyPrinter()
    fmt.Printf("\n")
    var inserted bool = listStrings.Insert(1, "Inserted String")
    if inserted {
        listStrings.PrettyPrinter()
    }

}

