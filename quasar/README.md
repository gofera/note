# 建立环境
根据[官方文档](http://docs.paralleluniverse.co/quasar/)，下载`quasar-mvn-archetype`脚手架创建工程，单元测试、运行。
```
git clone https://github.com/puniverse/quasar-mvn-archetype
cd quasar-mvn-archetype
mvn install
cd ..
mvn archetype:generate -DarchetypeGroupId=co.paralleluniverse -DarchetypeArtifactId=quasar-mvn-archetype -DarchetypeVersion=0.7.4 -DgroupId=testgrp -DartifactId=testprj
cd testprj
mvn test
mvn clean compile dependency:properties exec:exec
```
在`pom.xml`文件把`quasar`版本改为最新的`0.7.9`，测试，OK。

脚手架生成一个例子程序`QuasarIncreasingEchoApp`，可以用maven直接运行，但是在IDE执行需要加一个VM参数：
```
-javaagent:C:\Users\wenzhe\.m2\repository\co\paralleluniverse\quasar-core\0.7.9\quasar-core-0.7.9.jar
```

我改下QuasarIncreasingEchoApp类，加上`log`，打印线程信息，例子程序：
```
package testgrp;

import java.util.concurrent.ExecutionException;

import co.paralleluniverse.strands.SuspendableCallable;
import co.paralleluniverse.strands.SuspendableRunnable;
import co.paralleluniverse.strands.channels.Channels;
import co.paralleluniverse.strands.channels.IntChannel;

import co.paralleluniverse.fibers.Fiber;
import org.slf4j.LoggerFactory;

/**
 * Increasing-Echo Quasar Example
 *
 * @author circlespainter
 */
public class QuasarIncreasingEchoApp {
    static public Integer doAll() throws ExecutionException, InterruptedException {
        final IntChannel increasingToEcho = Channels.newIntChannel(0); // Synchronizing channel (buffer = 0)
        final IntChannel echoToIncreasing = Channels.newIntChannel(0); // Synchronizing channel (buffer = 0)

        log("Before Fiber INCREASER start.");

        Fiber<Integer> increasing = new Fiber<>("INCREASER", (SuspendableCallable<Integer>) () -> {
            ////// The following is enough to test instrumentation of synchronizing methods
            // synchronized(new Object()) {}
            log("Fiber INCREASER code is called.");
            int curr = 0;
            for (int i = 0; i < 10 ; i++) {
                Fiber.sleep(10);
                log("INCREASER sending: " + curr);
                increasingToEcho.send(curr);
                curr = echoToIncreasing.receive();
                log("INCREASER received: " + curr);
                curr++;
                log("INCREASER now: " + curr);
            }
            log("INCREASER closing channel and exiting");
            increasingToEcho.close();
            return curr;
        }).start();

        log("Before Fiber ECHO start.");

        Fiber<Void> echo = new Fiber<Void>("ECHO", (SuspendableRunnable) () -> {
            log("Fiber ECHO code is called.");
            Integer curr;
            while (true) {
                Fiber.sleep(1000);
                curr = increasingToEcho.receive();
                log("ECHO received: " + curr);

                if (curr != null) {
                    log("ECHO sending: " + curr);
                    echoToIncreasing.send(curr);
                } else {
                    log("ECHO detected closed channel, closing and exiting");
                    echoToIncreasing.close();
                    return;
                }
            }
        }).start();

        log("After Fiber ECHO start.");

        try {
            increasing.join();
            echo.join();
            log("After Fiber join.");
        } catch (ExecutionException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        return increasing.get();
    }

    static public void main(String[] args) throws ExecutionException, InterruptedException {
        log("get final increasing result: " + doAll());
    }

    private static void log(String msg) {
        //System.out.println(msg);
        LoggerFactory.getLogger("QuasarApp").info(msg);
    }
}
```
运行结果：
```
10:32:20.941 [main] INFO QuasarApp - Before Fiber INCREASER start.
10:32:21.241 [main] INFO QuasarApp - Before Fiber ECHO start.
10:32:21.247 [main] INFO QuasarApp - After Fiber ECHO start.
10:32:21.254 [ForkJoinPool-default-fiber-pool-worker-1] INFO QuasarApp - Fiber INCREASER code is called.
10:32:21.254 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - Fiber ECHO code is called.
10:32:21.860 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 0
10:32:22.260 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 0
10:32:22.260 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 0
10:32:22.261 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 0
10:32:22.261 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 1
10:32:22.312 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 1
10:32:23.262 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 1
10:32:23.262 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 1
10:32:23.263 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 1
10:32:23.263 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 2
10:32:23.310 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 2
10:32:24.285 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 2
10:32:24.285 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 2
10:32:24.285 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 2
10:32:24.285 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 3
10:32:24.301 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 3
10:32:25.312 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 3
10:32:25.312 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 3
10:32:25.312 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 3
10:32:25.312 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 4
10:32:25.328 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 4
10:32:26.315 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 4
10:32:26.315 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 4
10:32:26.315 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 4
10:32:26.315 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 5
10:32:26.330 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 5
10:32:27.317 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 5
10:32:27.317 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 5
10:32:27.317 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 5
10:32:27.317 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 6
10:32:27.333 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 6
10:32:28.327 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 6
10:32:28.327 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 6
10:32:28.327 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 6
10:32:28.327 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 7
10:32:28.343 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 7
10:32:29.343 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 7
10:32:29.343 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 7
10:32:29.343 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 7
10:32:29.343 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 8
10:32:29.344 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 8
10:32:30.334 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 8
10:32:30.334 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 8
10:32:30.334 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 8
10:32:30.334 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 9
10:32:30.336 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 9
10:32:31.345 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 9
10:32:31.345 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 9
10:32:31.345 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 9
10:32:31.345 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 10
10:32:31.345 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER closing channel and exiting
10:32:32.344 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: null
10:32:32.344 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO detected closed channel, closing and exiting
10:32:32.344 [main] INFO QuasarApp - After Fiber join.
10:32:32.360 [main] INFO QuasarApp - get final increasing result: 10
```
每次使用的线程都不一样，再次运行可能是这样的：
```
10:57:23.619 [main] INFO QuasarApp - Before Fiber INCREASER start.
10:57:23.847 [main] INFO QuasarApp - Before Fiber ECHO start.
10:57:23.847 [main] INFO QuasarApp - After Fiber ECHO start.
10:57:23.847 [ForkJoinPool-default-fiber-pool-worker-1] INFO QuasarApp - Fiber INCREASER code is called.
10:57:23.862 [ForkJoinPool-default-fiber-pool-worker-1] INFO QuasarApp - Fiber ECHO code is called.
10:57:23.878 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 0
10:57:24.864 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 0
10:57:24.864 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 0
10:57:24.864 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 0
10:57:24.864 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 1
10:57:24.879 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 1
10:57:25.875 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 1
10:57:25.875 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 1
10:57:25.875 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 1
10:57:25.875 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 2
10:57:25.891 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 2
10:57:26.890 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 2
10:57:26.890 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 2
10:57:26.890 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 2
10:57:26.890 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 3
10:57:26.906 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 3
10:57:27.895 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 3
10:57:27.895 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 3
10:57:27.895 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 3
10:57:27.895 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 4
10:57:27.896 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 4
10:57:28.908 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 4
10:57:28.908 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 4
10:57:28.908 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 4
10:57:28.908 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 5
10:57:28.909 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 5
10:57:29.914 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 5
10:57:29.914 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 5
10:57:29.914 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 5
10:57:29.914 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 6
10:57:29.930 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 6
10:57:30.925 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 6
10:57:30.925 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 6
10:57:30.925 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 6
10:57:30.925 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 7
10:57:30.941 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER sending: 7
10:57:31.938 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: 7
10:57:31.938 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO sending: 7
10:57:31.938 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 7
10:57:31.938 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 8
10:57:31.953 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 8
10:57:32.939 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 8
10:57:32.939 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 8
10:57:32.939 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER received: 8
10:57:32.939 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER now: 9
10:57:32.955 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - INCREASER sending: 9
10:57:33.944 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO received: 9
10:57:33.944 [ForkJoinPool-default-fiber-pool-worker-3] INFO QuasarApp - ECHO sending: 9
10:57:33.944 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER received: 9
10:57:33.944 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER now: 10
10:57:33.944 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - INCREASER closing channel and exiting
10:57:34.953 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO received: null
10:57:34.953 [ForkJoinPool-default-fiber-pool-worker-2] INFO QuasarApp - ECHO detected closed channel, closing and exiting
10:57:34.953 [main] INFO QuasarApp - After Fiber join.
10:57:34.953 [main] INFO QuasarApp - get final increasing result: 10
```

# 协程是个啥？
```
class Fiber<V> extends Strand implements Joinable<V>, Serializable, Future<V> {
	...
	FiberScheduler scheduler;      // 协程调度机制，缺省为支持协程的ForkJoinPool调度器
	FiberTask<V> task;             // 供线程池调度的并发任务
	Stack stack;                   // 用于协程间切换时保存运行现场以供恢复
	State state;                   // 协程状态：NEW，STARTED，RUNNING，WAITING，TIMED_WAITING，TERMINATED
	SuspendableCallable<V> target; // 协程执行的代码
	Object result;                 // 协程执行的代码的返回结果（线程没有这个）
	Object fiberLocals;            // 相当于ThreadLocal，只不过更细致，只是协程本地所有
	Object inheritableFiberLocals; // 继承自父协程或线程（即启动该协程所在的那个协程或线程）
	byte priority;                 // 优先级，缺省5，范围：[1-10]
	...
    public Fiber(String name, FiberScheduler scheduler, int stackSize, SuspendableCallable<V> target) {
    	...
        this.state = State.NEW;
        this.scheduler = scheduler;
        this.target = target;
        this.task = scheduler != null ? scheduler.newFiberTask(this) : new FiberForkJoinTask(this);
        this.priority = (byte)NORM_PRIORITY;
        ...
    }
```
使用协程与线程非常相似，Strand就是协程与线程的抽象，代表一个协程或一个线程。

构造协程时，为它创建一个Task供线程池调度（缺省为ForkJoinPool线程池的支持协程的Task）。

启动协程跟启动线程一样，都是调用start方法。在协程的start方法中向线程池提交任务。

```
    public final Fiber<V> start() {
        casState(State.NEW, State.STARTED); // CAS set Fiber State to STARTED
        ...
        task.submit();
        return this;
    }
```

# Fiber（协程）执行如何切换
线程阻塞时，整个线程挂起，无法执行其它任务；但协程阻塞时，线程不会挂起，即使只有一个线程，也可以去执行其它任务。

从一个协程执行体到另一个协程执行体，再切换回来，这个过程可能线程变化了，也可能还是同一个线程，这是不是很神奇！

Fiber执行体是一个线程池的task，由线程池调度（默认是ForkJoinPool），每个throws SuspendExecution的方法，编译器或运行时都会织入字节码(优化器可以通过分析决定要不要织入)，写入pop、push、sp、switch case，throw SuspendExecution。

下面m开头的方法为普通方法，s开头为可能挂起协程的方法：

```
s1() throws SuspendExecution {
  m1();
  s2();
  m2();
  s3();
  m3();
}
s2() throws SuspendExecution {
  m4();
  s4();
  m5();
}
```
协程代码被织入后是一个状态机。
```
s1() throws SuspendExecution {
  int pc = get pc from stack
  switch (pc) {
  case 0:
    m1();

    push data, pc = 1  // 压入数据，设置pc值，为了以后可以恢复
    s2();  // 当发生挂起时抛出SuspendExecution异常，跳出整个函数，
  case 1
    pop    // 重新执行时，由于pc为1，跳到这个case，从栈中恢复数据

    m2();

    push data, pc = 2
    s3();
  case 2:
    pop
    m3();
  }
}
s2() throws SuspendExecution {
  int pc = get pc from stack
  switch (pc) {
  case 0:
    m4();

    push data, pc = 1
    s4();
  case 1:
    pop

    m5();
}
```

当需要让协程挂起的时候，调用park方法，记住运行点信息（栈指针sp、push现场状态，包括操作数、局部变量），通过抛出SuspendExecution跳出执行体，协程task顶层catch这个异常，恢复线程状态。
```
class Fiber ...
    boolean exec() {
        ...
        installFiberDataInThread(currentThread);
        ...
        try {
            ...
            final V res = run1(); // we jump into the continuation
            ...
            setResult(res);
            return true;
            ...
        } catch (SuspendExecution ex) {
            ...
            stack.resumeStack();
            runningThread = null;
            ...
            clearRunSettings();
            ...
            return false;
        }
        ...
    }
```
```
boolean park(Object blocker, boolean exclusive) {
    ... update fiber state ...
    throw new SuspendExecution
}
```
当需要让一个协程停止挂起继续执行时，恢复运行点信息（sp、pop现场状态），调用该协程的unpack方法，fork该协程任务（加入到ForkJoinPool等待线程池调度）。
```
   public boolean unpark(ForkJoinPool fjPool, Object unblocker) {
        ... Update Filber State ...
        if (newState == RUNNABLE) {
        	...
            submit();  // call ForkJoinPool's fork method
        }
        ...
    }
```

# Fiber sleep原理
调用park方法，开启timeout scheduler，让线程池在timeout时间后重新调用该协程的执行体，由sp和pop恢复现场，跳到原来挂起的地方后继续执行。
```
  scheduler.schedule(this, blocker, timeout, unit);  // "this" is the fiber itself
  park(blocker);
```

# 协程间的通信：Channel
通道，是一个基于协程的阻塞队列（当溢出模式设置为阻塞时），发送是将数据写入通道，然后unpack等待读通道数据而挂起的协程（unpack具体就是在unpack中通过线程池调度fork那个协程对应的任务，上面已经说过）。如果写缓冲满而无法写入数据，则挂起该协程，等别人从通道读取数据时unpack它，或者等超时调度任务去重新调度它自己。

```
    public boolean send(int message, long timeout, TimeUnit unit) throws SuspendExecution, InterruptedException {
        if (isSendClosed())
            return true;
        if (!queue().enq(message))
            return super.send(message, timeout, unit);
        signalReceivers();  // 将unpack等待通道而挂起的协程
        return true;
    }
```
类似的，接收过程是将数据从通道读出，并且通知那些等待写通道数据而挂起的协程。如果写缓冲满而无法写入数据，则挂起该协程，等别人从通道写入数据时unpack它，或者等超时调度任务去重新调度它自己。
```
   public Message receive() throws SuspendExecution, InterruptedException {
        if (receiveClosed)
            return closeValue();

        Message m;
        boolean closed;
        final Object token = sync.register();
        try {
            for (int i = 0;; i++) {
                closed = isSendClosed(); // must be read BEFORE queue.poll()
                if ((m = queue.poll()) != null)
                    break;

                // i can be > 0 if task state is LEASED
                if (closed) {
                    setReceiveClosed();
                    return closeValue();
                }

                sync.await(i);
            }
        } finally {
            sync.unregister(token);
        }

        assert m != null;
        signalSenders();
        return m;
    }
```

```
    public enum OverflowPolicy {
        /**
         * The sender will get an exception (except if the channel is an actor's mailbox)
         */
        THROW,
        /**
         * The message will be silently dropped.
         */
        DROP,
        /**
         * The sender will block until there's a vacancy in the channel.
         */
        BLOCK,
        /**
         * The sender will block for some time, and retry.
         */
        BACKOFF,
        /**
         * The oldest message in the queue will be removed to make room for the new message.
         */
        DISPLACE
    }
```

# “协程安全”，“协程锁”

```
public class SimpleConditionSynchronizer extends ConditionSynchronizer implements Condition {
    private final Queue<Strand> waiters = new ConcurrentLinkedQueue<>();

    public SimpleConditionSynchronizer(Object owner) {
        super(owner);
    }

    @Override
    public Object register() {
        final Strand currentStrand = Strand.currentStrand();
        record("register", "%s register %s", this, currentStrand);
        waiters.add(currentStrand);
        return null;
    }

    @Override
    public void unregister(Object registrationToken) {
        final Strand currentStrand = Strand.currentStrand();
        record("unregister", "%s unregister %s", this, currentStrand);
        if (!waiters.remove(currentStrand))
            throw new IllegalMonitorStateException();
    }

    @Override
    public void signalAll() {
        for (Strand s : waiters) {
            record("signalAll", "%s signalling %s", this, s);
            Strand.unpark(s, owner);
        }
    }

    @Override
    public void signal() {
        /*
         * We must wake up the first waiter that is actually parked. Otherwise, by the time the awakened waiter calls
         * unregister(), another one may block, and we may need to wake that one.
         */
        for (final Strand s : waiters) {
            if (s.isFiber()) {
                if (FiberControl.unpark((Fiber) s, owner)) {
                    record("signal", "%s signalled %s", this, s);
                    return;
                }
            } else {
                // TODO: We can't tell (atomically) if a thread is actually parked, so we'll wake them all up.
                // We may consider a more complex solution, a-la AbstractQueuedSynchronizer for threads
                // (i.e. with a wrapper node, containing the state)
                record("signal", "%s signalling %s", this, s);
                Strand.unpark(s, owner);
            }
        }
    }
}
```

建议：不要用共享数据加锁来通信，而是用通信来共享数据。

如果在协程中使用锁（如synchrozed），将可能会阻塞线程，也阻塞了当前线程运行的协程。

Quasar会对java.util.concurrent包中的Thread.park，还有nio包的IO等待，替换成Fiber.park，这样park to thread就变成park to fiber，所以使用java.util.concurrent包和nio包的代码，可以不用修改的跑在Fiber上。

协程安全，就是线程安全，因为同一个线程不可能同时运行两个协程。

# Fiber Join原理

# Fiber如何指定线程池（非默认的ForkJoinPool）

# 协程的优先级调度原理


# Fiber如何运行UI代码
```
new Fiber(() -> 
  ...
  UI.asyncExec(() -> ...)  // 调用UI代码
  ...
).start();
```


如果是UI.syncExec, sync过程如何不挂起线程而只挂起协程。
```
new Fiber(() -> 
  ...
  Channel<Result> ch = new Channel(0);
  UI.asyncExec(() -> {
    ... 调用UI代码，结果为result
    ch.send(result);  // 通知
  });
  Result res = ch.receive(); // 等待UI线程拿到结果
  ...
).start();
```


# 优点
1. 轻量级，可以有成千上万个协程，适合非常高的并发情况（线程做不到啊）
2. 适合大量IO密集型的任务，这种情况比用线程池效率高很多
3. 易于编写并发代码：可以像写串行代码的风格去编写并发代码，更符合人类的思考方式。

# 缺点
1. 不是所有方法都支持协程，需要符合协程的条件。
2. switch case push pop需要开销，不适合计算密集型的任务

# Reference

[Coroutine in Java - Quasar Fiber实现](https://segmentfault.com/a/1190000006079389)

[次时代Java编程(一) Java里的协程](https://segmentfault.com/a/1190000005342905)

[Java中的纤程库 - Quasar](http://colobu.com/2016/07/14/Java-Fiber-Quasar/)

[Quasar和Akka比较](https://kaimingwan.com/post/java/quasarhe-akkabi-jiao)

[Parallel Universe Blog](http://blog.paralleluniverse.co/)

