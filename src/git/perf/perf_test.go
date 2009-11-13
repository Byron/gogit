package misc 

import ( "testing"; "time"; "fmt";"malloc"; )

type Clock struct {
	startTime int64;
	startMem uint64;
}

func (c *Clock) start() { c.startTime = time.Nanoseconds(); c.startMem = malloc.GetStats().Alloc; }
func (c *Clock) elapsed() float64 { return ( float64(time.Nanoseconds() - c.startTime ) / 1000000000)  }
func (c *Clock) elapsedMsg(msg string, iterations uint) float64 {
	e := c.elapsed();
	fmt.Printf( "Did %d %s in %v [s] ( %v / s )\n", iterations, msg, e, float64(iterations)/e );
	return e;
}

func (c *Clock) memUsageMsg(msg string) uint64 {
	usedMemoryKb := ( malloc.GetStats().Alloc - c.startMem ) / 1024;
	fmt.Printf( "%s used %d [kb]\n", msg, usedMemoryKb );
	return usedMemoryKb;
}


func TestLoops(t *testing.T) {
	var timer Clock;
	
	// LOOPING SPEED
	nit := uint(100000000);
	timer.start();
	for i:=uint(0); i < nit; i++ {}
	timer.elapsedMsg( "nop iterations", nit );
	timer.memUsageMsg("nop iterations" );
	
	// ALLOCATION SPEED
	nit = 250000;
	timer.start();
	for i:=uint(0); i < nit; i++ {
		var b *byte = new(byte);
		b = b;
	}
	timer.elapsedMsg( "byte allocations", nit );
	// This shows it collects them pretty fast
	timer.memUsageMsg("Byte Allocations" );
	
	// MAP PERFORMANCE
}
	
func TestVectorPerformance(t *testing.T){
	
}
	
func TestMapPerformance(t *testing.T) {
	t.FailNow();
}
