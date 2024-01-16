<script lang='ts'>
  import LayoutGrid, { Cell } from '@smui/layout-grid';
	import Button, { Label } from '@smui/button';
	import Textfield from '@smui/textfield';
	import Checkbox from '@smui/checkbox';
	import { onDestroy, onMount } from 'svelte';

  let me: IMe;

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