<script lang='ts'>
  import LayoutGrid, { Cell } from '@smui/layout-grid';
	import Button, { Label } from '@smui/button';
	import Textfield from '@smui/textfield';
	import Checkbox from '@smui/checkbox';
	import { onDestroy, onMount } from 'svelte';
	import type { IMe } from '../../types/me';
	import { Me } from '../../stores/me';
	import { LoginTokens } from '../../stores/loginTokens';

  let me: IMe;

	const logout = () => {
		Me.set({
			Username: '',
			FirstName: '',
			LastName: '',
			FullName: '',
			EmailVerified: false,
			Error: '',
			Message: '',
			Ok: true
		});
		LoginTokens.set({ accessToken: '', refreshToken: '', expiresIn: 0 });

  	// TODO: call /logout

    window.location.href = '/login';
	};

  const unsubscribe = Me.subscribe((v) => {
    me = v;
  });
  
  onDestroy(() => unsubscribe());
</script>

<form>
  <LayoutGrid>
    <Cell span={12}><p>Hello, {me.FullName}</p></Cell>
    <Cell span={6}>
      <Cell span={4}><Textfield input$readonly={true} bind:value={me.FirstName} label="First Name" /></Cell
      >
      <Cell span={4}><Textfield input$readonly={true} bind:value={me.LastName} label="Last Name" /></Cell
      >
      <Cell span={4}><Textfield input$readonly={true} bind:value={me.Username} label="Username" /></Cell
      >
      <Cell span={4} style="display:flex;"><Label>Email Verified</Label><Checkbox disabled bind:checked={me.EmailVerified} /></Cell
      >
    </Cell>
    <Cell span={6}>
      <p>This should be a page with your information and actions:</p>
      <Button variant="raised"><Label>Edit</Label></Button>
      <Button on:click={logout} variant="outlined"><Label>Log out</Label></Button>
    </Cell>
  </LayoutGrid>
</form>