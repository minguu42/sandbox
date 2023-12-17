resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags             = {
    Name = "teratera-stg"
  }
}

resource "aws_subnet" "public_a" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.0.0/20"
  availability_zone = "ap-northeast-1a"
  tags              = {
    Name = "teratera-stg-public-a"
  }
}

resource "aws_subnet" "private_a" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.16.0/20"
  availability_zone = "ap-northeast-1a"
  tags              = {
    Name = "teratera-stg-private-a"
  }
}

resource "aws_subnet" "public_c" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.80.0/20"
  availability_zone = "ap-northeast-1c"
  tags              = {
    Name = "teratera-stg-public-c"
  }
}

resource "aws_subnet" "private_c" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.96.0/20"
  availability_zone = "ap-northeast-1c"
  tags              = {
    Name = "teratera-stg-private-c"
  }
}

resource "aws_eip" "nat_a" {
  domain = "vpc"
  tags   = {
    Name = "teratera-stg-nat-a"
  }
}

resource "aws_eip" "nat_c" {
  domain = "vpc"
  tags   = {
    Name = "teratera-stg-nat-c"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
  tags   = {
    Name = "teratera-stg"
  }
}

resource "aws_nat_gateway" "a" {
  allocation_id = aws_eip.nat_a.id
  subnet_id     = aws_subnet.public_a.id
  tags          = {
    Name = "teratera-stg-a"
  }
  depends_on = [aws_internet_gateway.main]
}

resource "aws_nat_gateway" "c" {
  allocation_id = aws_eip.nat_c.id
  subnet_id     = aws_subnet.public_c.id
  tags          = {
    Name = "teratera-stg-c"
  }
  depends_on = [aws_internet_gateway.main]
}

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
    Name = "teratera-stg-public"
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

resource "aws_route_table" "private_a" {
  vpc_id = aws_vpc.main.id
  route {
    cidr_block = "10.0.0.0/16"
    gateway_id = "local"
  }
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.a.id
  }
  tags = {
    Name = "teratera-stg-private-a"
  }
}

resource "aws_route_table_association" "private_a" {
  subnet_id      = aws_subnet.private_a.id
  route_table_id = aws_route_table.private_a.id
}

resource "aws_route_table" "private_c" {
  vpc_id = aws_vpc.main.id
  route {
    cidr_block = "10.0.0.0/16"
    gateway_id = "local"
  }
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.c.id
  }
  tags = {
    Name = "teratera-stg-private-c"
  }
}

resource "aws_route_table_association" "private_c" {
  subnet_id      = aws_subnet.private_c.id
  route_table_id = aws_route_table.private_c.id
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
