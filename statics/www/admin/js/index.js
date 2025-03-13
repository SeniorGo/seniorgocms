const list = document.getElementById("list");

const fakeHeaders = {
  'X-Glue-Authentication': JSON.stringify({user: {id: 'user-fake-id'}}),
};

let itemSelected = null;
function AddPost(post) {
  const item = document.createElement("li");
  item.classList.add("list-group-item");
  item.addEventListener(
    "click",
    () => {
      if (itemSelected != null) {
        itemSelected.classList.remove("selected");
      }
      itemSelected = item;
      itemSelected.classList.add("selected");
      EditPost(post);
    },
    true
  );
  item.textContent = post.title;
  list.appendChild(item);
  return item;
}

let lastPost = null;
function EditPost(post) {
  if (lastPost != null) {
    // todo: save post? or warn user?
  }
  lastPost = post;
  document.getElementById("editor").style.display = "";
  document.getElementById("editor_id").innerText = post.id;
  document.getElementById("editor_created").innerText = post.creation_time;
  document.getElementById("editor_modified").innerText = post.modification_time;
  document.getElementById("editor_title").value = post.title;
  document.getElementById("editor_body").value = post.body;
}

function SavePost(post) {
  fetch(`/v0/posts/${encodeURIComponent(post.id)}`, {
    method: "PATCH",
    headers: fakeHeaders,
    body: JSON.stringify(post),
  })
    .then((resp) => resp.json())
    .then((payload) => {
        if (payload.error) {
            alert(payload.error.description);
            return;
        }
        Object.assign(lastPost, payload);
        EditPost(lastPost);
    });
}

function DeletePost(post) {
  fetch(`/v0/posts/${encodeURIComponent(post.id)}`, {
    method: "DELETE",
    headers: fakeHeaders,
  }).then((resp) => {
      if (resp.status !== 204) {
          resp.json().then(payload => {
              alert(payload.error.description)
          })
          return;
      }
    document.getElementById("editor").style.display = "none";
    itemSelected.style.display = "none";
    itemSelected = null;
  });
}

fetch("/v0/posts", {headers: fakeHeaders})
  .then((resp) => resp.json())
  .then((posts) => {
    posts.forEach(AddPost);
  });

document.getElementById("button_create").addEventListener(
  "click",
  () => {
    const post = { title: "Nuevo post", body: "Escribe algo asombroso" };
    fetch("/v0/posts", {
        method: "POST",
        headers: fakeHeaders,
        body: JSON.stringify(post),
    })
      .then((resp) => resp.json())
      .then((post) => {
        AddPost(post);
        //EditPost(post);
      });
  },
  true
);

document.getElementById("button_save").addEventListener(
  "click",
  () => {
    SavePost({
      id: lastPost.id,
      title: document.getElementById("editor_title").value,
      body: document.getElementById("editor_body").value,
    });
  },
  true
);

document.getElementById("button_delete").addEventListener(
  "click",
  (_) => {
    if (!confirm("¿Estás seguro de que lo quieres borrar?")) {
      return;
    }
    DeletePost(lastPost);
  },
  true
);
