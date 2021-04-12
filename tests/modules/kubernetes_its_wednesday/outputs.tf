output "public_address" {
  value = kubernetes_service.itswednesday.status.0.load_balancer.0.ingress.0.ip
}
