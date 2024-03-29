<div align="center">

# CARS

![License][license-badge]
![Build][build-badge]
![Go][go-version-badge]
![Version][release-version-badge]

Cars is an **Compile And Run Sandbox**.

</div>


<p align="center">
	<img src="./assets/simple-design.svg" alt="Size Limit CLI" width="1080">
</p>

## What?

Cars is a Sandbox environment which allows the execution of untrusted code in a fixed time frame, fixed memory
allocation
based on a selected programming language with the ability to provide expected output for user input testing. The facing
API allows posting untrusted code with some additional requirements into the queue. The loader(s) will run the
untrusted code inside a container with https://gvisor.dev/, tracking execution properties and verifying output.

## Documentation

* [Setup Guide (Running locally, Development)](./docs/RUNNING_LOCALLY.md)
* [API Endpoints (Compiling, Templates, Languages)](./docs/ENDPOINTS.md)

## Supported Queues

* Nsq (https://nsq.io/)
* AWS SQS (https://aws.amazon.com/sqs/)

## Supported File Systems

* Local file system (development)
* AWS S3 Bucket (https://aws.amazon.com/s3/)

## Supported Programming Languages

| Language | Version      | Url                          | Time | Memory |
|----------|--------------|------------------------------|------|--------|
| C        | GCC 12.3.x   | https://gcc.gnu.org          | ✅️   | ✅️     |
| C++      | GCC 12.3.x   | https://gcc.gnu.org          | ✅️   | ✅️     |
| C#       | .NET 8.0     | https://dotnet.microsoft.com | ✅️   | ✅️ ️    |
| F#       | .NET 8.0     | https://dotnet.microsoft.com | ✅️   | ✅️     |
| Java     | OpenJDK 21.0 | https://openjdk.java.net     | ✅️   | ✅️     |
| Kotlin   | 1.9.21       | https://kotlinlang.org/      | ✅️   | ✅️     |
| Scala    | 3.1.2        | https://www.scala-lang.org/  | ✅️   | ✅️     |
| NodeJs   | 20.x.x       | https://nodejs.org           | ✅️   | ✅️     |
| Python2  | 2.7.x        | https://pypy.org             | ✅️   | ✅️     |
| Python3  | 3.10.x       | https://pypy.org             | ✅️   | ✅️     |
| Go       | 1.21.x       | https://go.dev               | ✅️   | ✅️     |
| Ruby     | 3.2.x        | https://ruby-lang.org        | ✅️   | ✅️     |
| Rust     | 1.74.x       | https://rust-lang.org        | ✅️   | ✅️     |
| PHP      | 8.2.x        | https://www.php.net/         | ✅️   | ✅️     |

## gVisor (https://gvisor.dev/)

gVisor is an application kernel, written in Go, that implements a substantial portion of the Linux system call
interface. It provides an additional layer of isolation between running applications and the host operating system.


[license-badge]: https://img.shields.io/github/license/stephensli/Cars?style=flat-square

[go-version-badge]: https://img.shields.io/github/go-mod/go-version/stephensli/Cars?style=flat-square

[build-badge]: https://img.shields.io/github/workflow/status/stephensli/cars/Go?style=flat-square

[release-version-badge]: https://img.shields.io/github/v/release/stephensli/Cars?style=flat-square
