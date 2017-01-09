var info = document.getElementById("info");

info.innerHTML += "href: <strong>" + window.location.href + "</strong><br>";
info.innerHTML += "host: <strong>" + window.location.host + "</strong><br>";
info.innerHTML += "hostname: <strong>" + window.location.hostname + "</strong><br>";
info.innerHTML += "pathname: <strong>" + window.location.pathname + "</strong><br>";
info.innerHTML += "protocol: <strong>" + window.location.protocol + "</strong><br>";
info.innerHTML += "origin: <strong>" + window.location.origin + "</strong><br>";
info.innerHTML += "port: <strong>" + window.location.port + "</strong><br>";
info.innerHTML += "search: <strong>" + window.location.search + "</strong><br>";
info.innerHTML += "hash: <strong>" + window.location.hash + "</strong><br>";
