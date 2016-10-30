jq-httpd
========

Send your JSON to a specially formatted URL, transform them, and send them on
their way!

Usage
-----
Start the server. It defaults to listening on port 8081.

Send JSON via POST, PUT, whatever to `/jq/FILTER/to/DEST`

- FILTER: has to be a valid jq filter, urlencoded
- DEST: has to be a valid URL, urlencoded

If FILTER is left blank, the transformed JSON is returned to the client.

Examples
--------

Identify your vegetables:
```
curl -i -d '{"fruit": "watermelon"}' "localhost:8081/jq/%7Bvegetable%3A%20.name%7D/to/http%3A%2F%2F127.0.0.1%3A8082"
```
This requires a server on the same localhost as the server, on port 8082

Slack echo bot example:
```
TODO
```

License
-------
AGPLv3
