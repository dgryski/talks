class: center, middle, inverse

## Streaming Algorithms and Approximate Data Structures

## Part II: Heavy Hitters and Cardinality Estimation

---
## Overview

- TopK

- HyperLogLog

---

## TopK

- What are the TopK (for some constant 'k') elements in the stream

--

- Heavy Hitters
    - what are all elements in the stream with frequency > phi * N for some constant phi

    - no elements with frequency < (phi - eps) * N

    - difference between estimated and true frequency is <= eps * N

---

## TopK Algorithms

- Count-min sketch + heap

- Sampling:  O(1/eps^2)

- "Space Saving" (2005)

- "Filtered Space Saving" (2010)

---

## Space Saving

--

- Set k=ceil(1/eps)

--

- Keep exact (key, count) pairs for first k elements

--

- When a new element arrives, if it's not being tracked, remove the least frequent item and replace it ...

--

- ... with (newkey, *oldcount+1*)

---

## Filtered Space Saving

- Better estimates the error associated with each value

--

- SS: (newkey, oldcount+1, error=oldcount)

--

- maintain CM sketch with d=1, w=6k for estimates

- FSS: (newkey, hash[newkey]+1, error=hash[newkey])

---

## HyperLogLog

- Cardinality: how many distinct items have I seen?

- "how many leading 0s did I see?"

- At most n zero bits => cardinality ~2^n

--

- "Do this k times" (or k-times esitmate 1/kth of the stream)

--

- We don't need k hash functions
    - value only needs to count leading zeros == 5 bits
    - split a single 32-bit hash value into register + value

- estimate = harmonic mean of registers * correction factor

--

- error is 1.04/sqrt(registers)

- Redis uses 12k to store 16k registers => 0.81%

---

## Google: HyperLogLog++

--

- "Counting billions items isn't cool.  You know what's cool?  Counting *trillions* of items."

--

- Brute force improved correction factors

- space optimizations

---

## Reading

SS:
    - https://icmi.cs.ucsb.edu/research/tech_reports/reports/2005-23.pdf
    - http://www.l2f.inesc-id.pt/~fmmb/wiki/uploads/Work/misnis.ref0a.pdf

HLL:
    - http://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf
    - http://research.google.com/pubs/pub40671.html
    - http://antirez.com/news/75
    - http://research.neustar.biz/2013/01/24/hyperloglog-googles-take-on-engineering-hll/
    - http://druid.io/blog/2014/02/18/hyperloglog-optimizations-for-real-world-systems.html

---

class: center, middle, inverse

## Questions?

---

class: center, middle, inverse

## fin

???

vim: ft=markdown
