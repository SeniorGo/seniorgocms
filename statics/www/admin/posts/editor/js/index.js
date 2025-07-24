import { PostService } from "../../js/post-service.js";

const postForm = document.getElementById("post-form");
const postTitleInput = document.getElementById("post-title");
const postTagsInput = document.getElementById("post-tags");
const editorBtn = document.getElementById("editor-btn");

let currentPostId = null;

let easyMDE = null

const fakeHeaders = {
  "X-Glue-Authentication": JSON.stringify({
    user: {
      id: "user-fake-id",
      nick: "Fulanez",
    },
  }),
};

const postService = new PostService("/v0/posts", fakeHeaders);

function parseTags(tagsString) {
  if (!tagsString) return [];
  return tagsString.split(",").map(tag => tag.trim()).filter(tag => tag.length > 0);
}

postForm.addEventListener("submit", async (e) => {
  e.preventDefault();

  if (!easyMDE.value().trim()) {
    alert("El contenido del post no puede estar vacío");
    return;
  }

  const postRequest = {
    title: postTitleInput.value,
    body: easyMDE.value(),
    tags: parseTags(postTagsInput.value),
  };

  if (currentPostId !== null) {
    await postService.updatePost({
      ...postRequest,
      id: currentPostId,
    });
  } else {
    await postService.createPost(postRequest);
  }
  location.href = "/admin/posts";
});


document.addEventListener("DOMContentLoaded", async (event) => {
  const searchParams = new URLSearchParams(window.location.search);
  if (searchParams.has("id")) {
    currentPostId = searchParams.get("id");

    const post = await postService.getPost({ id: currentPostId });
    postTitleInput.value = post.title;
    postTagsInput.value = post.tags ? post.tags.join(", ") : "";

    document.getElementById("markdown-editor").value = post.body
  }

  easyMDE = new EasyMDE({
    element: document.getElementById("markdown-editor"),
    spellChecker: false,
    placeholder: "Escribe aquí con Markdown...",
  });

  if (currentPostId !== null) {
    easyMDE.value(document.getElementById("markdown-editor").value);
  }
});
