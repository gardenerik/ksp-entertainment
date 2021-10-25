(async function () {
    let currentState = null

    async function loadCurrent() {
        let res = await fetch('/api/playback/now/')
        if (!res.ok) {
            alert("Could not load current song.")
            return
        }

        let data = await res.json()
        currentState = data

        const disabledClass = "text-gray-600"

        document.querySelectorAll("#js-pb-wrapper > *").forEach(e => {
            e.classList.add(disabledClass)
        })

        if (data.playing) {
            document.getElementById("js-current-name").innerText = data.current.library_item.name

            document.querySelectorAll("#js-pb-wrapper > *").forEach(e => {
                e.classList.remove(disabledClass)
            })

            if (!data.paused) {
                document.getElementById("js-pb-pause").classList.remove("hidden")
                document.getElementById("js-pb-play").classList.add("hidden")
            } else {
                document.getElementById("js-pb-pause").classList.add("hidden")
                document.getElementById("js-pb-play").classList.remove("hidden")
            }

        } else {
            document.getElementById("js-current-name").innerText = "(not playing)"

            document.getElementById("js-pb-pause").classList.add("hidden")
            document.getElementById("js-pb-play").classList.remove("hidden")

            if (data.queue_stopped) {
                document.getElementById("js-pb-play").classList.remove(disabledClass)
            }
        }
    }

    async function playbackChange(action) {
        await fetch('/api/playback/'+action+'/', {
            method: 'POST',
        })
        await loadCurrent()
    }

    await loadCurrent()
    setInterval(loadCurrent, 15000)

    document.getElementById("js-pb-pause").addEventListener('click', () => playbackChange('pause'))
    document.getElementById("js-pb-play").addEventListener('click', () => {
        if (currentState.paused) {
            playbackChange('pause')
        } else {
            playbackChange('resume')
        }
    })
    document.getElementById("js-pb-skip").addEventListener('click', () => playbackChange('skip'))
    document.getElementById("js-pb-stop").addEventListener('click', () => playbackChange('stop'))
})()
