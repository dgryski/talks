class: center, middle, inverse

## dmrgo

---
## Overview

- MapReduce

- dmrgo

---

## MapReduce

- MapReduce: Simplified Data Processing on Large Clusters (Dean and Ghemawat, 2004)

- Google has moved beyond it.

- Hadoop hasn't.


---
## WordCount (input)

```
hello howdy world everybody hello everybody howdy howdy
```

---
## WordCount (map)

```
hello 1
howdy 1
world 1
everybody 1
hello 1
everybody 1
howdy 1
howdy 1
```

---
## WordCount (shuffle)

```
everybody 1
everybody 1
hello 1
hello 1
howdy 1
howdy 1
howdy 1
world 1
```

---
## WordCount (reduce)

```
everybody 2
hello 2
howdy 3
world 1
```

---
## Streaming Hadoop

- stdin/stdout

- TSV

---
## dmrgo

- statically typed Java API

- dynamically typed Python library mrjob ( https://github.com/Yelp/mrjob )

---

```go

type Emitter interface {
    Emit(key string, value string)
    Flush()
}

type MapReduceJob interface {
    Map(key string, value string, emitter Emitter)

    // Called at the end of the Map phase
    MapFinal(emitter Emitter)

    Reduce(key string, values []string, emitter Emitter)
}
```

---

class: center, middle, inverse

## Questions?

---

class: center, middle, inverse

## fin

???

vim: ft=markdown
