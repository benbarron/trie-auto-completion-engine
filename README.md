### Simple auto-complete engine implementation in golang.

### Quick start 
```shell
$ git clone https://github.com/benbarron/go-search-engine
$ cd go-search-engine
$ make build 
$ make run 
```

Initialize a trie
```shell
curl -XGET localhost:3000/trie/initialize
```

Add Words to tree 
```shell
curl -X POST http://localhost:3000/trie/enroll/:id \
  -H "Content-Type: application/json" \
  --data-binary "@words.json"
```

Preform autocompletion
```shell
curl -X POST http://localhost:3000/trie/search/:id \
  -H "Content-Type: application/json" \
  --data "{\"prefix\":\"A\"}"
```

### Structure of a prefix tree (trie)
![img.png](trie-image.png)

References
- https://gitlab.com/tsoding/trie
- https://en.wikipedia.org/wiki/Trie