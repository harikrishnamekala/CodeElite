#CODEELITE
The code-runner is a server which gets code to be run from the UI server and runs the code. Its operation is to compile, execute and return the output of the code to the requesting UI server.

Because the code is unsafe(written by some-random-guy-on-the-internet), it should run the code securely. Ideally, each submission of code should be executed isolated from other processes, restricted from making any network calls and accessing the file system out of its scope. These restrictions are the same for all submissions.

The other side of the coin is to limit the resource usage of each submission mainly the CPU time(infinite loops, fork bombs, compiler bombs), memory usage(huge allocations, more...). These limits should be tweakable per submission.

To do this, we make use of Docker. Docker takes care of all the above details by running a submission in its own `container`, that is isolated from the external world - the `host` environment.

*Container*: A container is a light weight environment in which a process runs, isolated from others with several restrictions imposable depending on the requirements. It is an instance of an "image". The container is stopped after PID 1 inside that process exits.

*Image*: An image is a blueprint of an environment.

Docker has a client-server architecture. The docker daemon (dockerd) is a background process that does actual work. Docker clients request the docker daemon to do work. They implement a protocol for requests and responses. The protocol works over TCP, UNIX sockets(and possibly more connection protocols).

So, code-runner should also act as a docker client and request docker daemon to spawn containers and run processes inside them. This requires us to build images from which containers are to be spawned. The image should contain libraries that are to be used by the executable and any environment.

*Issue*: There is a subtlety here because of compiled and dynamic languages. For compiled languages, we can run the compiler in its own container which is always running and have another image that will contain the executable produced by the compiled language and run the executable in its own tiny environment like the "hello-world" example. But interpreted languages like python, ruby and others can't be done because they require the interpreter to be available. This problem is yet to be solved. Perhaps, a separate treatment to compiled and interpreted languages is necessary. But the focus is to get compiled languages(C) running first.

TODO: Still have to choose a light weight REST framework(no need for a full-blown web framework).

https://lebkowski.name/docker-volumes/

sudo usermod -a -G docker $USER

Ada
--bash
--C
Clojure
CoffeScript
--C++
C#
D
Elixir
Fortan
F#
Go
Groovy
Haskell
--Java 7 and 8
Node Js
Julia
Kotlin
LOLCODE
Lua
Obj C
Ocaml
Octave
Pascal
Perl
--PHP
--Python 2
--Python 3
R
Racket
Ruby
Rust
Lisp
Scala
Smalltalk
Swift
Tcl
VB.NET
Whitespace
