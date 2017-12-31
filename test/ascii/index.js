
if (!process.argv[2]) {
    console.error("Missing text")
    process.exit(1)
}

const text = process.argv[2]
const font = process.argv[3] || 'Doom'

const fonts = ['Doom', 'rusted']
if (fonts.indexOf(font) == -1) {
    console.error("Unknown font, avail: %s", JSON.stringify(fonts))
    process.exit(1)
}

require('ascii-art').font(text, font, (asci) => console.log(asci))
