<script lang="ts">
  import LayoutGrid, { Cell } from '@smui/layout-grid';
  import Button, { Label } from '@smui/button';
  import Textfield from '@smui/textfield';
  import Checkbox from '@smui/checkbox';
  import { onDestroy } from 'svelte';

  import axios, { AxiosError } from 'axios';
  import type { ILoginTokens } from '../../types/login';
  import type { IMe } from '../../types/me';
	import { LoginTokens } from '../../stores/loginTokens';
  import { Me } from '../../stores/me';

  let me: IMe;

  let FormUsername: string = "a";
  let FormPassword: string = "admin";

  let message = "";
  let showMessage = false;
  const baseUrl = "http://localhost:8050/projectsapi";

  let error = "";

  const handleSubmit = async () => {
    error = "";
    message = "";

    const data = { username: FormUsername, password: FormPassword };

    axios.post(baseUrl + "/login", data)
    .then(d => {
      const loginTokens = d.data as ILoginTokens;
      LoginTokens.set(loginTokens);

      const m = d.data as IMe;

      Me.set(m);
    })
    .catch((e: AxiosError) => {
      const data = e.response?.data as { ok: boolean, message: string }
      message = data.message;

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
    })
  }

  const logout = () => {
    Me.set({ Username: "", FirstName: "", LastName: "", FullName: "", EmailVerified: false });
    LoginTokens.set({ accessToken: "", refreshToken: "", expiresIn: 0});

    // TODO: call /logout
  }

  const unsubscribe = Me.subscribe(v => {
    me = v;
  });

  onDestroy(() => unsubscribe());
</script>

{ #if ! me.Username }
  <form on:submit|preventDefault={handleSubmit}>
    <LayoutGrid>
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
            <Textfield label="username" type="text" bind:value={FormUsername} required />
          </div>
          <div class="item">
            <Textfield label="password" type="password" bind:value={FormPassword} required />
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
{ :else }
  <form>
    <LayoutGrid>
      <Cell span="12"><p>Hello, {me.FullName}</p></Cell>
      <Cell span="6">
        <Cell span="4"><Textfield input$readonly={true} bind:value={me.FirstName} label="First Name"></Textfield></Cell>
        <Cell span="4"><Textfield input$readonly={true} bind:value={me.LastName} label="Last Name"></Textfield></Cell>
        <Cell span="4"><Textfield input$readonly={true} bind:value={me.Username} label="Username"></Textfield></Cell>
        <Cell span="4" style="display:flex;"><Label>Email Verified</Label><Checkbox disabled bind:checked={me.EmailVerified} /></Cell>
      </Cell>
      <Cell span="6">
        <p>This should be a page with your information and actions:</p>
        <Button variant="raised"><Label>Edit</Label></Button>
        <Button on:click={logout} variant="outlined"><Label>Log out</Label></Button>
      </Cell>
    </LayoutGrid>
  </form>
{ /if }

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
