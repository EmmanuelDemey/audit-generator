# ES_HEAP_SIZE configuration

You should avoid allocate more than 50% of the total system memory to the JVM

```shell
GET /_cat/nodes?h=heap*&v
```
