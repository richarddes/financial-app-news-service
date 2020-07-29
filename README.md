# News-Service
This service fetches news from the *NewsAPI* service and saves them in a database. Furhermore, it returns the news saved in the db.

## Setup 
To run this service independently, you have to install all necessary dependencies, like so:
```sh
go get ./...
```
To run the app, run:
```sh
go run main.go
```

## Inner workings
This service fetches the top business headlines, in the specified languages, from the *NewsAPI* service every 15 minutes and saves them in the specified database. It only has one endpoint */api/news/top-headlines* (for now), which returns the saved top headlines in the language specified in the *Lang* header, in json format. 

## Contributing 
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
MIT License. Click [here](https://choosealicense.com/licenses/mit/) or see the LICENSE file for details.