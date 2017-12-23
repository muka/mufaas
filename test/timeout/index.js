// this container will keep open

console.log("argv >> %j", process.argv)

process.stdin.on("data", function(chunk) {
    console.log("stdin >> ", chunk)
})
