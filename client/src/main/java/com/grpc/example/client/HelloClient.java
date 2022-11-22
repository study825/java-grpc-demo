package com.grpc.example.client;

import com.grpc.example.HelloRequest;
import com.grpc.example.HelloResponse;
import com.grpc.example.HelloServiceGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class HelloClient {
    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 7080)
                .usePlaintext()
                .build();
        HelloServiceGrpc.HelloServiceBlockingStub stub = HelloServiceGrpc.newBlockingStub(channel);

        HelloResponse helloResponse = stub.hello(HelloRequest.newBuilder().setFirstName("qa").setLastName("q").build());

        channel.shutdown();
    }
}
