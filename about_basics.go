package go_koans

func aboutBasics() {
	assert(__bool__ == false)  // what is truth?
	assert(__bool__ != true) // in it there is nothing false

	var i int = __int__
	assert(i == -1) // precision is in the eye of the beholder

	k := __int__ //short assignment can be used, as well
	assert(k == -1)

	assert(5%2 == 1)
	assert(5*2 == 10)
	assert(5^2 == 7)

	var x int = __int__
	assert(x == __int__) // zero values are valued in Go

	var f float32 = __float32__
	assert(f == __float32__) // for types of all types

	var s string = __string__
	assert(s == __string__) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == 0)     // and types within composite types
	assert(c.f == 0) // which match the other types
	assert(c.s == "")  // in a typical way
}
