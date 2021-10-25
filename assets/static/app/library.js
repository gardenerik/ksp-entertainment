(async function () {
    async function enqueue(id) {
        await fetch('/api/library/'+id+'/enqueue/', {
            method: 'POST'
        })
    }

    document.querySelectorAll(".js-lib-enqueue").forEach(e => {
        e.addEventListener("click", () => {enqueue(e.dataset.id)})
    })
})()
