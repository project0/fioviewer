# Fioviewer
Visualize your [Fio](https://github.com/axboe/fio) logs in the Browser

> Note: This is just an early release and may not even be perfect

## Build and run
``` bash
# install dependencies
npm install

# build for production with minification
npm run build

# run server
go run fioviewer.go -listen ":8080" -dir /your/path/to/your/fiologs -static ./dist
```
