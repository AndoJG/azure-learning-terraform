# Its Wednesday

Manages a Kubernetes deployment for a stupid joke

## Vars

`container_port`: the port that the container exposes  
default: 80

`service_port`: the port that the k8s service listens on  
default: 80

`replicas`: # of pod replicas to deploy  
default: 1

`image_tag`: version of the image to deploy  
default: latest

## Outputs

`public_address`: returns the public ip of the deployed service