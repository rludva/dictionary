# dictionary
Practice and learn dictionary..

## Dictionary structure
hello:		 		   ahoj, nazdar, čau
hello world:		 nazdar světe

### Terms:
* In the dictionary there are two terms: `item` and `content`:
  - hello -> item
  - ahoj -> content
* it is possible to have multiple contents for an item separated by coma `,`
* item and content is separated by colon `:`


## Practice all items
```
// Print help...
$ dictionary --help
$ dictionary practice --help

// Practice...
$ dictionary practice --all

// Practice 5 times item and its content...
$ dictionary practice --all --counter=5

// Practice only items 5 times...
$ dictionary practice --item --counter=5
```