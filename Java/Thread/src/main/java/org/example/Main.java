package org.example;

import org.w3c.dom.ls.LSOutput;

import java.util.concurrent.Callable;

public class Main {
    public static void main(String[] args) throws Exception {
        System.out.println("Threading Start " + Thread.currentThread().getId());

        // function that we are passing inside thread is runnable
        Thread t = new Thread(() -> System.out.println("New Thread " + Thread.currentThread().getId()));
        t.start();

        Thread.sleep(5000);

        Thread rt = new Thread(new RunThread());
        rt.start();

        CustomThread ct = new CustomThread();
        ct.start();

        CallThread clt = new CallThread();
        System.out.println(clt.call());

    }
}

class CustomThread extends Thread {
    public void run() {
        System.out.println("I am a custom thread " + Thread.currentThread().getId());
    }
}

class RunThread implements Runnable {

    @Override
    public void run() {
        System.out.println("I am inside runnable " + Thread.currentThread().getId());
    }
}

class CallThread implements Callable<String> {

    @Override
    public String call() throws Exception {
        return "I am inside Callable";
    }
}