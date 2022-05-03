# TRUPHONE technical test

## WARNING!! THIS IS A BOILERPLATE AS ASKED IN THE TECHNICAL TEST DESCRIPTION!!

Test Overview:

This is a service boilerplate which has two main separate running services, 
auth-service and manager-service which holds the device-manager-service.

The application's architectures follow golang-standards 
(although not an official go-way of writing, it has been widely embraced by the go community and 
takes advantage of the go compiler) and the N-tier architecture is well-designed and strongly 
implemented and defended in this application model as well as followed by SOLID and 
clean-architecture principles.

- Why N-tier architecture?
  - N-tier architecture was chosen to make things easier in this boilerplate repo. How? 
  Domains are very well separated, and they are COMPLETELY isolated. 
  If a context needs to talk to or call another one - such as device-manager-service 
  accessing auth-service - this access is done by dependency injection, respecting SOLID principles, 
  making it easy to write, to read, to test, to execute and to compile.
  - Also, why dependency injection of contexts and domains? Well, since I 'boilerplated' a monolith, 
  if by the next few hours we decide to move towards a micro-service architecture, 
  that would be easily done and instead by passing the store, we would be passing a gRPC client, 
  for example, which would implement the SAME, LITERALLY THE SAME interface as the store adaptor, 
  so we would not need to change a thing!! Not even the tests (and we would have guaranteed 
  they would work as expected because we respected Liskov's injection principle!!); 
  we would only need to separate the contexts into their own private repos and be happy because 
  we know how to design an awesome api!! :)

I hope you all have liked the marvelous architectural design I made and, please, feel free to 
pose any questions :) 

Have a great day!!

Augusto Pietroski
