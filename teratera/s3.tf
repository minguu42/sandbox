resource "aws_s3_bucket" "tfstate" {
  bucket = "teratera-stg-tfstate"
  tags   = {
    Name = "teratera-stg-tfstate"
  }
  # 誤ってS3バケットを削除するのを防止する
  lifecycle {
    prevent_destroy = true
  }
}

# 明示的にバケットに対する全パブリックアクセスをブロックする
resource "aws_s3_bucket_public_access_block" "tfstate" {
  bucket                  = aws_s3_bucket.tfstate.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# デフォルトでサーバサイド暗号化を有効化する
resource "aws_s3_bucket_server_side_encryption_configuration" "tfstate" {
  bucket = aws_s3_bucket.tfstate.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# ステートファイルの完全な履歴が見れるように、バージョニングを有効化する
resource "aws_s3_bucket_versioning" "tfstate" {
  bucket = aws_s3_bucket.tfstate.id
  versioning_configuration {
    status = "Enabled"
  }
}
