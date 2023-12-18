<script lang="ts">
  import LayoutGrid, { Cell } from '@smui/layout-grid';
  import Button, { Label } from '@smui/button';
  import Textfield from '@smui/textfield';

  import axios, { Axios, AxiosError } from 'axios';
  import type { ILoginTokens } from '../../types/login';
	import { LoginTokens } from '../../stores/loginTokens';

  let username: string = "a";
  let password: string = "dmin";

  let message = "";
  let showMessage = false;
  const baseUrl = "http://localhost:8050/projectsapi";

  let error = "";

  const handleSubmit = async () => {
    error = "";
    message = "";

    const data = { username, password };

    axios.post(baseUrl + "/login", data)
    .then(d => {
      const loginTokens = d.data as ILoginTokens;
      LoginTokens.set(loginTokens);
      console.log(loginTokens);
    })
    .catch(e => {
      console.log(e)
      if ( e instanceof axios.AxiosError ) {
        message = e.response?.data?.message
        switch(e.code){
          case "ERR_NETWORK":
            error = "Sorry, backend is currently down &#128556"
          break;
          case "ERR_BAD_REQUEST":
            error = "Looks like your credentials are incorrect &#128517"
          break;
          default: {
            error = "some other exception"
          }
        } 
      }
    })
  }

</script>

<form on:submit|preventDefault={handleSubmit}>
  <LayoutGrid cols="3" row-gap="10px">
    <Cell span="4">
      { #if error != "" }
        <p
          on:mouseover={ () => showMessage = true }
          on:mouseleave={ () => showMessage = false }
          on:focus={ () => showMessage = true }
        >
          {@html error}
        </p>
      { /if }
    </Cell>
    <Cell span="4">
      <div class="centered-container">
        <div class="item">
          <Textfield label="username" type="text" bind:value={username} required />
        </div>
        <div class="item">
          <Textfield label="password" type="password" bind:value={password} required />
        </div>
        <div class="item">
          <Button color="primary" variant="raised" type="submit">
            <Label>login</Label>
          </Button>
        </div>
      </div>
    </Cell>
    <Cell span="4">
      { #if showMessage }
          <p>{@html message}</p>
      { /if }
    </Cell>
  </LayoutGrid>
</form>

<style>
  form {
    padding: 0.5%;
    background-color: var(--bg-2);
    color: var(--fg-1);
    margin: 1% 5%;
  }

  .centered-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
  }
  .item {
    margin-bottom: 0.5%;
  }
</style>
