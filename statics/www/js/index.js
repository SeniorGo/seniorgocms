fetch("/version")
  .then((req) => req.text())
  .then((version) => {
    document.getElementById("version").innerText = version;
  });

fetch("/auth/me")
  .then((req) => req.json())
  .then((payload) => {
    if (payload.error) {
      // No logged in
      document.getElementById("auth_login").style.display = "";
    } else {
      // Logged in
      console.log(payload);
      console.log(payload.nick);
      document.getElementById("auth_admin").style.display = "";
      document.getElementById("auth_logout").style.display = "";
      document.getElementById("auth_nick").style.display = "";
      document.getElementById("auth_login").style.display = "";
      document.getElementById("auth_nick").innerText = payload.nick;
    }
  });
