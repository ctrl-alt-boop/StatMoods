if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        navigator.serviceWorker.register('sw.js')
            .then((registration) => {
                console.log('Service Worker registered with scope:', registration.scope);
            })
            .catch((error) => {
                console.error('Service Worker registration failed:', error);
            });
    });
}

function onInit() {
    let greetingElement = document.getElementById("greeting")
    greetingElement.innerText = resources.strings.greeting
}

function onClick() {
    let greetingElement = document.getElementById("greeting")
    greetingElement.innerText = resources.strings.thanks
}

onInit()