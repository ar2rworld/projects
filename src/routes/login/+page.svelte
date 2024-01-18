<script lang="ts">
	import LayoutGrid, { Cell } from '@smui/layout-grid';
	import Button, { Label } from '@smui/button';
	import Textfield from '@smui/textfield';
	import { onDestroy, onMount } from 'svelte';
	import DOMPurify from 'dompurify';

	import type { AxiosError } from 'axios';
	import type { ILoginTokens } from '../../types/login';
	import type { IMe } from '../../types/me';
	import { LoginTokens } from '../../stores/loginTokens';
	import { Me } from '../../stores/me';
	import { getMe } from '../../axios/getMe';

	import { goto } from '$app/navigation';
	import { api } from '../../axios/axios';

	export let data: { new: boolean };
	let newUser = data.new;

	let me: IMe;
	let lt: ILoginTokens;

	let FormUsername: string = 'a';
	let FormPassword: string = 'admin';

	let message = '';
	let showMessage = false;
	const baseUrl = 'http://localhost:8050/projectsapi';

	let error = '';

	const handleSubmit = async () => {
		error = '';
		message = '';

		const data = { username: FormUsername, password: FormPassword };

		api
			.post('/login', data)
			.then((d) => {
				const loginTokens = d.data as ILoginTokens;
				LoginTokens.set(loginTokens);

				const m = d.data as IMe;

				Me.set(m);
				goto('/me');
			})
			.catch((e: AxiosError) => {
				const data = e.response?.data as { ok: boolean; message: string };
				message = data.message;

				switch (e.code) {
					case 'ERR_NETWORK':
						error = 'Sorry, backend is currently down &#128556';
						break;
					case 'ERR_BAD_REQUEST':
						error = 'Looks like your credentials are incorrect &#128517';
						break;
					default: {
						error = 'some other exception';
					}
				}
			});
	};

	const unsubscribe = Me.subscribe((v) => {
		me = v;
	});

	const unsubscribe2 = LoginTokens.subscribe(v => {
		console.log("tokens got updated", v)
		lt = v;
	})

	onMount(() => {
		setTimeout(() => {
			if (!me.Username && lt.accessToken) {
				getMe(lt)
					.then(v => {
						console.log(v);
						Me.set(v);
						goto('/me');
					})
					.catch((e) => {
						error = e?.message
						console.log(e);
					});
			}
		}, 1000)
	});

	onDestroy(() => { unsubscribe(); unsubscribe2() });
</script>
	
<LayoutGrid>
	<Cell span={4}>
		{#if error !== ''}
			<p
				on:mouseover={() => (showMessage = true)}
				on:mouseleave={() => (showMessage = false)}
				on:focus={() => (showMessage = true)}
			>
				{@html error}
			</p>
		{/if}
		{ #if newUser }
			<p>Thank you for registering!</p>
		{ /if }
	</Cell>

	<Cell span={4}>
		<form class='p-0.5 flex flex-col items-center justify-center h-full' on:submit|preventDefault={handleSubmit}>
					<Textfield label="username" type="text" bind:value={FormUsername} required />
					<Textfield class='mb-2' label="password" type="password" bind:value={FormPassword} required />
					<Button color="primary" variant="raised" type="submit">
						<Label>login</Label>
					</Button>
				<a href="/register">Don't have account?</a>
	</Cell>

	<Cell span={4}>
		{#if showMessage}
			<p>{DOMPurify.sanitize(message)}</p>
		{/if}
	</Cell>
</LayoutGrid>
