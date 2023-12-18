resource "aws_s3_bucket" "tfstate" {
  bucket = "${local.product}-${var.env}-tfstate"
  lifecycle {
    prevent_destroy = true
  }
}

# ステートファイルの完全な履歴が見れるように、バージョニングを有効化する
resource "aws_s3_bucket_versioning" "tfstate" {
  bucket = aws_s3_bucket.tfstate.id
  versioning_configuration {
    status = "Enabled"
  }
}
