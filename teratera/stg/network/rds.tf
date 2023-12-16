resource "aws_rds_cluster" "main" {
  cluster_identifier              = "teratera-stg"
  engine                          = "aurora-mysql"
  engine_version                  = "8.0.mysql_aurora.3.04.1"
  database_name                   = "maindb"
  master_username                 = "foo"
  master_password                 = "bar"
  port                            = 3306
  vpc_security_group_ids          = [aws_security_group.aurora.id]
  db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.main.name
  db_subnet_group_name            = aws_db_subnet_group.main.name
  skip_final_snapshot             = true # 検証のため
  tags                            = {
    Name = "teratera-stg"
  }
}

resource "aws_rds_cluster_instance" "main" {
  identifier           = "teratera-stg-01"
  cluster_identifier   = aws_rds_cluster.main.id
  instance_class       = "db.t4g.medium"
  engine               = aws_rds_cluster.main.engine
  engine_version       = aws_rds_cluster.main.engine_version
  db_subnet_group_name = aws_rds_cluster.main.db_subnet_group_name
  tags                 = {
    Name = "teratera-stg-01"
  }
}

resource "aws_security_group" "aurora" {
  name   = "aurora"
  vpc_id = aws_vpc.main.id
  ingress {
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_rds_cluster_parameter_group" "main" {
  name        = "teratera-stg"
  description = ""
  family      = "aurora-mysql8.0"
  tags        = {
    Name = "teratera-stg"
  }
}

resource "aws_db_subnet_group" "main" {
  name        = "teratera-stg"
  description = ""
  subnet_ids  = [aws_subnet.private_a.id, aws_subnet.private_c.id]
  tags        = {
    Name = "teratera-stg"
  }
}
