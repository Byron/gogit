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

