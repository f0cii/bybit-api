swagger generate client /C api.yaml

# https://openapi-generator.tech
# 安装:
# npm install @openapitools/openapi-generator-cli -g
openapi-generator generate -g go --additional-properties=prependFormOrBodyParameters=true \
    -o out -i petstore.yaml

npx openapi-generator generate -g go -o ./out -i api.yaml
# npx openapi-generator generate -i petstore.yaml -g ruby -o /tmp/test/
