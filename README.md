Ntram Indexing
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/ntram/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/ntram?status.svg)](https://godoc.org/github.com/kkdai/ntram)  [![Build Status](https://travis-ci.org/kkdai/ntram.svg?branch=master)](https://travis-ci.org/kkdai/ntram)


This package provide a simple way to "Trigram Indexing" in input document. It is refer from an article - [Google Code Search](https://github.com/google/codesearch).

Here is the [introduction](http://www.evanlin.com/trigram-study-note/) what is "trigram indexing" and how Google Code Search use it for search but it is in Chinese :) .


How it works
---------------

This package using [trigram indexing](https://swtch.com/~rsc/regexp/regexp4.html) to get all trigram in input string (what we call document).

Here is some trigram rule as follow:

- It will not transfer Upper case	 to Lower case. (follow code search rule)
- Includes "space"

 
Install
---------------
`go get github.com/kkdai/trigram`


Usage
---------------

```go

package main

import (
	"fmt"
	. "github.com/kkdai/trigram"
	)
func main() {	
	ti := NewTrigramIndex()
	ti.Add("Code is my life")			//doc 1
	ti.Add("Search")						//doc 2
	ti.Add("I write a lot of Codes") //doc 3
	
	//Print all trigram map 
	fmt.Println("It has ", len(ti.TrigramMap))
	for k, v := range ti.TrigramMap {
		fmt.Println("trigram=", k, " obj=", v)
	}

	//Search which doc include this code
	ret := ti.Query("Code")
	fmt.Println("Query ret=", ret)
	// [1, 3]
}
```


Benchmark
---------------

Still working to improve the query time.

```
//Original with Map
BenchmarkAddTwogram-4    	  200000	      7151 ns/op
BenchmarkAddTrigram-4    	  300000	      6713 ns/op
BenchmarkAddFourgran-4   	  300000	      5813 ns/op
BenchmarkDeleteTwogram-4 	  500000	      4591 ns/op
BenchmarkDeleteTrigram-4 	  500000	      3695 ns/op
BenchmarkDeleteFourgram-4	  500000	      3297 ns/op
BenchmarkQueryTwogran-4  	   10000	   8361813 ns/op
BenchmarkQueryTrigran-4  	   10000	   7650419 ns/op
BenchmarkQueryFourgram-4 	   10000	   6975925 ns/op
BenchmarkIntersectMap-4  	     200	   7834764 ns/op
BenchmarkIntersectSlice-4	     500	   2986213 ns/op


//Using slice in docIDs
BenchmarkAddTwogram-4    	   30000	    162352 ns/op
BenchmarkAddTrigram-4    	   30000	    142771 ns/op
BenchmarkAddFourgran-4   	   30000	    130655 ns/op
BenchmarkDeleteTwogram-4 	   20000	    169485 ns/op
BenchmarkDeleteTrigram-4 	   20000	    147598 ns/op
BenchmarkDeleteFourgram-4	   20000	    129833 ns/op
BenchmarkQueryTwogran-4  	   10000	   2295375 ns/op
BenchmarkQueryTrigran-4  	   10000	   2161023 ns/op
BenchmarkQueryFourgram-4 	   10000	   2032144 ns/op
BenchmarkIntersectMap-4  	     200	   8216303 ns/op
BenchmarkIntersectSlice-4	    1000	   2245023 ns/op

//Operate in map but intersect in slice
BenchmarkAddTwogram-4    	  200000	      7766 ns/op
BenchmarkAddTrigram-4    	  200000	      6313 ns/op
BenchmarkAddFourgran-4   	  300000	      6302 ns/op
BenchmarkDeleteTwogram-4 	  500000	      4743 ns/op
BenchmarkDeleteTrigram-4 	  500000	      3854 ns/op
BenchmarkDeleteFourgram-4	  500000	      3317 ns/op
BenchmarkQueryTwogran-4  	   10000	  28069290 ns/op
BenchmarkQueryTrigran-4  	   10000	  24906564 ns/op
BenchmarkQueryFourgram-4 	   10000	  20321440 ns/op
BenchmarkIntersectMap-4  	     200	   8216303 ns/op
BenchmarkIntersectSlice-4	    1000	   2245023 ns/op

```

BTW: Here is benchmark for [https://github.com/dgryski/go-trigram](https://github.com/dgryski/go-trigram) for my improvement record:



```
BenchmarkAdd-4       1000000          1063 ns/op
BenchmarkDelete-4     100000        140392 ns/op
BenchmarkQuery-4       10000        474320 ns/op
```

Inspired
---------------

- [Google Code Search (using Go)](https://github.com/google/codesearch)
- [Trigram Algorithm](http://ii.nlm.nih.gov/MTI/Details/trigram.shtml)
- [https://github.com/dgryski/go-trigram](https://github.com/dgryski/go-trigram)
- [Regular Expression Matching with a Trigram Index](https://swtch.com/~rsc/regexp/regexp4.html)
- [Approximate string-matching algorithms, part 2](http://www.morfoedro.it/doc.php?n=223&lang=en#SimilarityMetric)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

