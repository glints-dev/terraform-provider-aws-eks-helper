terraform {
  required_providers {
    aws-eks-helper = {
      source  = "glints/aws-eks-helper"
      version = "0.1.0"
    }
  }
}

provider "aws-eks-helper" {}

data "aws_eks_helper_kube_reserved" "t3a_large" {
  provider      = aws-eks-helper
  instance_type = "t3a.large"
  region        = "ap-southeast-1"
}
