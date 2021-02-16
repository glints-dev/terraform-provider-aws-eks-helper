terraform {
  required_providers {
    aws-eks-helper = {
      source  = "glints-dev/aws-eks-helper"
      version = "0.2.1"
    }
  }
}

provider "aws-eks-helper" {}

data "aws-eks-helper_kube_reserved" "t3a_large" {
  instance_type = "t3a.large"
  region        = "ap-southeast-1"
}
