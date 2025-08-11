// memory_phases.go
//
// Demonstrates Go memory concepts:
//   - stack vs heap (escape analysis)
//   - closures & captured state
//   - slice backing array behavior
//   - maps & channels allocation
//   - goroutine loop capture gotcha and fix
//   - sync.Pool reuse
//   - reading runtime mem stats and forcing GC
//
// Build hints:
//
//	go run memory_phases.go
//	go build -gcflags="-m" memory_phases.go   # shows what escapes to heap
//
// Try different GOGC values to observe GC behavior:
//
//	GOGC=50 go run memory_phases.go
//	GOGC=200 go run memory_phases.go
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== memory_phases demo ===")
	fmt.Println()

	stackVsHeapDemo()
	fmt.Println()

	closureDemo()
	fmt.Println()

	sliceMapChannelDemo()
	fmt.Println()

	goroutineCaptureDemo()
	fmt.Println()

	syncPoolDemo()
	fmt.Println()

	gcAndMemStatsDemo()
	fmt.Println()

	fmt.Println("Done.")
}

// ---------------------------------------------------------
// 1) Stack vs Heap / Escape analysis
// ---------------------------------------------------------

// heapEscape returns a pointer to a local variable -> must escape to heap
func heapEscape() *int {
	x := 42
	return &x
}

// noEscape returns an int value that can stay on the stack
func noEscape() int {
	y := 100
	return y
}

func stackVsHeapDemo() {
	fmt.Println("1) Stack vs Heap / Escape analysis")

	// This likely causes x to escape to the heap.
	p := heapEscape()
	fmt.Println("heapEscape returned pointer:", p, "value:", *p)

	// This can be kept on stack by compiler.
	v := noEscape()
	fmt.Println("noEscape returned value:", v)

	// If you build with: go build -gcflags="-m" memory_phases.go
	// you'll see messages indicating which variables were moved to heap.
	fmt.Println("Tip: rebuild with -gcflags=\"-m\" to see escape analysis output.")
}

// ---------------------------------------------------------
// 2) Closures and captured state
// ---------------------------------------------------------

func makeCounter(step int) func() int {
	count := 0
	// returned function is a closure capturing `count` and `step`
	return func() int {
		count += step
		return count
	}
}

func closureDemo() {
	fmt.Println("2) Closures & captured state")

	c1 := makeCounter(1)
	c2 := makeCounter(10)

	fmt.Println("c1:", c1()) // 1
	fmt.Println("c1:", c1()) // 2
	fmt.Println("c2:", c2()) // 10
	fmt.Println("c2:", c2()) // 20

	fmt.Println("Note: captured variables may live on heap due to closures.")
}

// ---------------------------------------------------------
// 3) Slices, maps, channels memory shape
// ---------------------------------------------------------

func sliceMapChannelDemo() {
	fmt.Println("3) Slices, Maps, Channels")

	// Slice shape: pointer -> len -> cap
	s := make([]int, 0, 2)
	fmt.Printf("slice initial len=%d cap=%d ptr=%p\n", len(s), cap(s), s)

	// Append will allocate new backing array when capacity exceeded.
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("after append %d: len=%d cap=%d ptr=%p\n", i, len(s), cap(s), s)
	}

	// Map: hash table allocated on heap
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("map contents: %+v (maps are heap allocated internally)\n", m)

	// Channel: always a heap object with internal buffer if buffered
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Printf("channel buffered len=%d cap=%d\n", len(ch), cap(ch))
	<-ch
	<-ch
	close(ch)
}

// ---------------------------------------------------------
// 4) Goroutine loop capture gotcha (and the fix)
// ---------------------------------------------------------

func goroutineCaptureDemo() {
	fmt.Println("4) Goroutine capture gotcha & fix")

	fmt.Println("BAD: capturing loop variable directly (may print same value):")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		// BAD pattern: capturing i from outer scope
		go func() {
			defer wg.Done()
			// Sleep a bit so chances are the loop advanced before goroutine runs
			time.Sleep(10 * time.Millisecond)
			fmt.Println("bad goroutine sees i =", i)
		}()
	}
	wg.Wait()

	fmt.Println("GOOD: pass loop variable as parameter to literal:")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			fmt.Println("good goroutine sees ii =", ii)
		}(i) // pass i explicitly
	}
	wg.Wait()
}

// ---------------------------------------------------------
// 5) sync.Pool reuse example
// ---------------------------------------------------------

func syncPoolDemo() {
	fmt.Println("5) sync.Pool reuse")

	var p = sync.Pool{
		New: func() interface{} {
			// allocate a moderately-sized buffer
			return make([]byte, 1024)
		},
	}

	// Get -> use -> Put
	buf := p.Get().([]byte)
	fmt.Println("got buffer len:", len(buf))
	// Use the buffer...
	buf[0] = 1
	// Return to pool for reuse
	p.Put(buf)

	// Later:
	buf2 := p.Get().([]byte)
	fmt.Println("got buffer again (reused ideally) len:", len(buf2))
	_ = buf2
}

// ---------------------------------------------------------
// 6) GC and runtime.MemStats demo
// ---------------------------------------------------------

func allocMany(n int) [][]byte {
	out := make([][]byte, 0, n)
	for i := 0; i < n; i++ {
		// Allocate ~1KB slices; these will be heap objects
		b := make([]byte, 1024)
		out = append(out, b)
	}
	return out
}

func printMem(prefix string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc = %d KB, Sys = %d KB, NumGC = %d\n",
		prefix, m.Alloc/1024, m.Sys/1024, m.NumGC)
}

func gcAndMemStatsDemo() {
	fmt.Println("6) GC and MemStats demo")

	// Ensure GC is enabled and debug GC percent is not limiting tracing
	debug.SetGCPercent(100)

	printMem("before allocations")

	// create some heap pressure
	huge := allocMany(2000) // ~2000 KB ~ 2 MB
	printMem("after allocations")

	// Drop references; now objects are eligible for GC
	huge = nil
	printMem("after drop references (before GC)")

	// Force GC to run
	runtime.GC()
	time.Sleep(100 * time.Millisecond) // give runtime a moment
	printMem("after runtime.GC()")

	// Show that increasing or lowering GOGC affects frequency/heap size;
	// try running with env GOGC=50 or GOGC=200 to experiment.
}
