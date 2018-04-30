# Max size for a shard

Your shards should exceed 32GB. If this is the case, you could add more shards to your indices.

You can get this informations, thanks to the following request :

```shell
GET /_cat/shards
```
