package org.example;

import org.apache.beam.sdk.coders.StringUtf8Coder;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.transforms.*;
import org.apache.beam.sdk.values.TypeDescriptor;
import org.apache.beam.sdk.values.TypeDescriptors;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class PipelineProperties {
    public static void main(String[] args) {
        PipelineOptions options = PipelineOptionsFactory.create();
        System.out.println("Runner: " + options.getRunner().getName());

        System.out.println("     |------------------|      ");
        final List<String> input = Arrays.asList(
                "Arshad,200",
                "Rohan,300",
                "Akash,400"
        );
        Pipeline pipeline = Pipeline.create(options);
        pipeline.apply(Create.of(input)).setCoder(StringUtf8Coder.of())
                .apply("PrintInput", MapElements.via(new SimpleFunction<String, String>() {
                    @Override
                    public String apply(String input) {
                        System.out.println(input);
                        return input;
                    }
                }))
                .apply("ExtractPaymentTypes", FlatMapElements
                        .into(TypeDescriptors.strings())
                        .via((String line) -> Collections.singletonList(line.split(",")[0])))
                .apply("PrintOutput", MapElements.via(new SimpleFunction<String, String>() {
                    @Override
                    public String apply(String input) {
                        System.out.println(input);
                        return input;
                    }
                }));

        pipeline.run();
    }
}
