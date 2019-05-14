class: center, middle, inverse

## randomized testing for go

---

## randomized testing

- `testing/quick`

- `github.com/dvyukov/go-fuzz`

- plus some others

---

# history

- University of Wisconsin Madison in 1989 by Professor Barton Miller and his students

- "An Empirical Study of the Reliability of UNIX Utilities" (1990) and "Fuzz Revisited: A Re-examination of the Reliability of UNIX Utilities and Services" (1995)

- some bugs reported in 1990 were still present in 1995

- mid-2000s picked up by the security community

---

# why we care

- writing tests is boring

- humans write biased tests

- have the computer write them for you

---

# type of random testing

- property-based testing

- generational fuzzing

- mutational fuzzing

- stateful testing

---

# testing/quick

```go
func TestQuick(t *testing.T) {
    q := func(i, j int) bool {
        quo, rem := Div(i, j)
        return i == quo*j+rem
    }

    if err := quick.Check(q, nil); err != nil {
        t.Error(err)
    }
}
```

---

# testing/quick

```go
func TestQuick(t *testing.T) {
    q := func(i, j int) bool {
        quo, rem := Div(i, j)
        return i == quo*j+rem
    }

    if err := quick.Check(q, nil); err != nil {
        err := err.(*quick.CheckError)
        t.Errorf("iteration %d failed: args=(%v,%v)",
            err.Count, err.In[0], err.In[1])
    }
}
```

---

# testing/quick

```go
func TestQuick(t *testing.T) {
    q := func(i, j int) bool {
        quo, rem := Div(i, j)
        if got := quo*j+rem; got != i {
            t.Errorf("Div(%d,%d)=(%d,%d), check %d", i, j, quo, rem, got)
        }
    }

    quick.Check(q, nil)
}
```

---

# testing/quick

```go
type small int

func (s small) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(small(rand.Intn(100)))
}

func TestQuick(t *testing.T) {
	f := func(si, sj small) bool {
                i, j := int(si), int(sj)
                quo, rem := Div(i, j+1)
                return i == quo*j+rem
	}

	if err := quick.Check(f, nil); err != nil {
		t.Log(err)
	}

}
```

---

# github.com/google/gofuzz

- split off from kubernetes

- random value generation

- register handlers for each type

- use as a component for other tests

---

# generational fuzzing

- lots of packages

- might need to write your own if your language is "tricky"

---

# github.com/zimmski/tavor

- framework with lots of features

- grammar based fuzzing

---

# github.com/zimmski/tavor

```
START = target
target = metric | function
function = name "(" arguments ")"
name = ([A-Za-z]) *([\w])
arguments = argument *( "," argument )
argument = metric | function | qstring | number
qstring = "\"" +([\w]) "\""
number = +([0-9]) | +([0-9]) "." +([0-9])
metric = name *( "." +([\w]) )
```

---

# github.com/MozillaSecurity/dharma

```
target :=
	+metric+
	+function+

function :=
	+node+(%repeat%(+argument+, ","))

argument :=
	+metric+
	+function+
	+qstring+
	+common:integer+
	+common:decimal_number+

node :=
	%repeat%(+alpha+)%repeat%(+word+)

qstring :=
	"+common:text+"

metric :=
	%repeat%(+node+, ".")
```

---

# sequitur

- dgryski/go-sequitur

---

# github.com/dvyukov/go-fuzz

- dvyukov

- mutation fuzzing

- coverage guided

- file formats, protocols, parsing *anything*

- based on afl

- understands simple text protocols

- clustering mode

---

# github.com/dvyukov/go-fuzz

- cd github.com/user/package
- mkdir corpus && cp inputs* corpus
- go-fuzz-build && go-fuzz

---

# github.com/dvyukov/go-fuzz

- `workdir/corpus`

    - `test1.in`
    - `<sha1>`

--

- `workdir/crashers`

    - `<sha1>`
    - `<sha1>.quoted`
    - `<sha1>.output`

--

- `workdir/supressions`

---

# github.com/dvyukov/go-fuzz

```go
func Fuzz(data []byte) int {
    if _, err := Decode(data); err != nil {
        return 0
    }

    return 1
}
```

---

# github.com/dvyukov/go-fuzz

```go
func Fuzz(data []byte) int {
    packed := Encode(data);

    var unpacked []byte
    var err error
    if unpacked, err = Decode(packed); err != nil {
        panic("unpacking packed data failed")
    }

    if !bytes.Equal(unpacked, data) {
        panic("roundtrip failed")
    }

    return 1
}
```

---

# github.com/dvyukov/go-fuzz

```go
func Fuzz(data []byte) int {
    fast := EncodeFast(data);
    slow := EncodeSlow(data)

    if !bytes.Equal(fast, slow) {
        panic("behaviour mismatch")
    }

    return 1
}
```

---

# github.com/dvyukov/go-fuzz

```go
func Fuzz(data []byte) int {
    purego := hashGo(data);
    asm, cgo := C.Hash(data), hashAsm(data)

    if !bytes.Equal(purego, asm) || !bytes.Equal(purego, cgo) {
        panic("behaviour mismatch")
    }

    return 1
}
```

---

# stateful fuzzing

- ideas from functional programming

- verifying invariants and post-conditions on API calls

- trying to figure out what a Go version would look like

---

# stateful fuzzing

- list of API calls

```
func (h heap) Put(key, prio int) error { ... }
func (h heap) Min() (int, error)  { ... }
func (h heap) Len() int { ... }
func (h heap) Cap() int { ... }
```

---

# stateful fuzzing

- list of API calls with claims

```
"put": {
    pre:  func() bool { return h.Len() < h.Cap() },
    call: h.Put,
    post: func() bool { return h.Len() != 0 },
}
```

---

# stateful fuzzing

- list of API calls with models

```
"put": {
    pre:  func() bool { return h.Len() < h.Cap() },
    call: func(x, prio int) { h.Put(x, prio); dumbHeap.Put(x, prio) }
    post: func() bool { h.Len() == dumbHeap.Len() },
}
```

---

# stateful fuzzing

- list of API calls that leads to a failure

```
put(1, 12); put(2, 19); getmin(); getmin(); put(4, 31);
```

- and minimize it; extract the signal from the noise

```
put(1, 12); getmin(); put(4, 31);
```

---

# stateful fuzzing

- dgryski/go-tinymap: real-world custom example

- mschoch/smat: sadly abandoned repo

---

# more reading

- randomized test case generation

- test case minimization

- https://www.fuzzingbook.org/

---
class: center, middle, inverse

## fin

???

vim: ft=markdown
