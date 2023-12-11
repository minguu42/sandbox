terraform {
  required_version = ">= 1.0.0, < 2.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket  = "teratera-stg-tfstate"
    key     = "stg/s3/terraform.tfstate"
    region  = "ap-northeast-1"
    encrypt = true
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_s3_bucket" "tfstate" {
  bucket = "teratera-stg-tfstate"
  # 誤ってS3バケットを削除するのを防止する
  lifecycle {
    prevent_destroy = true
  }
}

# 明示的にバケットに対する全パブリックアクセスをブロックする
resource "aws_s3_bucket_public_access_block" "tfstate_block" {
  bucket                  = aws_s3_bucket.tfstate.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# デフォルトでサーバサイド暗号化を有効化する
resource "aws_s3_bucket_server_side_encryption_configuration" "tfstate_default" {
  bucket = aws_s3_bucket.tfstate.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# ステートファイルの完全な履歴が見れるように、バージョニングを有効化する
resource "aws_s3_bucket_versioning" "tfstate_enabled" {
  bucket = aws_s3_bucket.tfstate.id
  versioning_configuration {
    status = "Enabled"
  }
}
