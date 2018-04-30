# Disable Swap

You should disable swapping on your nodes.

To check the current configuration, you can execute the following request :

```shell
GET http://localhost:9200/_nodes?filter_path=**.mlockall
```

You will get this response :

```json
{
  "nodes": {
    "KCeLjRoBSKWgw0Yr33y7lA": {
      "process": {
        "mlockall": false
      }
    }
  }
}
```

You should define the `bootstrap.memory_lock:true` in the `elasticsearch.yml` file.
