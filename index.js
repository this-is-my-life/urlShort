/**
 * @name urlShort
 * @description PMH Studio의 단축 URL서비스입니다
 * @copyright (c) 2019. PMH Studio / PMH. all rights reserved
 * @license MIT
 * @author PMH Studio / PMH
 * @version 1.0.0
 */

const PORT = 4000

const express = require('express')
const ejs = require('ejs')
const fs = require('fs')
const app = express()

if (!fs.existsSync('./lib/')) {
  fs.mkdirSync('./lib/', (err) => { console.error(err) })
}

if (!fs.existsSync('./lib/urls.json')) {
  fs.writeFileSync('./lib/urls.json', '{}', (err) => { console.error(err) })
}

const db = require('./lib/urls.json')

app.get('/', (req, res) => {
  res.redirect('/make')
})

app.get('/:short', (req, res) => {
  if (req.params.short === 'make') {

    // EJS 부분--------
    ejs.renderFile('./views/make.ejs', { db: db }, (err, data) => {
      if (err) console.error(err)
      res.send(data)
    })

  } else if (!db[req.params.short]) {
    res.sendStatus(404)
  } else {
    res.redirect(db[req.params.short])
  }
})

app.get('/make/:short/:long', (req, res) => {
  db[req.params.short] = 'http://' + req.params.long
  fs.writeFile('./lib/urls.json', JSON.stringify(db), (err) => { if (err) console.error(err) } )
  res.send('<script>\nalert(\'처리됨: \' + document.location.origin + "/' + req.params.short + '")\nwindow.location.replace(document.location.origin + "/' + req.params.short + '")\n</script>')
})

app.listen(PORT, () => { console.log('urlShort Load Completed') })
