let g5Ready = null;

const go = new Go();

async function waitForGlobalFunction(name, timeout = 1000) {
    const interval = 10;
    const maxTries = timeout / interval;
    let tries = 0;
    while (typeof window[name] !== "function") {
        if (++tries > maxTries) throw new Error(`Function "${name}" not found on window`);
        await new Promise(resolve => setTimeout(resolve, interval));
    }
}

async function initG5() {
    if (!g5Ready) {
        const result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject);
        go.run(result.instance); // don't await — it never resolves
        await waitForGlobalFunction("g5");
        g5Ready = true;
    }
}

export async function callG5(inputObj) {
    await initG5();
    const inputStr = JSON.stringify(inputObj);
    return window.g5(inputStr); // use `window.g5` explicitly
}
