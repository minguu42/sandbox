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
