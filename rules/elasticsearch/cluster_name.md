# Cluster Name

You should change the cluster name, in order to avoir other nodes to connect to it, if we still use the default cluster name.

In order to check if the cluster.name is defined, you can execute the following request ;

```json
GET http://localhost:9200/
```

You will get this response :

```json
{
  "name": "KCeLjRo",
  "cluster_name": "docker-cluster",
  "cluster_uuid": "2NEqti6KRWOMyYQPAJ8HIA",
  "version": {
    "number": "6.2.2",
    "build_hash": "10b1edd",
    "build_date": "2018-02-16T19:01:30.685723Z",
    "build_snapshot": false,
    "lucene_version": "7.2.1",
    "minimum_wire_compatibility_version": "5.6.0",
    "minimum_index_compatibility_version": "5.0.0"
  },
  "tagline": "You Know, for Search"
}
```

This parameter is named `cluster.name` in the `elasticsearch.yml`.
