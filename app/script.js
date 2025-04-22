const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);

    // POST handler
    document.getElementById("postForm").addEventListener("submit", (e) => {
        e.preventDefault();
        const name = document.getElementById("nameInput").value;
        const input = {
            path: "POST/hello",
            data: JSON.stringify({ name })
        };
        const result = g5(JSON.stringify(input));
        document.getElementById("response").textContent = result;
    });

    // GET handler
    document.getElementById("getButton").addEventListener("click", () => {
        const input = {
            path: "GET/ping",
            data: ""
        };
        const result = g5(JSON.stringify(input));
        document.getElementById("response").textContent = result;
    });
});
