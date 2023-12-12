<script lang="ts">
  let username: string = "";
  let password: string = "";

  const baseUrl = "https://192.18.155.170/projectsapi";

  const handleSubmit = async () => {  
    const data = {
      username,
      password
    };

    fetch(baseUrl + "/login",
      {
        method: "POST",
        body: JSON.stringify(data)
      }
    )
    .then(v => console.log(v.json()))
    .catch(e => console.log(e));
  }

</script>

<div>
  <form on:submit|preventDefault={handleSubmit}>
    <input placeholder="username" type="text" bind:value={username} required />
    <input placeholder="password" type="password" bind:value={password} required />

    <button type="submit">
      login
    </button>
  </form>
</div>

<style>
  form {
    margin: auto;
  }

  input {
    padding: 0.5%;
    color: var(--fg-1);
    background: var(--bg-3);
    border-color: var(--bg-2);
  }

  button {
    padding: 0.5%;
    margin: auto;
    border: none;
    color: var(--fg-2);
    background: var(--bg-3);
  }
</style>
