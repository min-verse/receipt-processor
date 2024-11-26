# Receipt Processor API

## Running the Code
```
go run cmd/api/main.go
```

## API Summary

This API, built with Golang and Chi, allows a client to both:

* `POST` a Receipt JSON to create a Receipt in the database with the assumed payload structure defined below
   * Path: `/receipts/process`
   * Example Response:
        ```json
        { "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
        ```
* `GET` a Receipt's "total points" based on certain criteria defined below
   * Path: `/receipts/{id}/points`
   * Example Response:
        ```json
            { "points": 32 }
        ```

## Structure of JSON Payload to Create Receipt in Database

```json
{
  "retailer": "Example Retailer Name",
  "purchaseDate": "YYYY-MM-DD",
  "purchaseTime": "HH:MM", // Assumes 24-hour time
  "items": [
    {
      "shortDescription": "Product Name 1",
      "price": "99.99"
    },{
      "shortDescription": "Product Name 2",
      "price": "19.99"
    }
    ...
  ],
  "total": "119.98"
}
```

## Criteria for Total Points on a Receipt

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of `0.25`.
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm.

---

## Data Models

I made a simulated "database" by creating a `DatabaseInterface` with shared methods to mock out an ORM (both retrieving and creating records in `receiptsDB` which is of the `DatabaseInterface`).

Also a part of the simulated "database" is the `receiptsTable` which mocks a key-value store where the UUID is the key and an instance of the Receipt Struct is the value.

I'm using Logrus as a structured logger for this application to log any errors.

* `POST /receipts/process`
   * Handled by `ProcessReceiptInformation`
* `GET /receipts/{id}/points`
   * Handled by `CalculateReceiptPoints`

I've added comments to each file to make the code more readable and the logic flow easier to track along with any assumptions I've made along the way.

![alt text](https://as2.ftcdn.net/v2/jpg/02/10/55/15/1000_F_210551556_gpiKP1jr7hkjxkclvldly0gPnZmzbzE8.jpg)