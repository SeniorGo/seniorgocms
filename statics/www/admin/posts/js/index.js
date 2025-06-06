import { PostService } from "./post-service.js";

const postsTableBody = document.getElementById("posts-tbody");
const confirmDialog = document.getElementById("confirm-dialog");
const deleteBtnOfDialog = document.getElementById("delete-btn-dialog");
const cancelBtnOfDialog = document.getElementById("cancel-btn-dialog");
const tagFilterInput = document.getElementById("tag-filter");

let selectedPostToDelete = null;
let allPosts = [];

const fakeHeaders = {
  "X-Glue-Authentication": JSON.stringify({
    user: {
      id: "user-fake-id",
      nick: "Fulanez",
    },
  }),
};

const postService = new PostService("/v0/posts", fakeHeaders);

function addPost(post) {
  console.log(post);

  const tr = document.createElement("tr");
  let td = document.createElement("td");
  td.textContent = post.id;
  tr.appendChild(td);

  td = document.createElement("td");
  td.textContent = post.title;
  tr.appendChild(td);

  td = document.createElement("td");
  td.textContent = post.tags ? post.tags.join(", ") : "";
  tr.appendChild(td);

  td = document.createElement("td");
  td.textContent = post.creation_time;
  tr.appendChild(td);

  td = document.createElement("td");
  td.textContent = post.modification_time;
  tr.appendChild(td);

  td = document.createElement("td");

  const editButton = document.createElement("a");
  editButton.href = `/admin/posts/editor?id=${post.id}`;
  editButton.classList.add("btn", "btn-sm", "btn-warning", "edit-btn");
  editButton.textContent = "Editar";
  td.appendChild(editButton);

  const deleteButton = document.createElement("button");
  deleteButton.classList.add("btn", "btn-sm", "btn-error", "delete-btn");
  deleteButton.textContent = "Eliminar";
  deleteButton.addEventListener("click", () => {
    selectedPostToDelete = post;
    confirmDialog.showModal();
  });
  td.appendChild(deleteButton);
  tr.appendChild(td);

  postsTableBody.appendChild(tr);
  return tr;
}

function filterPosts() {
  const tagFilter = tagFilterInput.value.toLowerCase().trim();
  postsTableBody.innerHTML = "";
  
  allPosts.forEach(post => {
    if (!tagFilter || (post.tags && post.tags.some(tag => tag.toLowerCase().includes(tagFilter)))) {
      addPost(post);
    }
  });
}

async function deletePost(id) {
  try {
    const response = await postService.deletePost({ id });
    if (response.status !== 204) {
      const payload = await response.json();
      alert(payload.error.description);
      return;
    }

    await loadPosts();
  } catch (error) {
    console.error("Error deleting post:", error);
  }
}

async function loadPosts() {
  try {
    allPosts = await postService.listPosts();
    filterPosts();
  } catch (error) {
    console.error("Error loading posts:", error);
  }
}

deleteBtnOfDialog.addEventListener("click", async () => {
  if (selectedPostToDelete !== null) {
    await deletePost(selectedPostToDelete.id);
    selectedPostToDelete = null;
  }
});

cancelBtnOfDialog.addEventListener("click", () => {
  selectedPostToDelete = null;
});

tagFilterInput.addEventListener("input", filterPosts);

loadPosts();
