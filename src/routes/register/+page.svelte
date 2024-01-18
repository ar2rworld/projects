<script lang="ts">

  import { Label } from "@smui/button";
  import Button from "@smui/button/src/Button.svelte";
  import Textfield from "@smui/textfield";
	import { api } from "../../axios/axios";
	import type { IMe } from "../../types/me";
	import { goto } from "$app/navigation";
	import axios, { AxiosError } from "axios";
	import LayoutGrid, { Cell } from '@smui/layout-grid';

  let username: string = "";
  let password: string = "";
  let favAnimal: string = "";

  let showError: boolean;
  let errorMessage: string;

  const handleSubmit = async () => {
    const payload = { username, password, favanimal: favAnimal };
    
    try {
      const response = await api.post("/register", payload);
      if (response.status == axios.HttpStatusCode.Ok) {
        goto("/login?new");
      }
    } catch (e) {
      const error = e as AxiosError;

      showError = true;
      const data = error.response?.data as {ok: boolean, message: string};
      console.log(data);
      errorMessage = data.message;
    }
  }
  
</script>
<LayoutGrid>

  <Cell class='text-center' span={4}>
    { #if showError }
      <Label>{errorMessage}</Label>
    { /if }
  </Cell>

  <Cell span={4}>
    <form class='flex flex-col items-center justify-center h-full' on:submit={handleSubmit}>
      <Textfield label="Username" type="text" bind:value={username} required />

      <Textfield label="Password" type="password" bind:value={password}  required />

      <Textfield label="Favorite animal?" type="text" bind:value={favAnimal} />
      
      <Button>
        <Label>Register</Label>
      </Button>
      <a href="/login">Already with us?</a>
    </form>
  </Cell>

  <Cell span={4}>

  </Cell>
</LayoutGrid>
