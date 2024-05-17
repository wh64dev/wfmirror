const form = document.getElementById("login");
const err = document.getElementById("err");

form.addEventListener("submit", (ev) => {
    ev.preventDefault();

    const path = document.getElementById("path");
    const username = form.querySelector("#username");
    const password = form.querySelector("#password");


    const encoded = new URLSearchParams();
    encoded.append("path", path.innerText);
    encoded.append("username", username.value);
    encoded.append("password", password.value);

    fetch("/action/login", {
        method: "POST",
        body: encoded
    }).then(res => res.json()).then(data => {
        if (data.status == 401) {
            err.innerText = "아이디 비번 틀림 ㅅㄱ";
            return;
        }
    });
});
