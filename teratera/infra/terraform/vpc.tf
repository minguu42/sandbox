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
