## Go Simple LRU Cache
Simple LRU Cache server to demonstrate least recently used algorithm
</br>
To build with Docker, from the root run:
```
docker build -t <image-name> .
```

Run the server
```
docker run --name simple-lrucache -p 8080:800 -e CACHE_SIZE=<cache-size> <image-name>
```

Default cache size is 10
</br>
To add an item to the cache:
POST request to `http://localhost:8080/cache/item` with body
```
{
    "key": "key1",
    "value": "value1"
}
```
</br>
To get item from the cache:
GET reuest to `http://localhost:8080/cache/item/{key}` 