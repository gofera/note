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
                log("after INCREASER sending: " + curr);
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
                    log("after ECHO sending: " + curr);
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