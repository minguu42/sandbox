resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags             = {
    Name = "${local.product}-${var.env}"
  }
}

resource "aws_subnet" "public_a" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.0.0/20"
  availability_zone = "ap-northeast-1a"
  tags              = {
    Name = "${local.product}-${var.env}-public-a"
  }
}

resource "aws_subnet" "private_a" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.16.0/20"
  availability_zone = "ap-northeast-1a"
  tags              = {
    Name = "${local.product}-${var.env}-private-a"
  }
}

resource "aws_subnet" "public_c" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.80.0/20"
  availability_zone = "ap-northeast-1c"
  tags              = {
    Name = "${local.product}-${var.env}-public-c"
  }
}

resource "aws_subnet" "private_c" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.96.0/20"
  availability_zone = "ap-northeast-1c"
  tags              = {
    Name = "${local.product}-${var.env}-private-c"
  }
}

# TODO: コメントアウトする
#resource "aws_eip" "nat_a" {
#  domain = "vpc"
#  tags   = {
#    Name = "${local.product}-${var.env}-nat-a"
#  }
#}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
  tags   = {
    Name = "${local.product}-${var.env}"
  }
}

# TODO: コメントアウトする
#resource "aws_nat_gateway" "a" {
#  allocation_id = aws_eip.nat_a.id
#  subnet_id     = aws_subnet.public_a.id
#  tags          = {
#    Name = "${local.product}-${var.env}-a"
#  }
#  depends_on = [aws_internet_gateway.main]
#}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id
  route {
    cidr_block = "10.0.0.0/16"
    gateway_id = "local"
  }
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }
  tags = {
    Name = "${local.product}-${var.env}-public"
  }
}

resource "aws_route_table_association" "public_a" {
  subnet_id      = aws_subnet.public_a.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_c" {
  subnet_id      = aws_subnet.public_c.id
  route_table_id = aws_route_table.public.id
}

# TODO: コメントアウトする
#resource "aws_route_table" "private" {
#  vpc_id = aws_vpc.main.id
#  route {
#    cidr_block = "10.0.0.0/16"
#    gateway_id = "local"
#  }
#  route {
#    cidr_block     = "0.0.0.0/0"
#    nat_gateway_id = aws_nat_gateway.a.id
#  }
#  tags = {
#    Name = "${local.product}-${var.env}-private"
#  }
#}
#
#resource "aws_route_table_association" "private_a" {
#  subnet_id      = aws_subnet.private_a.id
#  route_table_id = aws_route_table.private.id
#}
#
#resource "aws_route_table_association" "private_c" {
#  subnet_id      = aws_subnet.private_a.id
#  route_table_id = aws_route_table.private.id
#}

resource "aws_security_group" "alb" {
  name   = "${local.product}-${var.env}-alb"
  vpc_id = aws_vpc.main.id
}

resource "aws_vpc_security_group_ingress_rule" "alb_ingress" {
  security_group_id = aws_security_group.alb.id
  from_port         = 80
  to_port           = 80
  ip_protocol       = "tcp"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_vpc_security_group_egress_rule" "alb_egress" {
  security_group_id = aws_security_group.alb.id
  from_port         = 0
  to_port           = 0
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_security_group" "rds" {
  name   = "${local.product}-${var.env}-rds"
  vpc_id = aws_vpc.main.id
}

resource "aws_vpc_security_group_ingress_rule" "rds_ingress" {
  security_group_id = aws_security_group.rds.id
  from_port         = 3306
  to_port           = 3306
  ip_protocol       = "tcp"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_vpc_security_group_egress_rule" "rds_egress" {
  security_group_id = aws_security_group.rds.id
  from_port         = 0
  to_port           = 0
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}
