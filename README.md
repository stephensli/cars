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

# What?

Cars is a Sandbox environment which allows the execution of untrusted code in a fixed time frame, fixed memory
allocation
based on a selected programming language with the ability to provide expected output for user input testing. The facing
API allows posting untrusted code with some additional requirements into the queue. The loader(s) will run the
untrusted code inside a container with https://gvisor.dev/, tracking execution properties and verifying output.

## Supported Queues

* Nsq (https://nsq.io/)
* AWS SQS (https://aws.amazon.com/sqs/)

## Supported File Systems

* Local file system (development)
* AWS S3 Bucket (https://aws.amazon.com/s3/)

## Supported Programming Languages  

 * Python v3.9.x (https://www.python.org/)
 * NodeJs V16.x (https://nodejs.org/en/)


[license-badge]: https://img.shields.io/github/license/stephensli/cars?style=flat-square
[go-version-badge]: https://img.shields.io/github/go-mod/go-version/stephensli/cars?style=flat-square
[build-badge]: https://img.shields.io/github/workflow/status/stephensli/cars/go?style=flat-square
[release-version-badge]: https://img.shields.io/github/v/release/stephensli/cars?style=flat-square
