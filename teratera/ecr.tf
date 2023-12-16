resource "aws_ecr_repository" "main" {
  name = "teratera-stg"
  image_tag_mutability = "IMMUTABLE"
}
