/**
 * PostService provides methods to interact with the posts API.
 */
export class PostService {
  /**
   * Creates an instance of PostService.
   * @param {string} [baseUrl='/v0/posts'] - The base URL for the posts API.
   * @param {Object} [headers={}] - The headers to include with each request.
   */
  constructor(baseUrl = "/v0/posts", headers = {}) {
    this.baseUrl = baseUrl;
    this.headers = headers;
  }

  /**
   * Retrieves all posts.
   * @returns {Promise<Array>} A promise that resolves to an array of posts.
   */
  async listPosts(params = {}) {
    const url = new URL(this.baseUrl, window.location.origin);
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        url.searchParams.append(key, value);
      }
    });
    const response = await fetch(url, { headers: this.headers });
    return await response.json();
  }

  /**
   * Creates a new post.
   * @param {Object} post - The post object to create.
   * @param {string} post.title - The title of the post.
   * @param {string} post.body - The body content of the post.
   * @returns {Promise<Object>} A promise that resolves to the created post.
   */
  async createPost(post) {
    const response = await fetch(this.baseUrl, {
      method: "POST",
      headers: this.headers,
      body: JSON.stringify(post),
    });
    return await response.json();
  }

  /**
   * Updates an existing post.
   * @param {Object} post - The post object to update.
   * @param {string} post.id - The unique identifier of the post.
   * @param {string} post.title - The updated title of the post.
   * @param {string} post.body - The updated body content of the post.
   * @returns {Promise<Object>} A promise that resolves to the updated post.
   */
  async updatePost(post) {
    const url = `${this.baseUrl}/${encodeURIComponent(post.id)}`;
    const response = await fetch(url, {
      method: "PATCH",
      headers: this.headers,
      body: JSON.stringify(post),
    });
    return await response.json();
  }

  /**
   * Deletes a post.
   * @param {Object} post - The post object to delete.
   * @param {string} post.id - The unique identifier of the post.
   * @returns {Promise<Response>} A promise that resolves to the deletion response.
   */
  async deletePost(post) {
    const url = `${this.baseUrl}/${encodeURIComponent(post.id)}`;
    return await fetch(url, {
      method: "DELETE",
      headers: this.headers,
    });
  }

  /**
   * Get a post.
   * @param {Object} post - The post object to get.
   * @param {string} post.id - The unique identifier of the post.
   * @returns {Promise<Response>} A promise that resolves to get response.
   */
  async getPost(post) {
    const url = `${this.baseUrl}/${encodeURIComponent(post.id)}`;
    const response = await fetch(url, {
      method: "GET",
      headers: this.headers,
    });
    return await response.json();
  }
}
