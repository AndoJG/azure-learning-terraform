variable "service_port" {
  description = "port the k8s service should listen on"
  default = 80
  type = number
}

variable "container_port" {
  description = "port the k8s container is listening on"
  default = 80
  type = number
}

variable "replicas" {
  description = "number of pod replicas"
  default = 1
  type = number
}

variable "image_tag" {
  description = "version of the image to use"
  default = "latest"
  type = string
}
