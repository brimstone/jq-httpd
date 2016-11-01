jq-httpd
========

Send your JSON to a specially formatted URL, transform them, and send them on
their way!

Usage
-----
Start the server. It defaults to listening on port 8081. This can be changed
with the `PORT` environment variable.

Send JSON via POST, PUT, whatever to `/v1/jq/FILTER/to/DEST`

- FILTER: has to be a valid [jq](https://stedolan.github.io/jq/) filter,
urlencoded
- DEST: has to be a valid URL, urlencoded

If FILTER is left blank, the transformed JSON is returned to the client.

To actually use the DEST as a jq filter, use the `/v1/jq/FILTER/tq/DEST` url notation.

To use a FILTER of simply '.', it must be URL encoded to %2E.

Examples
--------
This requires a server on the same localhost as the server, on port 8082

Identify your vegetables:
```
curl -i -d '{"fruit": "watermelon"}' "localhost:8081/jq/%7Bvegetable%3A%20.name%7D/to/http%3A%2F%2F127.0.0.1%3A8082"
```

Use bits from the post data in the url with /tq/ intead of /to/:
```
curl -i -d "{\"date\": $(date +%s)}" "localhost:8081/v1/jq/$(urlencode '.')/tq/$(urlencode '"http://127.0.0.1:8082/"+ (.date | strftime("%Y-%m-%d"))')"
```

Slack echo bot example:
```
TODO
```

License
-------
AGPLv3
