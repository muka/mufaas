const ascii = require('ascii-art')

const hello = process.argv[2] || 'from function'

//NOTE this is used in unit test to validate input / output
//     (see ../../docker-api/exec_test.go)
const text = `hello ${hello}!`

ascii.font(text, 'Doom', (out) => {
    console.log(text)
    console.log(JSON.stringify(process.argv))
    console.log(out)
})
