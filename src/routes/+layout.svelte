<script>
  import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import Hint from '../components/Hint.svelte';
	import HintFeedback from '../components/HintFeedback.svelte';

  const STORAGE_KEY = 'theme';
  const DARK_PREFERENCE = '(prefers-color-scheme: dark)';

  const THEMES = {
    DARK: 'dark',
    LIGHT: 'light',
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
      document.body.classList.remove(THEMES.LIGHT);
      document.body.classList.add(THEMES.DARK);
    } else {
      document.body.classList.remove(THEMES.DARK);
      document.body.classList.add(THEMES.LIGHT);
    }
  };

  onMount(() => {
    applyTheme();
    window.matchMedia(DARK_PREFERENCE).addEventListener('change', applyTheme);
  });
</script>

<svelte:head>
{ #if currentTheme == THEMES.DARK }
  <link rel="stylesheet" href="/smui-dark.css" />
{ :else }
  <link rel="stylesheet" href="/smui.css" />
{ /if }
</svelte:head>


<nav>
	<a href="/" aria-current={$page.url.pathname === '/'}>
		home
	</a>

  <a href="/projects" aria-current={$page.url.pathname === '/projects'}>
		projects
	</a>
	<a href="/about" aria-current={$page.url.pathname === '/about'}>
		about
	</a>
	<a href="/blog" aria-current={$page.url.pathname === '/blog'}>
		blog
	</a>
	<a href="/shop" aria-current={$page.url.pathname === '/shop'}>
		shop
	</a>
  <a href="/login" aria-current={$page.url.pathname === '/login'}>
    login
  </a>
  <HintFeedback />
  <input class='toggleTheme' type="checkbox" checked={currentTheme !== THEMES.DARK} on:click={toggleTheme} />
  <Hint />
</nav>

<slot />

<style>
  .toggleTheme {
    margin-left:auto;
    text-align: end;
  }
</style>