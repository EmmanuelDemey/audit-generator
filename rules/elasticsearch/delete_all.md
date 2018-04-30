# Deleting all indices

For securing reasons, you should not be able to delete all indices in one request.

You should disable this functionnality thanks to the `action.destructive_requires_name` parameter, used in the `elasticsearch.yml` file
