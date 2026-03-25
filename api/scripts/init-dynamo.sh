#!/bin/sh
set -e

export AWS_ACCESS_KEY_ID=local
export AWS_SECRET_ACCESS_KEY=local
export AWS_REGION=us-east-1

echo "⏳ Waiting DynamoDB..."

until curl -s http://dynamodb:8000 > /dev/null; do
  sleep 1
done

echo "✅ DynamoDB ready"

echo "📦 Checking table: links"

# 👇 ignora erro do describe
if aws dynamodb describe-table \
  --table-name links \
  --endpoint-url http://dynamodb:8000 > /dev/null 2>&1; then

  echo "ℹ️ Table already exists"

else
  echo "📦 Creating table..."

  aws dynamodb create-table \
    --table-name links \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --endpoint-url http://dynamodb:8000

  echo "✅ Table created"
fi

echo "🚀 Done"