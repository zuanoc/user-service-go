resource "aws_iam_role" "lambda-role" {
  name = "${local.service_name}-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  tags = local.common_tags
}

resource "aws_iam_role_policy" "lambda-role-policy" {
  name = "${local.service_name}-lambda-role-policy"
  role   = aws_iam_role.lambda-role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
  
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*"
    }  
  ]
}
EOF
}
