resource "aws_dynamodb_table" "lock" {
  name         = "${local.product}-${var.env}-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"
  attribute {
    name = "LockID"
    type = "S"
  }
}
