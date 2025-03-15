const loginBtn = document.getElementById("auth_login");
const adminBtn = document.getElementById("auth_admin");
const logoutBtn = document.getElementById("auth_logout");
const nickLink = document.getElementById("auth_nick");
const avatarBtn = document.getElementById("auth_avatar");
const avatarImg = document.getElementById("auth_avatar_image");
const versionContainer = document.getElementById("version");

fetch("/version")
  .then((req) => req.text())
  .then((version) => {
    if (versionContainer) {
      versionContainer.innerText = version;
    }
  });

fetch("/auth/me")
  .then((req) => req.json())
  .then((payload) => {
    if (payload.error) {
      // No logged in
      loginBtn.style.display = "inline-flex";
      adminBtn.style.display = "none";
      logoutBtn.style.display = "none";
      nickLink.style.display = "none";
      nickLink.innerText = "";
    } else {
      // Logged in
      loginBtn.style.display = "none";
      adminBtn.style.display = "inline-flex";
      logoutBtn.style.display = "inline-flex";
      nickLink.style.display = "inline-flex";
      nickLink.innerText = payload.nick;
      avatarBtn.style.display = "inline-flex";
      avatarImg.src = payload.picture;
    }
  });
