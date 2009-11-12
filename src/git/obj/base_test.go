package obj

import "testing" 
import "fmt"

func TestSha1(t *testing.T)  {
	var s Sha1;
	_, err := FromHex("00");
	if err == nil { 
		t.Error("created Sha1 from hex being too short"); 
	}
	
	// test printing 
	fmt.Sprint(err);
	sp, _ := FromHex(s.String()); 
	if sp == nil || !sp.Equals(&s) {
		switch sp {
		case nil: t.Error( "Did not get any sha from ", s.String());
		default: t.Errorf( "From hex did not produce an equal sha: %v != %v", sp, s  ) 
		}
	}
	
	// make copy 
	s2 := *sp;
	s2.Equals(sp);
	// t.Fail();
}

func TestBasics(t *testing.T)  {
	t.FailNow();
}






/* ASSORTED TESTS */

func byteManip(a [10]byte) {
	a[0] = 2;
}

func byteManipPointer(a *[10]byte) {
	a[0] = 3;
}


type tenbyte [10]byte

// discover whether instance methods copy their instance when called
func (b tenbyte) makeChange() {
	b[0] = 4;
}

func (bp *tenbyte) makeChangePointer() {
	bp[0] = 4;
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
	
	// test whether methods on value types are true copies
	var tb tenbyte;
	tb.makeChange();
	if tb[0] == 0 {
		// t.Error( "Ups, it copies the value" );
	}
	
	// what happens to pointers ? 
	// be explict here
	var tbp *tenbyte= new(tenbyte);	
	tbp.makeChange();
	if tbp[0] == 0 {
		// t.Error( "Ups, even pointers are copied, fair enough though" );
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
}
