package main

func main() {

	var old struct {
		len int
		cap int
	}

	var cap int

	// START OMIT
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			for newcap < cap {
				newcap += newcap / 4
			}
		}
	}

	// END OMIT
}
