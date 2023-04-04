const btn = document.querySelector(".form > button")
if (btn) {
    btn.onclick = e => {
        localStorage.setItem('myCat', 'Tom')
        console.log(localStorage.getItem('myCat'))
    }
}
