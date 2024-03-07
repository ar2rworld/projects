<script lang="ts">
  let name = "a";
  let email = "a@gmjk.a";
  let message = "heillo";

  const handleSubmit = (e: SubmitEvent) => {
    e.preventDefault()
    console.log(name, email, message);
    fetch("/api/contact", {
      method: "POST",
      body: JSON.stringify({
        name, email, message
      })
    })
  }
</script>

<form on:submit={handleSubmit}>
  <h1>Enter your information:</h1>
  <div><p>Name:</p><input class="bg-bg-3 dark:bg-d-bg-3" type="text" bind:value={name} required/></div>
  <div><p>Email:</p><input class="bg-bg-3 dark:bg-d-bg-3" type="email" bind:value={email} required/></div>
  <div><p>Message:</p><input class="bg-bg-3 dark:bg-d-bg-3" type="text" bind:value={message} required/></div>
  <input type="submit" value="send" />
</form>
