Ngram Indexing
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/ngram/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/ngram?status.svg)](https://godoc.org/github.com/kkdai/ngram)  [![Go](https://github.com/kkdai/ngram/actions/workflows/go.yml/badge.svg)](https://github.com/kkdai/ngram/actions/workflows/go.yml)

This package provide a simple way to "Nrigram Indexing" in input document. It is refer from an article - [Google Code Search](https://github.com/google/codesearch).


Here is the [introduction](http://www.evanlin.com/trigram-study-note/) what is "trigram indexing" and how Google Code Search use it for search but it is in Chinese :) .


Performance Optimization
---------------

This package base on my another project [Trigram](https://github.com/kkdai/trigram), but it got better performance(**~3 times**). (refer Benchmark)

It has done as follow:

- Replace operation data in slice rather than map in intersect operation
- But still remain using map in query


How it works
---------------

This package using [trigram indexing](https://swtch.com/~rsc/regexp/regexp4.html) to get all trigram in input string (what we call document).

Here is some trigram rule as follow:

- It will not transfer Upper case	 to Lower case. (follow code search rule)
- Includes "space"

 
Install
---------------
`go get github.com/kkdai/ngram`


Usage
---------------

```go

package main

import (
	"fmt"
	. "github.com/kkdai/ngram"
	)
func main() {	
	//Currently is support Twogram, Trigram and Fourgram
	ti := NewNgramIndex(Trigram)
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
//Original benchmark in trigram
BenchmarkAddTwogram-4    	  200000	      7151 ns/op
BenchmarkAddTrigram-4    	  300000	      6713 ns/op
BenchmarkAddFourgran-4   	  300000	      5813 ns/op
BenchmarkDeleteTwogram-4 	  500000	      4591 ns/op
BenchmarkDeleteTrigram-4 	  500000	      3695 ns/op
BenchmarkDeleteFourgram-4	  500000	      3297 ns/op
BenchmarkQueryTwogran-4  	   10000	   8361813 ns/op
BenchmarkQueryTrigran-4  	   10000	   7650419 ns/op
BenchmarkQueryFourgram-4 	   10000	   6975925 ns/op


//Optimize result
BenchmarkAddTwogram-4    	  300000	      5737 ns/op
BenchmarkAddTrigram-4    	  500000	      4795 ns/op
BenchmarkAddFourgran-4   	  500000	      4158 ns/op
BenchmarkDeleteTwogram-4 	   20000	    167246 ns/op
BenchmarkDeleteTrigram-4 	   20000	    148756 ns/op
BenchmarkDeleteFourgram-4	   20000	    128022 ns/op
BenchmarkQueryTwogran-4  	   10000	   2461910 ns/op
BenchmarkQueryTrigran-4  	   10000	   2276625 ns/op
BenchmarkQueryFourgram-4 	   10000	   2172323 ns/op
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

