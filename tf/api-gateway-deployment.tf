resource "aws_api_gateway_deployment" "gateway-deployment" {
  depends_on = [
    aws_api_gateway_rest_api_policy.rest-gateway-policy,
    aws_api_gateway_integration.service-resource-integration
  ]

  rest_api_id = aws_api_gateway_rest_api.rest-api-gateway.id

  lifecycle {
    create_before_destroy = true
  }

  variables = {
    version = timestamp()
  }
}

resource "aws_api_gateway_stage" "gateway-stage" {
  deployment_id = aws_api_gateway_deployment.gateway-deployment.id
  rest_api_id = aws_api_gateway_rest_api.rest-api-gateway.id
  stage_name = local.env
  tags = local.common_tags
}