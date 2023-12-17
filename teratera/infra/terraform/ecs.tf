resource "aws_ecs_cluster" "main" {
  name = "teratera-stg"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecs_service" "main" {
  name    = "teratera-stg"
  cluster = aws_ecs_cluster.main.id
}

resource "aws_ecs_task_definition" "main" {
  family                   = "teratera-stg-api"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = ""
  network_mode             = "awsvpc"
  cpu                      = 256 # .25vCPU
  memory                   = 512 # 0.5GB
  container_definitions    = jsonencode([
    {
      name         = "main"
      image        = "service-first"
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
    }
  ])
  runtime_platform {
    cpu_architecture        = "X86_64"
    operating_system_family = "LINUX"
  }
}
