fetch("/version")
  .then((req) => req.text())
  .then((version) => {
    document.getElementById("version").innerText = version;
  });
