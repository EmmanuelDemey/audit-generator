# Split Brain

In order to avoid split brain, you should respect the following formula when configuring elasticsearch

```
discovery.zen.minimum_master_nodes: N/2 + 1
```

With N the number of ellible master nodes.
