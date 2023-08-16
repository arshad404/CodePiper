package org.example;

import org.apache.beam.sdk.transforms.DoFn;

import java.util.Objects;

public class ExtractPaymentTypeFn extends DoFn<String, String> {
    @ProcessElement
    public void processElement(ProcessContext c) {
        String[] tokens = Objects.requireNonNull(c.element()).split(",");
        if(tokens.length > 0) {
            c.output(tokens[0]);
        }
    }
}
