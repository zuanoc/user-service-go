resource "aws_api_gateway_rest_api" "rest-api-gateway" {
  name = "users"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
  binary_media_types = ["*/*"]

  tags = local.common_tags
}

resource "aws_api_gateway_rest_api_policy" "rest-gateway-policy" {
  rest_api_id = aws_api_gateway_rest_api.rest-api-gateway.id

  policy = <<EOF
    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Effect": "Allow",
          "Principal": "*",
          "Action": "execute-api:Invoke",
          "Resource": "${aws_api_gateway_rest_api.rest-api-gateway.execution_arn}/*"
        }
      ]
    }
  EOF
}
