<script lang="ts">
  let username: string = "";
  let password: string = "";

  const baseUrl = "https://192.18.155.170/projectsapi";

  const handleSubmit = (e: SubmitEvent) => {
    e.preventDefault();
    console.log("submit");

    fetch(baseUrl + "/login", { body: JSON.stringify({ "username": username, "password": password})})
    .then(v => console.log(v.json()))
    .catch(e => console.log(e))
  }

</script>

<div>
  <form on:submit={handleSubmit}>
    <input value={username} placeholder="username" />
    <input value={password} placeholder="password" type="password" />

    <button>
      login
    </button>
  </form>
</div>

<style>
  form {
    margin: auto;
  }

  input {
  }

  button {
    margin: auto;
    border: none;
    color: var(--fg-2);
    background: var(--bg-3);
  }
</style>
