const hello = process.argv[2] || 'from function'
const text = `hello ${hello}!`
console.error(text)
console.error(JSON.stringify(process.argv))

throw new Error("Fail")
