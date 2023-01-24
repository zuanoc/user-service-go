resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.rest-api-gateway.execution_arn}/*/*"
}

data "archive_file" "lambda_file" {
  type        = "zip"
  source_dir  = "${path.module}/../dist"
  output_path = "artifacts/bootstrap.zip"
}

resource "aws_lambda_function" "lambda" {
  function_name = "${local.service_name}-lambda"
  filename      = data.archive_file.lambda_file.output_path
  role          = aws_iam_role.lambda-role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2"
  architectures = ["x86_64"]
  source_code_hash = data.archive_file.lambda_file.output_base64sha256

  environment {
    variables = {
      "GIN_MODE": "release"
    }
  }
}

resource "aws_cloudwatch_log_group" "lambda_log_group" {
  name = "/aws/lambda/${aws_lambda_function.lambda.function_name}"

  tags = local.common_tags
}