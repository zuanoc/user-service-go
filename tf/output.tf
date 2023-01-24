output "base_url" {
  value = "${aws_api_gateway_deployment.gateway-deployment.invoke_url}dev/users"
}