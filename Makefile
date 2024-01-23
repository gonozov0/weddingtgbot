zip_and_push_to_s3:
	zip -r ./zip/weddingtgbot.zip .
	s3cmd put ./zip/weddingtgbot.zip s3://weddingtgbot/
