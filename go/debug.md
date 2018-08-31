There are 2 tools to debug Go, one is gdb, and the other is dlv.

# code for test
```
-bash-4.1$ cat -n main.go 
     1  package main
     2  
     3  import (
     4    "fmt"
     5  )
     6  
     7  var ch chan int
     8  var exit chan int
     9  
    10  func main() {
    11    ch = make(chan int)
    12    exit = make(chan int)
    13    go produce()
    14    go consume()
    15    <-exit
    16  }
    17  
    18  func produce() {
    19    for i := 0; i < 100; i++ {
    20      ch <- i
    21    }
    22    close(ch)
    23  }
    24  
    25  func consume() {
    26    for i := range ch {
    27      fmt.Println(i)
    28    }
    29    close(exit)
    30  }
    31  
-bash-4.1$ go build -gcflags "-N -l"
-bash-4.1$ ls
main.go  testdebug
```

# gdb
We can use gdb to debug Go, you can find in its offical website: https://golang.org/doc/gdb
## When set GOMAXPROCS=1
In the same thread, when a Go-routine block, the thread is not blocked, it will execute another Go-routine that is not blocked. In the test case, I use 3 Go-routines (main, produce, consume), when program started, only one thread started to run 3 Go-routines, it can switch Go-routines in the same thread. For the reason, you can study assemble code: go/src/runtime/asm_amd64.s:2361.
```
-bash-4.1$ export GOMAXPROCS=1
-bash-4.1$ gdb testdebug
GNU gdb (GDB) 7.9
... Skip GDB initialize info ...
(gdb) b 15
Breakpoint 1 at 0x4832fb: file /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go, line 15.
(gdb) b 20
Breakpoint 2 at 0x483385: file /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go, line 20.
(gdb) b 27
Breakpoint 3 at 0x48346a: file /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go, line 27.
(gdb) r
Starting program: /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/testdebug 

Breakpoint 1, main.main () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:15
15    <-exit
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 767414 "testdebug" main.main () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:15
(gdb) bt
#0  main.main () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:15
(gdb) c
Continuing.
Breakpoint 2, main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
20      ch <- i
(gdb) p i
$1 = 1
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 767414 "testdebug" main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
(gdb) c
Continuing.

Breakpoint 3, main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
27      fmt.Println(i)
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 767414 "testdebug" main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
(gdb) bt
#0  main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) up
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
2361    BYTE  $0x90 // NOP
(gdb) l
2356    RET
2357  
2358  // The top-most function running on a goroutine
2359  // returns to goexit+PCQuantum.
2360  TEXT runtime��goexit(SB),NOSPLIT,$0-0
2361    BYTE  $0x90 // NOP
2362    CALL  runtime��goexit1(SB)  // does not return
2363    // traceback from goexit1 must hit code range of goexit
2364    BYTE  $0x90 // NOP
2365  
Breakpoint 2, main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
20      ch <- i
(gdb) bt
#0  main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 767414 "testdebug" main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
(gdb) c
Continuing.

Breakpoint 2, main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
20      ch <- i
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 767414 "testdebug" main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
(gdb) bt
#0  main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) c
Continuing.

Breakpoint 3, main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
27      fmt.Println(i)
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 767414 "testdebug" main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
(gdb) bt
#0  main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
```
## When set GOMAXPROCS=4
Even we set GOMAXPROCS=4, but only one thread started when program starts. The new thread is created when need to do. In the test case, I use 3 Go-routines (main, produce, consume), when program started, only one thread started. Then when produce and consume Go-routines creates and main Go-routine blocked, another thread started. Only 2 threads works in the whole process duration, because no more needed, even we allow max 4 procedure. 
```
-bash-4.1$ export GOMAXPROCS=4
-bash-4.1$ gdb testdebug 
GNU gdb (GDB) 7.9
... Skip GDB initialize info ...
(gdb) b 15
Breakpoint 1 at 0x4832fb: file /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go, line 15.
(gdb) b 20
Breakpoint 2 at 0x483385: file /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go, line 20.
(gdb) b 27
Breakpoint 3 at 0x48346a: file /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go, line 27.
(gdb) r
Starting program: /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/testdebug 

Breakpoint 1, main.main () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:15
15    <-exit
(gdb) info threads
  Id   Target Id         Frame 
* 1    LWP 781891 "testdebug" main.main () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:15
(gdb) bt
#0  main.main () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:15
(gdb) c
Continuing.
[New LWP 781898]
[Switching to LWP 781898]

Breakpoint 2, main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
20      ch <- i
(gdb) info threads
  Id   Target Id         Frame 
* 2    LWP 781898 "testdebug" main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
  1    LWP 781891 "testdebug" runtime.futex () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:527
(gdb) bt
#0  main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) c
Continuing.
[Switching to LWP 781891]

Breakpoint 3, main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
27      fmt.Println(i)
(gdb) info threads
  Id   Target Id         Frame 
  2    LWP 781898 "testdebug" main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
* 1    LWP 781891 "testdebug" main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
(gdb) bt
#0  main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) print i
$1 = 0
(gdb) c
Continuing.
[Switching to LWP 781898]

Breakpoint 2, main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
20      ch <- i
(gdb) info threads
  Id   Target Id         Frame 
* 2    LWP 781898 "testdebug" main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
  1    LWP 781891 "testdebug" 0x000000000048346f in main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
(gdb) bt
#0  main.produce () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:20
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) c
Continuing.
0
[Switching to LWP 781891]

Breakpoint 3, main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
27      fmt.Println(i)
(gdb) info threads
  Id   Target Id         Frame 
  2    LWP 781898 "testdebug" runtime.usleep () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:140
* 1    LWP 781891 "testdebug" main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
(gdb) bt
#0  main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
#1  0x000000000044f671 in runtime.goexit () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
#2  0x0000000000000000 in ?? ()
(gdb) thread 1
[Switching to thread 1 (LWP 781891)]
#0  main.consume () at /gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug/main.go:27
27      fmt.Println(i)
(gdb) thread 2
[Switching to thread 2 (LWP 781898)]
#0  runtime.usleep () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:140
140   RET
(gdb) l
135   MOVL  $0, R10
136   MOVQ  SP, R8
137   MOVL  $0, R9
138   MOVL  $SYS_pselect6, AX
139   SYSCALL
140   RET
141 
142 TEXT runtime��gettid(SB),NOSPLIT,$0-4
143   MOVL  $SYS_gettid, AX
144   SYSCALL
(gdb) bt
#0  runtime.usleep () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:140
#1  0x0000000000433a5d in runtime.runqgrab (_p_=0xc420026500, batch=0xc4200245e8, batchHead=1, stealRunNextG=true, ~r4=0) at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:4795
#2  0x0000000000433b57 in runtime.runqsteal (_p_=0xc420024000, p2=0xc420026500, stealRunNextG=true, ~r3=0x0) at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:4830
#3  0x000000000042d119 in runtime.findrunnable (gp=0xc42008a180, inheritTime=false) at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:2281
#4  0x000000000042e06b in runtime.schedule () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:2536
#5  0x000000000042e396 in runtime.park_m (gp=0xc42008a180) at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:2599
#6  0x000000000044d0bb in runtime.mcall () at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:351
#7  0x0000000000000000 in ?? ()
```

# dlv
Delve is a better debugger for the Go programming language. The goal of the project is to provide a simple, full featured debugging tool for Go. Delve should be easy to invoke and easy to use. Chances are if you're using a debugger, things aren't going your way. With that in mind, Delve should stay out of your way as much as possible. 

You may see in its github: https://github.com/derekparker/delve. 

It is a debugger tool written in Go, and many Go IDEs (such Goland) uses dlv as its default debugger, it is very good and powerful.

It has better debug features than gdb, as it is designed for Go language. 

For example, 

1. we can list go-routines and switch between go-routines in dlv but cannot in gdb.

2. we can see go source code to schedule (park, unpark) the go-routines.

Its commands are very similar than gdb, easy to learn:
```
(dlv) help
The following commands are available:
    args ------------------------ Print function arguments.
    break (alias: b) ------------ Sets a breakpoint.
    breakpoints (alias: bp) ----- Print out info for active breakpoints.
    clear ----------------------- Deletes breakpoint.
    clearall -------------------- Deletes multiple breakpoints.
    condition (alias: cond) ----- Set breakpoint condition.
    config ---------------------- Changes configuration parameters.
    continue (alias: c) --------- Run until breakpoint or program termination.
    disassemble (alias: disass) - Disassembler.
    exit (alias: quit | q) ------ Exit the debugger.
    frame ----------------------- Executes command on a different frame.
    funcs ----------------------- Print list of functions.
    goroutine ------------------- Shows or changes current goroutine
    goroutines ------------------ List program goroutines.
    help (alias: h) ------------- Prints the help message.
    list (alias: ls | l) -------- Show source code.
    locals ---------------------- Print local variables.
    next (alias: n) ------------- Step over to next source line.
    on -------------------------- Executes a command when a breakpoint is hit.
    print (alias: p) ------------ Evaluate an expression.
    regs ------------------------ Print contents of CPU registers.
    restart (alias: r) ---------- Restart process.
    set ------------------------- Changes the value of a variable.
    source ---------------------- Executes a file containing a list of delve commands
    sources --------------------- Print list of source files.
    stack (alias: bt) ----------- Print stack trace.
    step (alias: s) ------------- Single step through program.
    step-instruction (alias: si)  Single step a single cpu instruction.
    stepout --------------------- Step out of the current function.
    thread (alias: tr) ---------- Switch to the specified thread.
    threads --------------------- Print out info for every traced thread.
    trace (alias: t) ------------ Set tracepoint.
    types ----------------------- Print list of types
    vars ------------------------ Print package variables.
    whatis ---------------------- Prints type of an expression.
Type help followed by a command for full documentation.
```

Here is to debug the above code using dlv in the case when GOMAXPROCS=1.

```
-bash-4.1$ pwd
/gpfs/DEV/PWO/weliu/code/go/src/lab/testdebug
-bash-4.1$ go install
-bash-4.1$ export GOMAXPROCS=1
-bash-4.1$ dlv debug
-bash-4.1$ dlv debug
Type 'help' for list of commands.
(dlv) b main.go:15
Breakpoint 1 set at 0x49c2fb for main.main() ./main.go:15
(dlv) b main.go:20
Breakpoint 2 set at 0x49c385 for main.produce() ./main.go:20
(dlv) b main.go:27
Breakpoint 3 set at 0x49c46a for main.consume() ./main.go:27
(dlv) r
Process restarted with PID 806668
(dlv) c
> main.main() ./main.go:15 (hits goroutine(1):1 total:1) (PC: 0x49c2fb)
    10: func main() {
    11:   ch = make(chan int)
    12:   exit = make(chan int)
    13:   go produce()
    14:   go consume()
=>  15:   <-exit
    16: }
    17: 
    18: func produce() {
    19:   for i := 0; i < 100; i++ {
    20:     ch <- i
(dlv) threads
* Thread 806668 at 0x49c2fb ./main.go:15 main.main
  Thread 806698 at 0x4565a3 /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:140 runtime.usleep
  Thread 806699 at 0x456ad3 /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:527 runtime.futex
(dlv) bt
0  0x000000000049c2fb in main.main
   at ./main.go:15
1  0x000000000042c290 in runtime.main
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:198
2  0x0000000000455531 in runtime.goexit
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
(dlv) goroutines
[6 goroutines]
* Goroutine 1 - User: ./main.go:15 main.main (0x49c2fb) (thread 806668)
  Goroutine 2 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:240 runtime.forcegchelper (0x42c450)
  Goroutine 3 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
  Goroutine 4 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/mfinal.go:161 runtime.runfinq (0x416460)
  Goroutine 5 - User: ./main.go:18 main.produce (0x49c350)
  Goroutine 6 - User: ./main.go:25 main.consume (0x49c3f0)
(dlv) c
> main.produce() ./main.go:20 (hits goroutine(5):1 total:1) (PC: 0x49c385)
    15:   <-exit
    16: }
    17: 
    18: func produce() {
    19:   for i := 0; i < 100; i++ {
=>  20:     ch <- i
    21:   }
    22:   close(ch)
    23: }
    24: 
    25: func consume() {
(dlv) threads
* Thread 806668 at 0x49c385 ./main.go:20 main.produce
  Thread 806698 at 0x4565a3 /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:140 runtime.usleep
  Thread 806699 at 0x456ad3 /gpfs/DEV/PWO/weliu/tools/go/src/runtime/sys_linux_amd64.s:527 runtime.futex
(dlv) bt
0  0x000000000049c385 in main.produce
   at ./main.go:20
1  0x0000000000455531 in runtime.goexit
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
(dlv) goroutins
Command failed: command not available
(dlv) goroutines
[6 goroutines]
  Goroutine 1 - User: ./main.go:15 main.main (0x49c314)
  Goroutine 2 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
  Goroutine 3 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
  Goroutine 4 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
* Goroutine 5 - User: ./main.go:20 main.produce (0x49c385) (thread 806668)
  Goroutine 6 - User: ./main.go:26 main.consume (0x49c445)
(dlv) goroutine 1
Switched from 5 to 1 (thread 806668)
(dlv) bt
0  0x000000000042c6b9 in runtime.gopark
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292
1  0x000000000042c76e in runtime.goparkunlock
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:297
2  0x0000000000404929 in runtime.chanrecv
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/chan.go:518
3  0x00000000004046cb in runtime.chanrecv1
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/chan.go:400
4  0x000000000049c314 in main.main
   at ./main.go:15
5  0x000000000042c290 in runtime.main
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:198
6  0x0000000000455531 in runtime.goexit
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
(dlv) goroutines
[6 goroutines]
* Goroutine 1 - User: ./main.go:15 main.main (0x49c314)
  Goroutine 2 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
  Goroutine 3 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
  Goroutine 4 - User: /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 runtime.gopark (0x42c6b9)
  Goroutine 5 - User: ./main.go:20 main.produce (0x49c385) (thread 806668)
  Goroutine 6 - User: ./main.go:26 main.consume (0x49c445)
(dlv) goroutine 5
Switched from 1 to 5 (thread 806668)
(dlv) bt
0  0x000000000049c385 in main.produce
   at ./main.go:20
1  0x0000000000455531 in runtime.goexit
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
(dlv) 
0  0x000000000049c385 in main.produce
   at ./main.go:20
1  0x0000000000455531 in runtime.goexit
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
(dlv) frame 1 list
Goroutine 5 frame 1 at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361 (PC: 0x455531)
  2356:   RET
  2357: 
  2358: // The top-most function running on a goroutine
  2359: // returns to goexit+PCQuantum.
  2360: TEXT runtime��goexit(SB),NOSPLIT,$0-0
=>2361:   BYTE  $0x90 // NOP
  2362:   CALL  runtime��goexit1(SB)  // does not return
  2363:   // traceback from goexit1 must hit code range of goexit
  2364:   BYTE  $0x90 // NOP
  2365: 
  2366: // This is called from .init_array and follows the platform, not Go, ABI.
(dlv) goroutine 6 frame 0 list
Goroutine 6 frame 0 at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 (PC: 0x42c6b9)
   287:   mp.waittraceev = traceEv
   288:   mp.waittraceskip = traceskip
   289:   releasem(mp)
   290:   // can't do anything that might move the G between Ms here.
   291:   mcall(park_m)
=> 292: }
   293: 
   294: // Puts the current goroutine into a waiting state and unlocks the lock.
   295: // The goroutine can be made runnable again by calling goready(gp).
   296: func goparkunlock(lock *mutex, reason string, traceEv byte, traceskip int) {
   297:   gopark(parkunlock_c, unsafe.Pointer(lock), reason, traceEv, traceskip)
(dlv) goroutine 6 frame 1 list
Goroutine 6 frame 1 at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:297 (PC: 0x42c76e)
   292: }
   293: 
   294: // Puts the current goroutine into a waiting state and unlocks the lock.
   295: // The goroutine can be made runnable again by calling goready(gp).
   296: func goparkunlock(lock *mutex, reason string, traceEv byte, traceskip int) {
=> 297:   gopark(parkunlock_c, unsafe.Pointer(lock), reason, traceEv, traceskip)
   298: }
   299: 
   300: func goready(gp *g, traceskip int) {
   301:   systemstack(func() {
   302:     ready(gp, traceskip, true)
(dlv) goroutine 6 frame 2 list
Goroutine 6 frame 2 at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/chan.go:518 (PC: 0x404929)
   513:   mysg.g = gp
   514:   mysg.isSelect = false
   515:   mysg.c = c
   516:   gp.param = nil
   517:   c.recvq.enqueue(mysg)
=> 518:   goparkunlock(&c.lock, "chan receive", traceEvGoBlockRecv, 3)
   519: 
   520:   // someone woke us up
   521:   if mysg != gp.waiting {
   522:     throw("G waiting list is corrupted")
   523:   }
(dlv) goroutine 6 frame 3 list
Goroutine 6 frame 3 at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/chan.go:405 (PC: 0x40470b)
   400:   chanrecv(c, elem, true)
   401: }
   402: 
   403: //go:nosplit
   404: func chanrecv2(c *hchan, elem unsafe.Pointer) (received bool) {
=> 405:   _, received = chanrecv(c, elem, true)
   406:   return
   407: }
   408: 
   409: // chanrecv receives on channel c and writes the received data to ep.
   410: // ep may be nil, in which case received data is ignored.
(dlv) goroutine 6 list
Goroutine 6 frame 0 at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292 (PC: 0x42c6b9)
   287:   mp.waittraceev = traceEv
   288:   mp.waittraceskip = traceskip
   289:   releasem(mp)
   290:   // can't do anything that might move the G between Ms here.
   291:   mcall(park_m)
=> 292: }
   293: 
   294: // Puts the current goroutine into a waiting state and unlocks the lock.
   295: // The goroutine can be made runnable again by calling goready(gp).
   296: func goparkunlock(lock *mutex, reason string, traceEv byte, traceskip int) {
   297:   gopark(parkunlock_c, unsafe.Pointer(lock), reason, traceEv, traceskip)
(dlv) goroutine 6 list 30
Command failed: no code at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:30
(dlv) goroutine 6 bt
0  0x000000000042c6b9 in runtime.gopark
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:292
1  0x000000000042c76e in runtime.goparkunlock
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/proc.go:297
2  0x0000000000404929 in runtime.chanrecv
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/chan.go:518
3  0x000000000040470b in runtime.chanrecv2
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/chan.go:405
4  0x000000000049c445 in main.consume
   at ./main.go:26
5  0x0000000000455531 in runtime.goexit
   at /gpfs/DEV/PWO/weliu/tools/go/src/runtime/asm_amd64.s:2361
```
From the stack trace by the debugger dlv, we can understand how to switch go-routine in the same thread: when channel blocked, runtime.gopark() will be called, it will release M (thread) so M can switch to other go-routine.

