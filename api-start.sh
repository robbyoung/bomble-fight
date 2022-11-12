export ENV=LOCAL
export VERSION=VERSION
export PORT=3001
export FIXTURES=./fixtures.json
cd apps/api/cmd/api-service
go build && ./api-service