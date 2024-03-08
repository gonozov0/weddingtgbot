include .env
export

lint:
	go fmt ./...
	find . -name '*.go' -exec goimports -local containerh/ -w {} +
	find . -name '*.go' -exec golines -w {} -m 120 \;
	golangci-lint run ./...

zip_and_push_to_cloud:
	rm -f ./zip/weddingtgbot.zip
	zip -x ".*" -x "zip/*" -x "tests/*" -x "git-hooks/*" -r ./zip/weddingtgbot.zip .
	yc serverless function version create \
	   --function-name=weddingtgbot \
	   --runtime golang121 \
	   --entrypoint cmd/webhook_handler/webhook_handler.Handler \
	   --memory 128m \
	   --execution-timeout 5s \
	   --source-path ./zip/weddingtgbot.zip \
	   --service-account-id aje6c77di14c8bcs77ib \
	   --log-group-name default \
	   --network-name default \
	   --secret environment-variable=TG_BOT_TOKEN,id=e6qroi83kor4oin8r05q,version-id=e6qnhv1alns3r3j7tpjh,key=TG_BOT_TOKEN \
	   --secret environment-variable=YC_ACCESS_KEY_ID,id=e6qroi83kor4oin8r05q,version-id=e6qnhv1alns3r3j7tpjh,key=YC_ACCESS_KEY_ID \
	   --secret environment-variable=YC_SECRET_ACCESS_KEY,id=e6qroi83kor4oin8r05q,version-id=e6qnhv1alns3r3j7tpjh,key=YC_SECRET_ACCESS_KEY \
	   --environment YC_BUCKET_NAME=weddingtgbot
