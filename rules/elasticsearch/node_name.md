# Disable Swap

You should change the name of every node of your cluster

To check the current configuration, you can execute the following request :

```shell
GET http://localhost:9200/_nodes
```

You will get this response :

```json
{
  "_nodes": {
    "total": 1,
    "successful": 1,
    "failed": 0
  },
  "cluster_name": "docker-cluster",
  "nodes": {
    "KCeLjRoBSKWgw0Yr33y7lA": {
      "name": "KCeLjRo"
    }
  }
}
```

You should define the `node.name:true` in the `elasticsearch.yml` file.
