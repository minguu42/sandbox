terraform {
  backend "s3" {
    bucket         = "teratera-state-456"
    key            = "global/s3/terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "teratera-locks"
    encrypt        = true
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_s3_bucket" "terraform_state" {
  bucket = "teratera-state-456"
  lifecycle {
    prevent_destroy = true # 誤ってS3バケットを削除するのを防止する
  }
}

# ステートファイルの完全な履歴が見れるように、バージョニングを有効化する
resource "aws_s3_bucket_versioning" "enabled" {
  bucket = aws_s3_bucket.terraform_state.id
  versioning_configuration {
    status = "Enabled"
  }
}

# デフォルトでサーバサイド暗号化を有効化する
resource "aws_s3_bucket_server_side_encryption_configuration" "default" {
  bucket = aws_s3_bucket.terraform_state.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# 明示的にバケットに対する全パブリックアクセスをブロックする
resource "aws_s3_bucket_public_access_block" "public_access" {
  bucket                  = aws_s3_bucket.terraform_state.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_dynamodb_table" "terraform_locks" {
  name         = "teratera-locks"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }
}
