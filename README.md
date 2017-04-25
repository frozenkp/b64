# b64
b64 is a base64 encoder & decoder with different table. 

You can use default table "A~Za~z0~9+/" or a custom table.

## format
```bash
b64 (-d | -e) [-r outputTable.txt | -t table.txt] file.txt
```

## argument
### -d
```-d``` is the decode mode of b64

### -e
```-e``` is the encode mode of b64

### -r string
when ```-r``` is on, b64 will build a random table, and store the table in the file you specify. Then, encoding file with the random table.

**It should use with ```-e```.**

### -t string
when ```-t``` is on, b64 will decode/encode the file with table you specify.
