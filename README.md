# Go Practice
I set up this repository to screw around with Go and gRPC a little bit.
Got around to leveraging gRPC and Protobuf to build a _super_ simple application which provides the client with show recommendations.
I wrote some fairly detailed comments throughout just in case anybody wants to look this over, but I will briefly describe the setup below.

## Protobuf
With protobuf, we can define the structure of our incoming and outgoing messages, along with services which will provide specific pieces of functionality.
If you look at `proto/application.proto` you will see that we define our input (`ShowRecommendationRequest`), output (`ShowRecommendation`) and our service (`ShowRecommendationService`).
We will use Protobuf to generate Go code from this file so that we can instantiate these messages and services, and handle them with gRPC. The Protobuf-generated code
can be found under `src/application/application.pb.go`.

## gRPC
gRPC allows us to pass Protobuf messages back and forth, negating the use of JSON.
In our `server.go` file, you can see that we use the gRPC Go library to serve up our `ShowRecommendationService`, which will handle the input that
our client provides.

## Running the Example
When running our example, the expectation is that we can pass our Protobuf-generated messages back and forth between the client and server.
We will send the server a `ShowRecommendationRequest` with our name, and expect to receive a `ShowRecommendation` which includes our name and the name
of the show that is being recommended.

Let's run our server like so:
```
$ go run server.go 
Server running...
```

Now, let's execute our client code a few times and observe the results:
```
$ go run client.go 
2022/03/03 13:15:34 Response from server: name:"Matty" show:"Daredevil"
$ go run client.go 
2022/03/03 13:15:36 Response from server: name:"Matty" show:"Midnight Mass"
$ go run client.go 
2022/03/03 13:15:39 Response from server: name:"Matty" show:"The Haunting of Hill House"
```

If you were to change the name that we provide in our `ShowRecommendationRequest` message, you would see that change reflected in the resulting
`ShowRecommendation` as well.

## Further Considerations
A couple of things worth noting that I will list out below.
1. At the time of setting up this project, I am brand new to writing Go. I am by no means recommending anything in here as standard practice; I am just familiarizing myself with these tools.
2. The Protobuf-generated code encompasses the server _and_ client functionality. In this example our client lives in the same codebase as our server, but if we had a client that lived in a separate codebase (very likely if implementing microservices), we would need to share that generated code somehow. This could possibly be achieved by storing all Protobuf definitions and generated code in a separate repository, which is then programmatically copied into other projects with Terraform. It could also possibly be achieved by storing these results in a shared library like a Rubygem which could then be versioned and required by multiple projects. A coworker also recommended Git sub-modules as a potential solution.
3. This example is only intended to run on a local machine, so of course there is no proper routing and SSL is not supported. As a result, the setup of our gRPC server is exceedingly simple. It is safe to assume that doing something like this in a production environment would require some extra work and configuration.
