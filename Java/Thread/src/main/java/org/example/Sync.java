package org.example;

import java.util.concurrent.ThreadPoolExecutor;

public class Sync {
    public static void main(String[] args) {


        syncMethod();
    }

    static void syncMethod() {
        System.out.println("Inside Sync Method" + Thread.currentThread().getId());
    }
}


