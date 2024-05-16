document.getElementById("drop_zone").addEventListener("dragover", function(event) {
    event.preventDefault();
    document.getElementById("drop_zone").classList.add("dragover");
});

document.getElementById("drop_zone").addEventListener("dragleave", function(event) {
    event.preventDefault();
    document.getElementById("drop_zone").classList.remove("dragover");
});

document.getElementById("drop_zone").addEventListener("drop", function(event) {
    event.preventDefault();
    document.getElementById("drop_zone").classList.remove("dragover");
    const files = event.dataTransfer.files;
    handleFiles(files);
});

function handleFiles(files) {
    const data = new FormData();
    data.append("path", window.location.href.replace(/(http|https):\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}/g, ""));
    
    for (const file of files) {
        data.append("upload[]", file);
    }

    fetch("/action/upload", {
        method: "POST",
        body: data
    })
        .then(res => res.json())
        .then(data => {
            if (data.ok !== 1) {
                console.error("error!");
                return;
            }

            console.log(data);
            window.location.reload();
        })
}
