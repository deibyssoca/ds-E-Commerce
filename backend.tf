# backend.tf
terraform {
  backend "s3" {
    # Replace this with your bucket name!
    bucket = "ds-e-commerce-bucket"
    key    = "go-lambda-test.tfstate"
    region = "us-east-1"
  }
}