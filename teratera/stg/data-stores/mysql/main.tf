terraform {
  backend "s3" {
    bucket         = "teratera-state-456"
    key            = "stg/data-stores/mysql/terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "teratera-locks"
    encrypt        = true
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_db_instance" "example" {
  identifier_prefix   = "teratera"
  engine              = "mysql"
  allocated_storage   = 10
  instance_class      = "db.t2.micro"
  skip_final_snapshot = true
  db_name             = "example_database"
  username            = var.db_username
  password            = var.db_password
}
