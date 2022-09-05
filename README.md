# tagLevenshtein
Get taglsit levenshtein for Golang.

-----

## Install

```
go get github.com/ArsFy/tagLevenshtein@v1.0.0
```

## Example

```go
package main

import (
    "github.com/ArsFy/tagLevenshtein"
)

func main(){
    var list1 []string = []string{"ice", "apple", "book", "test"}
	var list2 []string = []string{"ice", "apple", "test", "bird"}
	fmt.Println(levenshtein(list1, list2))
}
```