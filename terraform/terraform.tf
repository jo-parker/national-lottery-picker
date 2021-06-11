provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region
  version = "~> 2.7"
}

locals {
  # The name of the CloudFormation stack to be created for the ECS service and related resources
  aws_ecs_service_stack_name = "${var.aws_resource_prefix}-svc-stack"
  # The name of the ECS cluster to be created
  aws_ecs_cluster_name = "${var.aws_resource_prefix}-cluster"
  # The name of the ECS service to be created
  aws_ecs_service_name = "${var.aws_resource_prefix}-service"
  # The name of the execution role to be created
  aws_ecs_execution_role_name = "${var.aws_resource_prefix}-ecs-execution-role"
}

# Note: creates task definition and task definition family with the same name as the ServiceName parameter value
resource "aws_cloudformation_stack" "ecs_service" {
  name = local.aws_ecs_service_stack_name
  template_body = file("cloudformation-templates/public-service.yml")

  parameters = {
    ContainerMemory = 1024
		ClusterName = local.aws_ecs_cluster_name
    ServiceName = local.aws_ecs_service_name
    ECSTaskExecutionRole = local.aws_ecs_execution_role_name
  }
}
