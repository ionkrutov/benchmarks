const turbo = require('turbo-http')

const server = turbo.createServer(function (req, res) {
  res.write(Buffer.from('Hello world!\n'))
})

server.listen(8080)