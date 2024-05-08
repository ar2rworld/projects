import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	server: {
		proxy: {
			'^/api': 'https://zxdkzryj7yikjsjvk2gzoao4mq0fjmqu.lambda-url.us-east-2.on.aws/'
		}
	}
});
