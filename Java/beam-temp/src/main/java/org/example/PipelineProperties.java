package org.example;

import org.apache.beam.sdk.coders.StringUtf8Coder;
import org.apache.beam.sdk.options.PipelineOptions;
import org.apache.beam.sdk.options.PipelineOptionsFactory;
import org.apache.beam.sdk.Pipeline;
import org.apache.beam.sdk.transforms.Create;
import org.apache.beam.sdk.transforms.DoFn;
import org.apache.beam.sdk.transforms.ParDo;

import java.util.Arrays;
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
                .apply(ParDo.of(new PrintToConsoleFn()))
                .apply(ParDo.of(new ExtractPaymentTypeFn()))
                .apply(ParDo.of(new PrintToConsoleFn()));

        pipeline.run();
    }
}
