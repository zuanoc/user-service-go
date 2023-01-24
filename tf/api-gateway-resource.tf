resource "aws_api_gateway_resource" "service-resource" {
  rest_api_id = aws_api_gateway_rest_api.rest-api-gateway.id
  parent_id = aws_api_gateway_rest_api.rest-api-gateway.root_resource_id
  path_part = "{proxy+}"
}

resource "aws_api_gateway_method" "service-resource-method" {
  rest_api_id = aws_api_gateway_rest_api.rest-api-gateway.id
  resource_id = aws_api_gateway_resource.service-resource.id
  http_method = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "service-resource-integration" {
  rest_api_id          = aws_api_gateway_rest_api.rest-api-gateway.id
  resource_id          = aws_api_gateway_resource.service-resource.id
  http_method          = aws_api_gateway_method.service-resource-method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda.invoke_arn
}
