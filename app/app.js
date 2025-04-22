import {callG5} from './g5.js';

document.getElementById("getButton").addEventListener("click", async () => {
    document.getElementById("response").textContent = await callG5({
        path: "GET/ping",
        data: ""
    });
});

document.getElementById("postForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const name = document.getElementById("nameInput").value;
    document.getElementById("response").textContent = await callG5({
        path: "POST/hello",
        data: JSON.stringify({name})
    });
});

