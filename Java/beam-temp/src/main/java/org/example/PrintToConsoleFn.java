package org.example;

import org.apache.beam.sdk.transforms.DoFn;

public class PrintToConsoleFn extends DoFn<String, String> {
    @ProcessElement
    public void processElement(ProcessContext c) {
        System.out.println(c.element());
        c.output(c.element());
    }
}
