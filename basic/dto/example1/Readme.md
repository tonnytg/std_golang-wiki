How to test your dto:

    curl --location --request POST 'http://localhost:8080/product' \
        --header 'Content-Type: text/plain' \
        --data-raw '{
            "id": "1",
            "name":"test",
            "price": 0.1,
            "quantity": 1,
            "status":"test"
         }'