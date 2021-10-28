(async function () {
    async function enqueue(id) {
        await fetch('/api/playlists/'+id+'/play/', {
            method: 'POST'
        })
    }

    document.querySelectorAll(".js-playlist-enqueue").forEach(e => {
        e.addEventListener("click", () => {enqueue(e.dataset.id)})
    })

    async function add_item(id) {
        const url = prompt("Video URL")
        if (url.trim() === "") {
            return
        }

        let body = new FormData()
        body.append('url', url)

        await fetch('/api/playlists/'+id+'/', {
            method: 'POST',
            body: body
        })
        window.location.reload()
    }

    document.querySelectorAll(".js-playlist-add-item").forEach(e => {
        e.addEventListener("click", () => {add_item(e.dataset.id)})
    })

    async function importPlaylist() {
        const name = prompt("Playlist name")
        if (name.trim() === "") {
            return
        }

        const url = prompt("Playlist URL")
        if (url.trim() === "") {
            return
        }

        let body = new FormData()
        body.append('url', url)
        body.append('name', name)

        await fetch('/api/playlists/import/', {
            method: 'POST',
            body: body
        })
        window.location.reload()
    }

    document.querySelectorAll(".js-playlist-import").forEach(e => {
        e.addEventListener("click", () => {importPlaylist()})
    })
})()
