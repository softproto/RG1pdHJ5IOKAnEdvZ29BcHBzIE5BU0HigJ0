# How to

## run
```
docker build . -t gogospace
docker run -p 8080:8080 gogospace
```
## test
```
go test -v ./...
```

# Description

Let’s say GogoApps is building an awesome space visualization tool. In order to accomplish this task a
beautiful database of space pictures and videos is required. Such database is going to be built using two
independent microservices.
A microservice that is able to download media files from provided urls and store them onto some
internal storage (called media-downloader)
A microservice that would prepare a list of urls for the first microservice to use (called url-collector)
While your teammates are building the media-downloader, your task is to create the latter - the urlcollector. You will be building a microservice responsible for gathering image URLs from the open NASA's APOD API.


# Requirements

The solution should consist of a HTTP server with one HTTP endpoint available. Endpoint should be
exposed under following path:

        GET /pictures?start_date=2020-01-04&end_date=2020-02-05

Since the NASA API publishes one image per day the start_date and end_date parameters define
range of pictures to be processed. Expected response from the endpoint should be a valid JSON message
containing all the urls in the following format:

        {“urls”: ["https://apod.nasa.gov/apod/image/2008/AlienThrone_Zajac_3807.jpg", ...]}

In case of an error the server is asked to return an appropriate HTTP status code and a descriptive JSON
message in the following format:

        {“error”: “error message”} .

As the provided date range in a single request might be broad, the NASA API should be queried
concurrently. However, in order not to be recognized as a malicious user, a limit of concurrent
requests to this external API must exist. Bear in mind, that this limit should never be exceeded
regardless of how many concurrent requests is the url-collector receiving.
## Note
- We are not using the Apod server's ability to handle a date range in a single request due to the Task Description.
- Considering the fact that there can be both successful and unsuccessful requests for date range, the collected errors are contained in the same report.


# Issues found

- The Apod service does not provide an error result in a unified format. For this reason, we cannot correctly collect errors from the Apod server.

Response code 400:
```
{
"code": 400,
"msg": "Bad Request: incorrect field passed. Allowed request fields for apod method are 'concept_tags', 'date', 'hd', 'count', 'start_date', 'end_date', 'thumbs'",
"service_version": "v1"
}
```

Response code 403
```
{
"error": {
"code": "API_KEY_INVALID",
"message": "An invalid api_key was supplied. Get one at https://api.nasa.gov:443"
}
}
```

# Demo results

- Successful request
```
{
"Urls": [
"https://apod.nasa.gov/apod/image/0104/distantsn_hst.jpg",
"https://apod.nasa.gov/apod/image/0104/auroraiceland_shs_big.jpg",
"https://apod.nasa.gov/apod/image/0104/quiddich_sts92.jpg",
"https://apod.nasa.gov/apod/image/0104/ngc1748_hst.jpg"
]
}
```

- The response contains errors
```
{
"Errors": [
"with getDatesList(2001-04-01, 2001-04-00) got error: parsing time \"2001-04-00\": day out of range"
]
}
```

- The response contains both errors and correct data
```
{
"Urls": [
"https://apod.nasa.gov/apod/image/0104/auroraiceland_shs_big.jpg",
"https://apod.nasa.gov/apod/image/0104/ngc1748_hst.jpg",
"https://apod.nasa.gov/apod/image/0104/quiddich_sts92.jpg"
],
"Errors": [
"with 2001-04-05 got error: 429 Too Many Requests",
"with 2001-04-04 got error: 429 Too Many Requests",
"with 2001-04-06 got error: 429 Too Many Requests"
]
}
```

- Demo screenshot

<img src="./images/res_URLs and Errors.png">