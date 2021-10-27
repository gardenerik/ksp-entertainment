(async function () {
    const template = '<div class="mb-2 border border-gray-700 py-2 px-3 rounded flex">' +
        '<a href="#">{{o.library_item.name}} <span class="text-xs text-gray-400">#{{o.library_item.id}}</span></a>' +
        '<button class="text-red-600 hover:text-red-500 hover:underline text-sm ml-auto js-queue-remove" data-id="{{o.id}}">Remove</button>' +
    '</div>';

    async function deleteQueueItem(id) {
        await fetch('/api/queue/'+id+'/', {
            method: 'DELETE'
        })
        await loadQueue()
    }

    async function clearQueue() {
        await fetch('/api/queue/', {
            method: 'DELETE'
        })
        await loadQueue()
    }

    async function addToQueue() {
        const url = prompt("Video URL")
        if (url.trim() === "") {
            return
        }

        let body = new FormData()
        body.append('url', url)

        await fetch('/api/queue/', {
            method: 'POST',
            body: body
        })
        await loadQueue()
    }

    async function loadQueue() {
        let res = await fetch('/api/queue/')
        if (!res.ok) {
            alert("Could not load queue.")
            return
        }
        let data = (await res.json()).queue

        document.getElementById("js-queue-list").innerHTML = ""
        if (data.length) {
            document.getElementById("js-queue-empty").classList.add("hidden")
        } else {
            document.getElementById("js-queue-empty").classList.remove("hidden")
        }

        for (const item of data) {
            document.getElementById("js-queue-list").innerHTML += Mustache.render(template, {o: item})
        }

        document.querySelectorAll(".js-queue-remove").forEach(e => {
            e.addEventListener("click", () => {deleteQueueItem(e.dataset.id)})
        })
    }

    await loadQueue()
    setInterval(loadQueue, 15000)
    document.getElementById("js-queue-add").addEventListener('click', addToQueue)
    document.getElementById("js-queue-clear").addEventListener('click', clearQueue)
})()
