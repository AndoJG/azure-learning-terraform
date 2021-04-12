resource "kubernetes_deployment" "itswednesday" {
  metadata {
    name = "its-wednesday"
    labels = {
      app = "its-wednesday"
    }
  }
  spec {
    replicas = var.replicas
    selector {
      match_labels = {
        app = "its-wednesday"
      }
    }
  
    template {
      metadata {
        labels = {
          app = "its-wednesday"
        }
      }

      spec {
        container {
          image = "breadwatcher/its-wednesday:${var.image_tag}"
          name  = "its-wednesday"
        }
      }
    }
  }
}

resource "kubernetes_service" "itswednesday" {
  metadata {
    name = "its-wednesday-service"
  }
  spec {
    selector = {
        app = "its-wednesday"
      }
    port {
      port        = var.service_port
      target_port = var.container_port
    }
    type = "LoadBalancer"
  }
}

resource "kubernetes_ingress" "itswednesday" {
  metadata {
    name = "its-wednesday-ingress"
  }

  spec {
    backend {
      service_name = kubernetes_service.itswednesday.metadata[0].name
      service_port = var.service_port
    }

    rule {
      http {
        path {
          backend {
            service_name = kubernetes_service.itswednesday.metadata[0].name
            service_port = var.service_port
          }
          
          path = "/"
        }
      }
    }
  }
}