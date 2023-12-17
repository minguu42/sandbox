provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      Product   = "teratera"
      Env       = "stg"
      Owner     = "minguu42"
      ManagedBy = "terraform"
    }
  }
}
