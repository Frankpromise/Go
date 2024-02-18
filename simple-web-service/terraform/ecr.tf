resource "aws_ecr_repository" "go-ecr" {
  name                 = "my-ecr-repo"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}