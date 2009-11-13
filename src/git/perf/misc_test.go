package misc

import "testing"
import "fmt"

/* ASSORTED TESTS */

func byteManip(a [10]byte) {
	a[0] = 2;
}

func byteManipPointer(a *[10]byte) {
	a[0] = 3;
}

func TestArrayPassing(t *testing.T) {
	var b [10]byte;
	byteManip(b);
	if b[0] != 0{
		t.Error( "bytes passed by reference")
	}
	
	byteManipPointer(&b);
	if b[0] != 3 {
		t.Error( "pointer bytes passed by copy" );
	}
	
}


// type derived from built-in one
type tenbyte [10]byte

// discover whether instance methods copy their instance when called
func (b tenbyte) makeChange() {
	b[0] = 4;
}

func (bp *tenbyte) makeChangePointer() {
	bp[0] = 4;
}

// type derived from struct
type simple struct {
	v int8;
}

func (s simple) makeChange() {
	s.v = 1;
}

func (s *simple) makeChangePointer() {
	s.v = 1;
}


func TestFunctionInvocation(t *testing.T) {
	// test whether value methods on value types are true copies
	var tb tenbyte;
	tb.makeChange();
	if tb[0] != 0 {
		t.Error( "value method on stack instance should have copied itself into method" );
	}
	
	// what happens if pointers are used with value methods ? 
	// be explict here
	var tbp *tenbyte= new(tenbyte);	
	tbp.makeChange();
	if tbp[0] != 0 {
		t.Error( "value method on heap instance should have copied itself into method" );
	}
	
	// call pointer func on non-pointer
	tb.makeChangePointer();
	if tb[0] != 4 {
		t.Error( "Change did not propagate to stack instance" );
	}
	
	tbp.makeChangePointer();
	if tbp[0] != 4 {
		t.Error( "Change did not propagate to heap instance" );
	}
	
	// same for struct type
	var s simple;
	s.makeChange();
	if s.v != 0 {
		t.Error( "value method on stack instance should have copied itself into method" );
	}
	
	var sp *simple= new(simple);	
	sp.makeChange();
	if sp.v != 0 {
		t.Error( "value method on heap instance should have copied itself into method" );
	}
	
	// call pointer func on non-pointer
	s.makeChangePointer();
	if s.v != 1 {
		t.Error( "Change did not propagate to stack instance" );
	}
	
	sp.makeChangePointer();
	if sp.v != 1 {
		t.Error( "Change did not propagate to heap instance" );
	}
}


func every( out chan<- uint, number uint ) {
	c := uint(2);
	for {
		if c % number == 0 {
			out <- c;
		}
		c += 1;
	}
}

func TestChannelThreading(t *testing.T) {
	result := make( chan uint, 10 );
	
	// should take a while - would expect at least 2 threads here
	go every( result, 100000000 );
	go every( result, 50000001 );
	
	// get some values
	for i := 0; i < 20; i++ {
		fmt.Println( "Got result: ", <-result );
		
	}
}
