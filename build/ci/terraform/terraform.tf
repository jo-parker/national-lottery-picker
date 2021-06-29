provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region
  version = "~> 2.7"
}

locals {
  aws_vpc_stack_name = "${var.aws_resource_prefix}-vpc-stack"
  aws_ecs_service_stack_name = "${var.aws_resource_prefix}-svc-stack"
  aws_ecs_cluster_name = "${var.aws_resource_prefix}-cluster"
  aws_ecs_service_name = "${var.aws_resource_prefix}-service"
  aws_ecs_execution_role_name = "${var.aws_resource_prefix}-ecs-execution-role"
}

resource "aws_cloudformation_stack" "vpc" {
  name = local.aws_vpc_stack_name
  template_body = file("cloudformation-templates/public-vpc.yml")
  capabilities = ["CAPABILITY_NAMED_IAM"]
  parameters = {
    ClusterName = local.aws_ecs_cluster_name
    ExecutionRoleName = local.aws_ecs_execution_role_name
  }
}
resource "aws_cloudformation_stack" "ecs_service" {
  name = local.aws_ecs_service_stack_name
  template_body = file("cloudformation-templates/public-service.yml")
  depends_on = [aws_cloudformation_stack.vpc]

  parameters = {
    ContainerMemory = 1024
    StackName = local.aws_vpc_stack_name
    ServiceName = local.aws_ecs_service_name
    RepositoryCredentials = var.aws_repository_creds
    Role = var.aws_task_role_arn
  }
}
