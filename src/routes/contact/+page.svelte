<script lang="ts">
  let name = "";
  let email = "";
  let message = "";

  let output = "";

  const reset = () => { name = ""; email = ""; message = ""; }

  const handleSubmit = (e: SubmitEvent) => {
    e.preventDefault()
    output = "";
    console.log(name, email, message);
    fetch("/api/contact", {
      method: "POST",
      body: JSON.stringify({
        name, email, message
      })
    })
    .then((e:Response) => {
      console.log(e);
      if ( e.status != 200 ) {
        output = "Something went wrong("
        return;
      }
      reset();
      output = "Got it!"
    })
    .catch(() => {
      output = "Something went wrong((("
    })
  }
</script>

<form on:submit={handleSubmit}>
  <h1>Enter your information:</h1>
  <div><p>Name:</p><input class="bg-bg-3 dark:bg-d-bg-3" type="text" bind:value={name} required/></div>
  <div><p>Email:</p><input class="bg-bg-3 dark:bg-d-bg-3" type="email" bind:value={email} required/></div>
  <div><p>Message:</p><input class="bg-bg-3 dark:bg-d-bg-3" type="text" bind:value={message} required/></div>
  <input type="submit" value="send" />
  {output ? output : ""}
</form>
