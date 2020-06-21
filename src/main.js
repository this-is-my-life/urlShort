function sub () {
  if (document.forms[0].do.value !== 'create') {
    alert('경고: 잘못된 명령을 수행중입니다')
    return false
  }

  if (!document.forms[0].short.value) {
    alert('출발지가 입력되지 않았습니다\n랜덤 기능을 사용해 보시는건 어떠신가요?')
    return false
  }

  if (!document.forms[0].long.value) {
    alert('목적지가 입력되지 않았습니다')
    return false
  }

  if (!(document.forms[0].long.value.startsWith('https://') || document.forms[0].long.value.startsWith('http://'))) {
    alert('목적지는 "http://"나 "https://"로 시작해야 합니다\nex) https://google.com/get/noto')
    return false
  }

  if (!document.forms[0].short.value.startsWith('/')) {
    document.forms[0].short.value = '/' + document.forms[0].short.value
  }
}

function random () {
  let str = '', len = Math.floor(Math.random() * 7) + 3
  for (let c = 0; c < len; c++) {
    str += String.fromCharCode(Math.floor(Math.random() * 0x2BAF) + 0xAC00)
  }
  document.forms[0].short.value = str
}
