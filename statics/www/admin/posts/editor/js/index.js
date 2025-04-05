import { PostService } from "../../js/post-service.js";

const postForm = document.getElementById("post-form");
const postTitleInput = document.getElementById("post-title");
const postBodyInput = document.getElementById("post-body");
const postTagsInput = document.getElementById("post-tags");
const editorBtn = document.getElementById("editor-btn");

let currentPostId = null;

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

  const postRequest = {
    title: postTitleInput.value,
    body: postBodyInput.value,
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
    postBodyInput.value = post.body;
    postTagsInput.value = post.tags ? post.tags.join(", ") : "";
  }
});
