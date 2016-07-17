class: center, middle, inverse

## Streaming Algorithms and Approximate Data Structures

---
## Overview

- Introduction

--

- Bloom Filters

--

- Count-Min Sketch

--

- Reading

---

## Intro

- Many of these problems have exact solutions if you have enough memory and/or CPU

--

- Reduce CPU and memory by approximating

--

- I will skim the math

---

## Terms you should know

--

- Cash Register Model
    - things go in
    - x z x x y w z x y x

--

- Turnstile model
    - things go in and out
    - [x,1], [z,1], [x,2], [y,1], [w,1], [z,-2], [x,-1], [y,-1], [x,-1]

--

- epsilon/delta
    - answer is within epsilon of true value, with probability of failure delta

--

- k independent hash functions
    - murmur3 hash function with random seed
    - other tricks work too: h(x,i) = h1(x) + i * h2(x)
    - universal hashing

---

## What I'm not going to talk about

- HyperLogLog

- TopK

- Streaming Quantiles

---

## Bloom Filters

- B.H. Bloom, 1970

--

- Approximate Set

--

- "No I haven't seen this" or "Yes, I've probably seen this"

--

- Cache heavy queries: disk or network lookup
    - BigTable/Cassandra
    - Chrome's "malicious URL" check
    - Bloom join
    - "Network Applications of Bloom Filters: A Survey" (Broder, Mitzenmacher 2005)

---

### How would we build this?

- One hash table of bits

- One hash function

--

- Do this 'k' times

--

- ~10 bits per element + 5 hash functions give you <1% false positive rate

--
- "Given n elements and false-positive rate p, how many bits do I need?"
    - m = (-n ln(p)) / ( ln(2)^2 )

--

- "Given m bits of storage and n elements, how many hash functions do I need?"
    - k = m/n * ln(2)

---
class: center, middle

### Wikipedia

![bf](Bloom_filter.svg)

---

### Bloom Filters (cont)

--

- "What have I put in the set?"

--

- "How many items have I put in the set?"

--

- "How do I remove elements from the set?"

--

- "How many items can I put in the set?"

--

- Union, Intersection, Halving

---

### Count-Min Sketch

--

- Approximate Frequencies
    - "How many times have I seen this?"

--

- "How would we build this?"
    - one hash table of counters
    - one hash function
    - "do it k times"

--

- "How do we query this?"
    - collisions means the buckets are biased estimators
    - but they're all upper bounds
    - take the minimum across all the buckets

---
class: center, middle

![cm](count-min-sketch.png)

w = ceil(E/epsilon)

d = ceil(ln(1/delta))

estimate <= count + eps * N

eps = 0.001, delta = 0.001 => w=2719, d=7, 32-bit counters is 73k of space


---

### Count Min Sketch (cont)

--

- "What have I put in the set?"

--

- "How many items have I put in the set?"

--

- "How do I remove elements from the set?"

--

- Union, Intersection, Halving

---

## Count-Min Sketch: Applications

- Count tracking on large data sets

- Heavy Hitters
    - Elephants from Mice
    - Count-min sketch + heap
    - Will see another (simpler, magic) algorithm later

- Variations:
    - better low-frequency estimates (CountMeanMin),
        - (but larger under-estimation error for large items)
    - better cash-register estimates (Conservative Update)
        - (but deleting isn't allowed)

- Many many others

---

### More Stuff I didn't talk about

- Multi-pass streaming algorithms

- Sliding windows

---

### Links

- https://github.com/dgryski
    - libcmsketch (p5-cmsketch)
    - hokusai

- http://research.neustar.biz/

- https://gist.github.com/debasishg/8172796

- https://highlyscalable.wordpress.com/2012/05/01/probabilistic-structures-web-analytics-data-mining/

---

class: center, middle, inverse

## Questions?

---

class: center, middle, inverse

## fin

???

vim: ft=markdown
