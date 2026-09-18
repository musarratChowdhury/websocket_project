
console.log("Hello, World!");

window.onload = () => {
    if(window["WebSocket"]) {
        const ws =  new WebSocket("ws://" + window.location.host + "/ws");
    }
}