<script lang="ts">
	import "../app.css";
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import Hint from '../components/Hint.svelte';
	import HintFeedback from '../components/HintFeedback.svelte';
	import { LoginTokens } from '../stores/loginTokens';
	import { Me } from '../stores/me';

	const STORAGE_KEY = 'theme';
	const DARK_PREFERENCE = '(prefers-color-scheme: dark)';

	const THEMES = {
		DARK: 'dark',
		LIGHT: 'light'
	};
	let currentTheme = THEMES.DARK;

	const prefersDarkThemes = () => window.matchMedia(DARK_PREFERENCE).matches;

	const toggleTheme = () => {
		const stored = localStorage.getItem(STORAGE_KEY);

		if (stored) {
			// clear storage
			localStorage.removeItem(STORAGE_KEY);
		} else {
			// store opposite of preference
			localStorage.setItem(STORAGE_KEY, prefersDarkThemes() ? THEMES.LIGHT : THEMES.DARK);
		}
		applyTheme();
	};
	const applyTheme = () => {
		const preferredTheme = prefersDarkThemes() ? THEMES.DARK : THEMES.LIGHT;
		currentTheme = localStorage.getItem(STORAGE_KEY) ?? preferredTheme;

		if (currentTheme === THEMES.DARK) {
			// for smui
			document.body.classList.remove(THEMES.LIGHT);
			document.body.classList.add(THEMES.DARK);

			// for tailwindcss
			document.documentElement.classList.remove(THEMES.LIGHT);
			document.documentElement.classList.add(THEMES.DARK);
		} else {
			// for smui
			document.body.classList.remove(THEMES.DARK);
			document.body.classList.add(THEMES.LIGHT);
			
			// for tailwindcss
			document.documentElement.classList.remove(THEMES.DARK);
			document.documentElement.classList.add(THEMES.LIGHT);
		}
	};

	const setupLoginTokens = () => {
		let accessToken = '';
		let refreshToken = '';
		let expiresIn = '';
		accessToken = localStorage.getItem('accessToken') || '';
		refreshToken = localStorage.getItem('refreshToken') || '';
		expiresIn = localStorage.getItem('expiresIn') || '';
		if (accessToken) {
			LoginTokens.update((v) => {
				v.accessToken = accessToken;
				return v;
			});
		}
		if (refreshToken) {
			LoginTokens.update((v) => {
				v.refreshToken = refreshToken;
				return v;
			});
		}
		if (expiresIn) {
			LoginTokens.update((v) => {
				v.expiresIn = Number(expiresIn);
				return v;
			});
		}

		LoginTokens.subscribe((v) => {
			localStorage.setItem('accessToken', v.accessToken);
			localStorage.setItem('refreshToken', v.refreshToken);
			localStorage.setItem('expiresIn', v.expiresIn.toString());
		});
	};

	let username: string;
	const setupMe = () => {
		Me.subscribe((v) => {
			username = v.Username;
		});
	};

	onMount(() => {
		applyTheme();
		window.matchMedia(DARK_PREFERENCE).addEventListener('change', applyTheme);
		setupLoginTokens();
		setupMe();
	});
</script>

<svelte:head>
	{#if currentTheme == THEMES.DARK}
		<link rel="stylesheet" href="/smui-dark.css" />
	{:else}
		<link rel="stylesheet" href="/smui.css" />
	{/if}
</svelte:head>

<nav class="flex relative gap-4 p-4 rounded-lg bg-bg-2 dark:bg-d-bg-2">
	<a href="/" aria-current={$page.url.pathname === '/'}> home </a>

	<a href="/projects" aria-current={$page.url.pathname === '/projects'}> projects </a>
	<a href="/about" aria-current={$page.url.pathname === '/about'}> about </a>
	<a href="/blog" aria-current={$page.url.pathname === '/blog'}> blog </a>
	<a href="/shop" aria-current={$page.url.pathname === '/shop'}> shop </a>
	{ #if ! username }
		<a href="/login" aria-current={$page.url.pathname === '/login'}>login</a>
	{ :else }
		<a href="/me" aria-current={$page.url.pathname === '/me'}>me</a>
	{ /if}
	<HintFeedback />
	<input
		class="ml-auto"
		type="checkbox"
		checked={currentTheme !== THEMES.DARK}
		on:click={toggleTheme}
	/>
	<Hint />
</nav>

<slot />

<style>
	nav a[aria-current='true'] {
		border-bottom: 2px solid;
	}
</style>
