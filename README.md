# MapData

A package to help traversing a nested map.

This is helpful when unmarshalling data from a yaml or json file to a variable.
We convert this variable to an "object" and have the function to retrieve a value 
by providing the path to value string. 


## Usage

```
    mapTest := map[string]interface{}{
        "a": map[string]interface{}{
          "a1": "alpha",
        },
        "b": "beta",
        "c": "charlie",
        "d": "delta",
    }
    myMap, err := NewMapData(mapTest)
    if err != nil {
        log.Fatal(err)
    }

    v, err := myMap.GetPath("a.a1")
    # v is `alpha`
    # err is `nil`
```
